// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "graph-analyzer/api/graph-property/domain"

	mock "github.com/stretchr/testify/mock"
)

// FullGraphService is an autogenerated mock type for the FullGraphService type
type FullGraphService struct {
	mock.Mock
}

type FullGraphService_Expecter struct {
	mock *mock.Mock
}

func (_m *FullGraphService) EXPECT() *FullGraphService_Expecter {
	return &FullGraphService_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields:
func (_m *FullGraphService) Invoke() domain.FullGraphDTO {
	ret := _m.Called()

	var r0 domain.FullGraphDTO
	if rf, ok := ret.Get(0).(func() domain.FullGraphDTO); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.FullGraphDTO)
	}

	return r0
}

// FullGraphService_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type FullGraphService_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
func (_e *FullGraphService_Expecter) Invoke() *FullGraphService_Invoke_Call {
	return &FullGraphService_Invoke_Call{Call: _e.mock.On("Invoke")}
}

func (_c *FullGraphService_Invoke_Call) Run(run func()) *FullGraphService_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FullGraphService_Invoke_Call) Return(_a0 domain.FullGraphDTO) *FullGraphService_Invoke_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FullGraphService_Invoke_Call) RunAndReturn(run func() domain.FullGraphDTO) *FullGraphService_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFullGraphService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFullGraphService creates a new instance of FullGraphService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFullGraphService(t mockConstructorTestingTNewFullGraphService) *FullGraphService {
	mock := &FullGraphService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}