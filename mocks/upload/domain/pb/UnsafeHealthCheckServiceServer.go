// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeHealthCheckServiceServer is an autogenerated mock type for the UnsafeHealthCheckServiceServer type
type UnsafeHealthCheckServiceServer struct {
	mock.Mock
}

type UnsafeHealthCheckServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeHealthCheckServiceServer) EXPECT() *UnsafeHealthCheckServiceServer_Expecter {
	return &UnsafeHealthCheckServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedHealthCheckServiceServer provides a mock function with given fields:
func (_m *UnsafeHealthCheckServiceServer) mustEmbedUnimplementedHealthCheckServiceServer() {
	_m.Called()
}

// UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedHealthCheckServiceServer'
type UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedHealthCheckServiceServer is a helper method to define mock.On call
func (_e *UnsafeHealthCheckServiceServer_Expecter) mustEmbedUnimplementedHealthCheckServiceServer() *UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call {
	return &UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedHealthCheckServiceServer")}
}

func (_c *UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call) Run(run func()) *UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call) Return() *UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call) RunAndReturn(run func()) *UnsafeHealthCheckServiceServer_mustEmbedUnimplementedHealthCheckServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUnsafeHealthCheckServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeHealthCheckServiceServer creates a new instance of UnsafeHealthCheckServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeHealthCheckServiceServer(t mockConstructorTestingTNewUnsafeHealthCheckServiceServer) *UnsafeHealthCheckServiceServer {
	mock := &UnsafeHealthCheckServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
