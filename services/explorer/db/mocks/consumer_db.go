// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gorm "gorm.io/gorm"

	immutable "github.com/benbjohnson/immutable"

	mock "github.com/stretchr/testify/mock"

	model "github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"

	sql "github.com/synapsecns/sanguine/services/explorer/db/sql"
)

// ConsumerDB is an autogenerated mock type for the ConsumerDB type
type ConsumerDB struct {
	mock.Mock
}

// GetAddressChainRanking provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetAddressChainRanking(ctx context.Context, query string) ([]*model.AddressChainRanking, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.AddressChainRanking
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.AddressChainRanking); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AddressChainRanking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressDailyData provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetAddressDailyData(ctx context.Context, query string) ([]*model.AddressDailyCount, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.AddressDailyCount
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.AddressDailyCount); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AddressDailyCount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressData provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetAddressData(ctx context.Context, query string) (float64, float64, int, error) {
	ret := _m.Called(ctx, query)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, string) float64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func(context.Context, string) float64); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(context.Context, string) int); ok {
		r2 = rf(ctx, query)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, string) error); ok {
		r3 = rf(ctx, query)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetAddressRanking provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetAddressRanking(ctx context.Context, query string) ([]*model.AddressRanking, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.AddressRanking
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.AddressRanking); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AddressRanking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBridgeEvents provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetAllBridgeEvents(ctx context.Context, query string) ([]sql.HybridBridgeEvent, error) {
	ret := _m.Called(ctx, query)

	var r0 []sql.HybridBridgeEvent
	if rf, ok := ret.Get(0).(func(context.Context, string) []sql.HybridBridgeEvent); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sql.HybridBridgeEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllMessageBusEvents provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetAllMessageBusEvents(ctx context.Context, query string) ([]sql.HybridMessageBusEvent, error) {
	ret := _m.Called(ctx, query)

	var r0 []sql.HybridMessageBusEvent
	if rf, ok := ret.Get(0).(func(context.Context, string) []sql.HybridMessageBusEvent); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sql.HybridMessageBusEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBridgeEvent provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetBridgeEvent(ctx context.Context, query string) (*sql.BridgeEvent, error) {
	ret := _m.Called(ctx, query)

	var r0 *sql.BridgeEvent
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.BridgeEvent); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.BridgeEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBridgeEvents provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetBridgeEvents(ctx context.Context, query string) ([]sql.BridgeEvent, error) {
	ret := _m.Called(ctx, query)

	var r0 []sql.BridgeEvent
	if rf, ok := ret.Get(0).(func(context.Context, string) []sql.BridgeEvent); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sql.BridgeEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDailyTotals provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetDailyTotals(ctx context.Context, query string) ([]*model.DateResultByChain, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.DateResultByChain
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.DateResultByChain); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DateResultByChain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDateResults provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetDateResults(ctx context.Context, query string) ([]*model.DateResult, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.DateResult
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.DateResult); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFloat64 provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetFloat64(ctx context.Context, query string) (float64, error) {
	ret := _m.Called(ctx, query)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, string) float64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastStoredBlock provides a mock function with given fields: ctx, chainID, contractAddress
func (_m *ConsumerDB) GetLastStoredBlock(ctx context.Context, chainID uint32, contractAddress string) (uint64, error) {
	ret := _m.Called(ctx, chainID, contractAddress)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint32, string) uint64); ok {
		r0 = rf(ctx, chainID, contractAddress)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, string) error); ok {
		r1 = rf(ctx, chainID, contractAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeaderboard provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetLeaderboard(ctx context.Context, query string) ([]*model.Leaderboard, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.Leaderboard
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.Leaderboard); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Leaderboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMVBridgeEvent provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetMVBridgeEvent(ctx context.Context, query string) (*sql.HybridBridgeEvent, error) {
	ret := _m.Called(ctx, query)

	var r0 *sql.HybridBridgeEvent
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.HybridBridgeEvent); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.HybridBridgeEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingByChain provides a mock function with given fields: ctx
func (_m *ConsumerDB) GetPendingByChain(ctx context.Context) (*immutable.Map[int, int], error) {
	ret := _m.Called(ctx)

	var r0 *immutable.Map[int, int]
	if rf, ok := ret.Get(0).(func(context.Context) *immutable.Map[int, int]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*immutable.Map[int, int])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRankedChainsByVolume provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetRankedChainsByVolume(ctx context.Context, query string) ([]*model.VolumeByChainID, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.VolumeByChainID
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.VolumeByChainID); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VolumeByChainID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetString provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetString(ctx context.Context, query string) (string, error) {
	ret := _m.Called(ctx, query)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringArray provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetStringArray(ctx context.Context, query string) ([]string, error) {
	ret := _m.Called(ctx, query)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenCounts provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetTokenCounts(ctx context.Context, query string) ([]*model.TokenCountResult, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.TokenCountResult
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.TokenCountResult); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TokenCountResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxCounts provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetTxCounts(ctx context.Context, query string) ([]*model.TransactionCountResult, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.TransactionCountResult
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.TransactionCountResult); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TransactionCountResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint64 provides a mock function with given fields: ctx, query
func (_m *ConsumerDB) GetUint64(ctx context.Context, query string) (uint64, error) {
	ret := _m.Called(ctx, query)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreEvent provides a mock function with given fields: ctx, event
func (_m *ConsumerDB) StoreEvent(ctx context.Context, event interface{}) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreEvents provides a mock function with given fields: ctx, events
func (_m *ConsumerDB) StoreEvents(ctx context.Context, events []interface{}) error {
	ret := _m.Called(ctx, events)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}) error); ok {
		r0 = rf(ctx, events)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreLastBlock provides a mock function with given fields: ctx, chainID, blockNumber, contractAddress
func (_m *ConsumerDB) StoreLastBlock(ctx context.Context, chainID uint32, blockNumber uint64, contractAddress string) error {
	ret := _m.Called(ctx, chainID, blockNumber, contractAddress)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint64, string) error); ok {
		r0 = rf(ctx, chainID, blockNumber, contractAddress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreSwapFee provides a mock function with given fields: ctx, chainID, timestamp, contractAddress, fee, feeType
func (_m *ConsumerDB) StoreSwapFee(ctx context.Context, chainID uint32, timestamp uint64, contractAddress string, fee uint64, feeType string) error {
	ret := _m.Called(ctx, chainID, timestamp, contractAddress, fee, feeType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint64, string, uint64, string) error); ok {
		r0 = rf(ctx, chainID, timestamp, contractAddress, fee, feeType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreTokenIndex provides a mock function with given fields: ctx, chainID, tokenIndex, tokenAddress, contractAddress
func (_m *ConsumerDB) StoreTokenIndex(ctx context.Context, chainID uint32, tokenIndex uint8, tokenAddress string, contractAddress string) error {
	ret := _m.Called(ctx, chainID, tokenIndex, tokenAddress, contractAddress)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, string, string) error); ok {
		r0 = rf(ctx, chainID, tokenIndex, tokenAddress, contractAddress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UNSAFE_DB provides a mock function with given fields:
func (_m *ConsumerDB) UNSAFE_DB() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

type mockConstructorTestingTNewConsumerDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumerDB creates a new instance of ConsumerDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumerDB(t mockConstructorTestingTNewConsumerDB) *ConsumerDB {
	mock := &ConsumerDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
