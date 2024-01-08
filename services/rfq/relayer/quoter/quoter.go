// Package quoter submits quotes to the RFQ API for which assets the relayer is willing to relay.
package quoter

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"golang.org/x/exp/slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
)

var logger = log.Logger("quoter")

// Quoter submits quotes to the RFQ API.
type Quoter interface {
	// SubmitAllQuotes submits all quotes to the RFQ API.
	SubmitAllQuotes(ctx context.Context) (err error)
	// ShouldProcess determines if a quote should be processed.
	// We do this by either saving all quotes in-memory, and refreshing via GetSelfQuotes() through the API
	// The first comparison is does bridge transaction OriginChainID+TokenAddr match with a quote + DestChainID+DestTokenAddr, then we look to see if we have enough amount to relay it + if the price fits our bounds (based on that the Relayer is relaying the destination token for the origin)
	// validateQuote(BridgeEvent)
	ShouldProcess(ctx context.Context, quote reldb.QuoteRequest) bool
}

// Manager submits quotes to the RFQ API.
// TODO: should be unexported.
type Manager struct {
	// config is the relayer's config.
	config relconfig.Config
	// inventoryManager is used to get the relayer's inventory.
	inventoryManager inventory.Manager
	// rfqClient is used to communicate with the RFQ API.
	rfqClient rfqAPIClient.AuthenticatedClient
	// relayerSigner is the signer used by the relayer to interact on chain
	relayerSigner signer.Signer
	// feePricer is used to price fees.
	feePricer pricer.FeePricer
	// metricsHandler handles traces, etc
	metricsHandler metrics.Handler
	// quotableTokens is a map of token -> list of quotable tokens.
	// should be removed in config overhaul
	quotableTokens map[string][]string
}

// NewQuoterManager creates a new QuoterManager.
func NewQuoterManager(config relconfig.Config, metricsHandler metrics.Handler, inventoryManager inventory.Manager, relayerSigner signer.Signer, feePricer pricer.FeePricer) (Quoter, error) {
	apiClient, err := rfqAPIClient.NewAuthenticatedClient(metricsHandler, config.GetRfqAPIURL(), relayerSigner)
	if err != nil {
		return nil, fmt.Errorf("error creating RFQ API client: %w", err)
	}

	qt := make(map[string][]string)

	// fix any casing issues.
	for token, destTokens := range config.QuotableTokens {
		processedDestTokens := make([]string, len(destTokens))
		for i := range destTokens {
			processedDestTokens[i] = strings.ToLower(destTokens[i])
		}
		qt[strings.ToLower(token)] = processedDestTokens
	}

	return &Manager{
		config:           config,
		inventoryManager: inventoryManager,
		rfqClient:        apiClient,
		quotableTokens:   qt,
		relayerSigner:    relayerSigner,
		metricsHandler:   metricsHandler,
		feePricer:        feePricer,
	}, nil
}

// ShouldProcess determines if a quote should be processed.
func (m *Manager) ShouldProcess(ctx context.Context, quote reldb.QuoteRequest) bool {
	// allowed pairs for this origin token on the destination
	destPairs := m.quotableTokens[quote.GetOriginIDPair()]
	if !(slices.Contains(destPairs, strings.ToLower(quote.GetDestIDPair()))) {
		return false
	}

	// handle decimals.
	// this will never get hit if we're operating correctly.
	if quote.OriginTokenDecimals != quote.DestTokenDecimals {
		logger.Errorf("Pairing tokens with two different decimals is disabled as a safety feature right now.")
		return false
	}

	// then check if we'll make money on it
	isProfitable, err := m.isProfitableQuote(ctx, quote)
	if err != nil {
		logger.Errorf("Error checking if quote is profitable: %v", err)
		return false
	}
	return isProfitable
}

// isProfitableQuote determines if a quote is profitable, i.e. we will not lose money on it, net of fees.
func (m *Manager) isProfitableQuote(ctx context.Context, quote reldb.QuoteRequest) (bool, error) {
	destTokenID, err := m.config.GetTokenName(quote.Transaction.DestChainId, quote.Transaction.DestToken.String())
	if err != nil {
		return false, fmt.Errorf("error getting dest token ID: %w", err)
	}
	fee, err := m.feePricer.GetTotalFee(ctx, quote.Transaction.OriginChainId, quote.Transaction.DestChainId, destTokenID)
	if err != nil {
		return false, fmt.Errorf("error getting total fee: %w", err)
	}
	// NOTE: this logic assumes that the origin and destination tokens have the same price.
	cost := new(big.Int).Add(quote.Transaction.DestAmount, fee)
	return quote.Transaction.OriginAmount.Cmp(cost) >= 0, nil
}

