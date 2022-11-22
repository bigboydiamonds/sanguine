package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewAttestationCollectorContract returns a bound attestation collector contract.
// nolint: staticcheck
func NewAttestationCollectorContract(ctx context.Context, client chain.Chain, attestationAddress common.Address) (domains.AttestationCollectorContract, error) {
	boundCountract, err := attestationcollector.NewAttestationCollectorRef(attestationAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &attestationcollector.AttestationCollectorRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return attestationCollectorContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type attestationCollectorContract struct {
	// contract contains the conract handle
	contract *attestationcollector.AttestationCollectorRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

func (a attestationCollectorContract) SubmitAttestation(ctx context.Context, signer signer.Signer, attestation types.SignedAttestation) error {
	transactor, err := signer.GetTransactor(a.client.GetBigChainID())
	if err != nil {
		return fmt.Errorf("could not sign tx: %w", err)
	}

	transactOpts, err := a.nonceManager.NewKeyedTransactor(transactor)
	if err != nil {
		return fmt.Errorf("could not create tx: %w", err)
	}

	transactOpts.Context = ctx

	encodedAttestation, err := types.EncodeSignedAttestation(attestation)
	if err != nil {
		return fmt.Errorf("could not get signed attestations: %w", err)
	}

	_, err = a.contract.SubmitAttestation(transactOpts, encodedAttestation)
	if err != nil {
		return fmt.Errorf("could not submit attestation: %w", err)
	}

	return nil
}

func (a attestationCollectorContract) GetLatestNonce(ctx context.Context, origin uint32, destination uint32, signer signer.Signer) (nonce uint32, err error) {
	// TODO (joe): This needs to be refactored after we do the GlobalRegistry stuff
	return 0, nil
	/*latestNonce, err := a.contract.GetLatestNonce(&bind.CallOpts{Context: ctx}, origin, destination, signer)
	if err != nil {
		return 0, fmt.Errorf("could not retrieve latest nonce: %w", err)
	}

	return latestNonce, nil*/
}
