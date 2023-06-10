// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeGexfServiceServer is an autogenerated mock type for the UnsafeGexfServiceServer type
type UnsafeGexfServiceServer struct {
	mock.Mock
}

type UnsafeGexfServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeGexfServiceServer) EXPECT() *UnsafeGexfServiceServer_Expecter {
	return &UnsafeGexfServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedGexfServiceServer provides a mock function with given fields:
func (_m *UnsafeGexfServiceServer) mustEmbedUnimplementedGexfServiceServer() {
	_m.Called()
}

// UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedGexfServiceServer'
type UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedGexfServiceServer is a helper method to define mock.On call
func (_e *UnsafeGexfServiceServer_Expecter) mustEmbedUnimplementedGexfServiceServer() *UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call {
	return &UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedGexfServiceServer")}
}

func (_c *UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call) Run(run func()) *UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call) Return() *UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call) RunAndReturn(run func()) *UnsafeGexfServiceServer_mustEmbedUnimplementedGexfServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUnsafeGexfServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeGexfServiceServer creates a new instance of UnsafeGexfServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeGexfServiceServer(t mockConstructorTestingTNewUnsafeGexfServiceServer) *UnsafeGexfServiceServer {
	mock := &UnsafeGexfServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
