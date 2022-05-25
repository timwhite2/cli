// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	"context"
	"testing"

	types "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// IBCClient is an autogenerated mock type for the QueryClient type
type IBCClient struct {
	mock.Mock
}

// ClientParams provides a mock function with given fields: ctx, in, opts
func (_m *IBCClient) ClientParams(ctx context.Context, in *types.QueryClientParamsRequest, opts ...grpc.CallOption) (*types.QueryClientParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryClientParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryClientParamsRequest, ...grpc.CallOption) *types.QueryClientParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryClientParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryClientParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientState provides a mock function with given fields: ctx, in, opts
func (_m *IBCClient) ClientState(ctx context.Context, in *types.QueryClientStateRequest, opts ...grpc.CallOption) (*types.QueryClientStateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryClientStateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryClientStateRequest, ...grpc.CallOption) *types.QueryClientStateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryClientStateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryClientStateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientStates provides a mock function with given fields: ctx, in, opts
func (_m *IBCClient) ClientStates(ctx context.Context, in *types.QueryClientStatesRequest, opts ...grpc.CallOption) (*types.QueryClientStatesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryClientStatesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryClientStatesRequest, ...grpc.CallOption) *types.QueryClientStatesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryClientStatesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryClientStatesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientStatus provides a mock function with given fields: ctx, in, opts
func (_m *IBCClient) ClientStatus(ctx context.Context, in *types.QueryClientStatusRequest, opts ...grpc.CallOption) (*types.QueryClientStatusResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryClientStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryClientStatusRequest, ...grpc.CallOption) *types.QueryClientStatusResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryClientStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryClientStatusRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsensusState provides a mock function with given fields: ctx, in, opts
func (_m *IBCClient) ConsensusState(ctx context.Context, in *types.QueryConsensusStateRequest, opts ...grpc.CallOption) (*types.QueryConsensusStateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryConsensusStateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryConsensusStateRequest, ...grpc.CallOption) *types.QueryConsensusStateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryConsensusStateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryConsensusStateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsensusStates provides a mock function with given fields: ctx, in, opts
func (_m *IBCClient) ConsensusStates(ctx context.Context, in *types.QueryConsensusStatesRequest, opts ...grpc.CallOption) (*types.QueryConsensusStatesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryConsensusStatesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryConsensusStatesRequest, ...grpc.CallOption) *types.QueryConsensusStatesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryConsensusStatesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryConsensusStatesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpgradedClientState provides a mock function with given fields: ctx, in, opts
func (_m *IBCClient) UpgradedClientState(ctx context.Context, in *types.QueryUpgradedClientStateRequest, opts ...grpc.CallOption) (*types.QueryUpgradedClientStateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryUpgradedClientStateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryUpgradedClientStateRequest, ...grpc.CallOption) *types.QueryUpgradedClientStateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryUpgradedClientStateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryUpgradedClientStateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpgradedConsensusState provides a mock function with given fields: ctx, in, opts
func (_m *IBCClient) UpgradedConsensusState(ctx context.Context, in *types.QueryUpgradedConsensusStateRequest, opts ...grpc.CallOption) (*types.QueryUpgradedConsensusStateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryUpgradedConsensusStateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryUpgradedConsensusStateRequest, ...grpc.CallOption) *types.QueryUpgradedConsensusStateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryUpgradedConsensusStateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryUpgradedConsensusStateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIBCClient creates a new instance of IBCClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIBCClient(t testing.TB) *IBCClient {
	mock := &IBCClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}