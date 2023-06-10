// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "graph-analyzer/api/graph-property/domain"

	mock "github.com/stretchr/testify/mock"
)

// GeneralInformationResponder is an autogenerated mock type for the GeneralInformationResponder type
type GeneralInformationResponder struct {
	mock.Mock
}

type GeneralInformationResponder_Expecter struct {
	mock *mock.Mock
}

func (_m *GeneralInformationResponder) EXPECT() *GeneralInformationResponder_Expecter {
	return &GeneralInformationResponder_Expecter{mock: &_m.Mock}
}

// CreateSuccessfulResponse provides a mock function with given fields: dto
func (_m *GeneralInformationResponder) CreateSuccessfulResponse(dto domain.GeneralInformationDTO) {
	_m.Called(dto)
}

// GeneralInformationResponder_CreateSuccessfulResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSuccessfulResponse'
type GeneralInformationResponder_CreateSuccessfulResponse_Call struct {
	*mock.Call
}

// CreateSuccessfulResponse is a helper method to define mock.On call
//   - dto domain.GeneralInformationDTO
func (_e *GeneralInformationResponder_Expecter) CreateSuccessfulResponse(dto interface{}) *GeneralInformationResponder_CreateSuccessfulResponse_Call {
	return &GeneralInformationResponder_CreateSuccessfulResponse_Call{Call: _e.mock.On("CreateSuccessfulResponse", dto)}
}

func (_c *GeneralInformationResponder_CreateSuccessfulResponse_Call) Run(run func(dto domain.GeneralInformationDTO)) *GeneralInformationResponder_CreateSuccessfulResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.GeneralInformationDTO))
	})
	return _c
}

func (_c *GeneralInformationResponder_CreateSuccessfulResponse_Call) Return() *GeneralInformationResponder_CreateSuccessfulResponse_Call {
	_c.Call.Return()
	return _c
}

func (_c *GeneralInformationResponder_CreateSuccessfulResponse_Call) RunAndReturn(run func(domain.GeneralInformationDTO)) *GeneralInformationResponder_CreateSuccessfulResponse_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGeneralInformationResponder interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeneralInformationResponder creates a new instance of GeneralInformationResponder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeneralInformationResponder(t mockConstructorTestingTNewGeneralInformationResponder) *GeneralInformationResponder {
	mock := &GeneralInformationResponder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
