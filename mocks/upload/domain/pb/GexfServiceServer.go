// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	pb "graph-analyzer/api/upload/domain/pb"

	mock "github.com/stretchr/testify/mock"
)

// GexfServiceServer is an autogenerated mock type for the GexfServiceServer type
type GexfServiceServer struct {
	mock.Mock
}

type GexfServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *GexfServiceServer) EXPECT() *GexfServiceServer_Expecter {
	return &GexfServiceServer_Expecter{mock: &_m.Mock}
}

// ProcessGexf provides a mock function with given fields: _a0, _a1
func (_m *GexfServiceServer) ProcessGexf(_a0 context.Context, _a1 *pb.GexfRequest) (*pb.GexfResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GexfResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GexfRequest) (*pb.GexfResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GexfRequest) *pb.GexfResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GexfResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.GexfRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GexfServiceServer_ProcessGexf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessGexf'
type GexfServiceServer_ProcessGexf_Call struct {
	*mock.Call
}

// ProcessGexf is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *pb.GexfRequest
func (_e *GexfServiceServer_Expecter) ProcessGexf(_a0 interface{}, _a1 interface{}) *GexfServiceServer_ProcessGexf_Call {
	return &GexfServiceServer_ProcessGexf_Call{Call: _e.mock.On("ProcessGexf", _a0, _a1)}
}

func (_c *GexfServiceServer_ProcessGexf_Call) Run(run func(_a0 context.Context, _a1 *pb.GexfRequest)) *GexfServiceServer_ProcessGexf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pb.GexfRequest))
	})
	return _c
}

func (_c *GexfServiceServer_ProcessGexf_Call) Return(_a0 *pb.GexfResponse, _a1 error) *GexfServiceServer_ProcessGexf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GexfServiceServer_ProcessGexf_Call) RunAndReturn(run func(context.Context, *pb.GexfRequest) (*pb.GexfResponse, error)) *GexfServiceServer_ProcessGexf_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedGexfServiceServer provides a mock function with given fields:
func (_m *GexfServiceServer) mustEmbedUnimplementedGexfServiceServer() {
	_m.Called()
}

// GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedGexfServiceServer'
type GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedGexfServiceServer is a helper method to define mock.On call
func (_e *GexfServiceServer_Expecter) mustEmbedUnimplementedGexfServiceServer() *GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call {
	return &GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedGexfServiceServer")}
}

func (_c *GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call) Run(run func()) *GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call) Return() *GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call) RunAndReturn(run func()) *GexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGexfServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewGexfServiceServer creates a new instance of GexfServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGexfServiceServer(t mockConstructorTestingTNewGexfServiceServer) *GexfServiceServer {
	mock := &GexfServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
