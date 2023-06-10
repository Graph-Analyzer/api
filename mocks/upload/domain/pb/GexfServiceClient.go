// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "graph-analyzer/api/upload/domain/pb"
)

// GexfServiceClient is an autogenerated mock type for the GexfServiceClient type
type GexfServiceClient struct {
	mock.Mock
}

type GexfServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *GexfServiceClient) EXPECT() *GexfServiceClient_Expecter {
	return &GexfServiceClient_Expecter{mock: &_m.Mock}
}

// ProcessGexf provides a mock function with given fields: ctx, in, opts
func (_m *GexfServiceClient) ProcessGexf(ctx context.Context, in *pb.GexfRequest, opts ...grpc.CallOption) (*pb.GexfResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.GexfResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GexfRequest, ...grpc.CallOption) (*pb.GexfResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GexfRequest, ...grpc.CallOption) *pb.GexfResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GexfResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.GexfRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GexfServiceClient_ProcessGexf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessGexf'
type GexfServiceClient_ProcessGexf_Call struct {
	*mock.Call
}

// ProcessGexf is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pb.GexfRequest
//   - opts ...grpc.CallOption
func (_e *GexfServiceClient_Expecter) ProcessGexf(ctx interface{}, in interface{}, opts ...interface{}) *GexfServiceClient_ProcessGexf_Call {
	return &GexfServiceClient_ProcessGexf_Call{Call: _e.mock.On("ProcessGexf",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *GexfServiceClient_ProcessGexf_Call) Run(run func(ctx context.Context, in *pb.GexfRequest, opts ...grpc.CallOption)) *GexfServiceClient_ProcessGexf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pb.GexfRequest), variadicArgs...)
	})
	return _c
}

func (_c *GexfServiceClient_ProcessGexf_Call) Return(_a0 *pb.GexfResponse, _a1 error) *GexfServiceClient_ProcessGexf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GexfServiceClient_ProcessGexf_Call) RunAndReturn(run func(context.Context, *pb.GexfRequest, ...grpc.CallOption) (*pb.GexfResponse, error)) *GexfServiceClient_ProcessGexf_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGexfServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewGexfServiceClient creates a new instance of GexfServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGexfServiceClient(t mockConstructorTestingTNewGexfServiceClient) *GexfServiceClient {
	mock := &GexfServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
