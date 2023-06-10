// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "graph-analyzer/api/graph-property/domain"

	mock "github.com/stretchr/testify/mock"
)

// DegreeDistributionResponder is an autogenerated mock type for the DegreeDistributionResponder type
type DegreeDistributionResponder struct {
	mock.Mock
}

type DegreeDistributionResponder_Expecter struct {
	mock *mock.Mock
}

func (_m *DegreeDistributionResponder) EXPECT() *DegreeDistributionResponder_Expecter {
	return &DegreeDistributionResponder_Expecter{mock: &_m.Mock}
}

// CreateSuccessfulResponse provides a mock function with given fields: dto
func (_m *DegreeDistributionResponder) CreateSuccessfulResponse(dto domain.DegreeDistributionDTO) {
	_m.Called(dto)
}

// DegreeDistributionResponder_CreateSuccessfulResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSuccessfulResponse'
type DegreeDistributionResponder_CreateSuccessfulResponse_Call struct {
	*mock.Call
}

// CreateSuccessfulResponse is a helper method to define mock.On call
//   - dto domain.DegreeDistributionDTO
func (_e *DegreeDistributionResponder_Expecter) CreateSuccessfulResponse(dto interface{}) *DegreeDistributionResponder_CreateSuccessfulResponse_Call {
	return &DegreeDistributionResponder_CreateSuccessfulResponse_Call{Call: _e.mock.On("CreateSuccessfulResponse", dto)}
}

func (_c *DegreeDistributionResponder_CreateSuccessfulResponse_Call) Run(run func(dto domain.DegreeDistributionDTO)) *DegreeDistributionResponder_CreateSuccessfulResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.DegreeDistributionDTO))
	})
	return _c
}

func (_c *DegreeDistributionResponder_CreateSuccessfulResponse_Call) Return() *DegreeDistributionResponder_CreateSuccessfulResponse_Call {
	_c.Call.Return()
	return _c
}

func (_c *DegreeDistributionResponder_CreateSuccessfulResponse_Call) RunAndReturn(run func(domain.DegreeDistributionDTO)) *DegreeDistributionResponder_CreateSuccessfulResponse_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDegreeDistributionResponder interface {
	mock.TestingT
	Cleanup(func())
}

// NewDegreeDistributionResponder creates a new instance of DegreeDistributionResponder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDegreeDistributionResponder(t mockConstructorTestingTNewDegreeDistributionResponder) *DegreeDistributionResponder {
	mock := &DegreeDistributionResponder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
