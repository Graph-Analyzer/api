// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ReadinessResponder is an autogenerated mock type for the ReadinessResponder type
type ReadinessResponder struct {
	mock.Mock
}

type ReadinessResponder_Expecter struct {
	mock *mock.Mock
}

func (_m *ReadinessResponder) EXPECT() *ReadinessResponder_Expecter {
	return &ReadinessResponder_Expecter{mock: &_m.Mock}
}

// CreateSuccessfulResponse provides a mock function with given fields:
func (_m *ReadinessResponder) CreateSuccessfulResponse() {
	_m.Called()
}

// ReadinessResponder_CreateSuccessfulResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSuccessfulResponse'
type ReadinessResponder_CreateSuccessfulResponse_Call struct {
	*mock.Call
}

// CreateSuccessfulResponse is a helper method to define mock.On call
func (_e *ReadinessResponder_Expecter) CreateSuccessfulResponse() *ReadinessResponder_CreateSuccessfulResponse_Call {
	return &ReadinessResponder_CreateSuccessfulResponse_Call{Call: _e.mock.On("CreateSuccessfulResponse")}
}

func (_c *ReadinessResponder_CreateSuccessfulResponse_Call) Run(run func()) *ReadinessResponder_CreateSuccessfulResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadinessResponder_CreateSuccessfulResponse_Call) Return() *ReadinessResponder_CreateSuccessfulResponse_Call {
	_c.Call.Return()
	return _c
}

func (_c *ReadinessResponder_CreateSuccessfulResponse_Call) RunAndReturn(run func()) *ReadinessResponder_CreateSuccessfulResponse_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewReadinessResponder interface {
	mock.TestingT
	Cleanup(func())
}

// NewReadinessResponder creates a new instance of ReadinessResponder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReadinessResponder(t mockConstructorTestingTNewReadinessResponder) *ReadinessResponder {
	mock := &ReadinessResponder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}