// SubmitAllQuotes submits all quotes to the RFQ API.
func (m *Manager) SubmitAllQuotes(ctx context.Context) (err error) {
	ctx, span := m.metricsHandler.Tracer().Start(ctx, "submitQuotes")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	inv, err := m.inventoryManager.GetCommitableBalances(ctx)
	if err != nil {
		return fmt.Errorf("error getting commitable balances: %w", err)
	}
	return m.prepareAndSubmitQuotes(ctx, inv)
}

// Prepares and submits quotes based on inventory.
func (m *Manager) prepareAndSubmitQuotes(ctx context.Context, inv map[int]map[common.Address]*big.Int) error {
	var allQuotes []model.PutQuoteRequest

	// First, generate all quotes
	for chainID, balances := range inv {
		for address, balance := range balances {
			quotes, err := m.generateQuotes(ctx, chainID, address, balance)
			if err != nil {
				return err
			}
			allQuotes = append(allQuotes, quotes...)
		}
	}

	// Now, submit all the generated quotes
	for _, quote := range allQuotes {
		if err := m.submitQuote(quote); err != nil {
			return err
		}
	}

	return nil
}

// generateQuotes TODO: THIS LOOP IS BROKEN
// Essentially, if we know a destination chain token balance, then we just need to find which tokens are bridgeable to it.
// We can do this by looking at the quotableTokens map, and finding the key that matches the destination chain token.
// Generates quotes for a given chain ID, address, and balance.
func (m *Manager) generateQuotes(ctx context.Context, chainID int, address common.Address, balance *big.Int) ([]model.PutQuoteRequest, error) {
	destChainCfg, ok := m.config.Chains[chainID]
	if !ok {
		return nil, fmt.Errorf("error getting chain config for destination chain ID %d", chainID)
	}
	destTokenID := fmt.Sprintf("%d-%s", chainID, address.Hex())
	var quotes []model.PutQuoteRequest
	for keyTokenID, itemTokenIDs := range m.quotableTokens {
		for _, tokenID := range itemTokenIDs {
			// TODO: probably a better way to do this.
			if strings.ToLower(tokenID) == strings.ToLower(destTokenID) {
				originStr := strings.Split(keyTokenID, "-")[0]
				origin, err := strconv.Atoi(originStr)
				if err != nil {
					return nil, fmt.Errorf("error converting origin chainID: %w", err)
				}
				destToken, err := m.config.GetTokenName(uint32(chainID), address.Hex())
				if err != nil {
					return nil, fmt.Errorf("error getting dest token ID: %w", err)
				}
				fee, err := m.feePricer.GetTotalFee(ctx, uint32(origin), uint32(chainID), destToken)
				if err != nil {
					return nil, fmt.Errorf("error getting total fee: %w", err)
				}
				originChainCfg, ok := m.config.Chains[origin]
				if !ok {
					return nil, fmt.Errorf("error getting chain config for origin chain ID %d", origin)
				}
				quote := model.PutQuoteRequest{
					OriginChainID:           origin,
					OriginTokenAddr:         strings.Split(keyTokenID, "-")[1],
					DestChainID:             chainID,
					DestTokenAddr:           address.Hex(),
					DestAmount:              balance.String(),
					MaxOriginAmount:         balance.String(),
					FixedFee:                fee.String(),
					OriginFastBridgeAddress: originChainCfg.Bridge,
					DestFastBridgeAddress:   destChainCfg.Bridge,
				}
				quotes = append(quotes, quote)
			}
		}
	}
	return quotes, nil
}

// Submits a single quote.
func (m *Manager) submitQuote(quote model.PutQuoteRequest) error {
	err := m.rfqClient.PutQuote(&quote)
	if err != nil {
		return fmt.Errorf("error submitting quote: %w", err)
	}
	return nil
}
