// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "graph-analyzer/api/graph-property/domain"

	mock "github.com/stretchr/testify/mock"
)

// HasCutVertexService is an autogenerated mock type for the HasCutVertexService type
type HasCutVertexService struct {
	mock.Mock
}

type HasCutVertexService_Expecter struct {
	mock *mock.Mock
}

func (_m *HasCutVertexService) EXPECT() *HasCutVertexService_Expecter {
	return &HasCutVertexService_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields:
func (_m *HasCutVertexService) Invoke() domain.HasCutVertexDTO {
	ret := _m.Called()

	var r0 domain.HasCutVertexDTO
	if rf, ok := ret.Get(0).(func() domain.HasCutVertexDTO); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.HasCutVertexDTO)
	}

	return r0
}

// HasCutVertexService_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type HasCutVertexService_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
func (_e *HasCutVertexService_Expecter) Invoke() *HasCutVertexService_Invoke_Call {
	return &HasCutVertexService_Invoke_Call{Call: _e.mock.On("Invoke")}
}

func (_c *HasCutVertexService_Invoke_Call) Run(run func()) *HasCutVertexService_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HasCutVertexService_Invoke_Call) Return(_a0 domain.HasCutVertexDTO) *HasCutVertexService_Invoke_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HasCutVertexService_Invoke_Call) RunAndReturn(run func() domain.HasCutVertexDTO) *HasCutVertexService_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewHasCutVertexService interface {
	mock.TestingT
	Cleanup(func())
}

// NewHasCutVertexService creates a new instance of HasCutVertexService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHasCutVertexService(t mockConstructorTestingTNewHasCutVertexService) *HasCutVertexService {
	mock := &HasCutVertexService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
