// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "graph-analyzer/api/graph-property/domain"

	mock "github.com/stretchr/testify/mock"
)

// DiameterService is an autogenerated mock type for the DiameterService type
type DiameterService struct {
	mock.Mock
}

type DiameterService_Expecter struct {
	mock *mock.Mock
}

func (_m *DiameterService) EXPECT() *DiameterService_Expecter {
	return &DiameterService_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields:
func (_m *DiameterService) Invoke() domain.DiameterDTO {
	ret := _m.Called()

	var r0 domain.DiameterDTO
	if rf, ok := ret.Get(0).(func() domain.DiameterDTO); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.DiameterDTO)
	}

	return r0
}

// DiameterService_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type DiameterService_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
func (_e *DiameterService_Expecter) Invoke() *DiameterService_Invoke_Call {
	return &DiameterService_Invoke_Call{Call: _e.mock.On("Invoke")}
}

func (_c *DiameterService_Invoke_Call) Run(run func()) *DiameterService_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DiameterService_Invoke_Call) Return(_a0 domain.DiameterDTO) *DiameterService_Invoke_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DiameterService_Invoke_Call) RunAndReturn(run func() domain.DiameterDTO) *DiameterService_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDiameterService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDiameterService creates a new instance of DiameterService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDiameterService(t mockConstructorTestingTNewDiameterService) *DiameterService {
	mock := &DiameterService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}