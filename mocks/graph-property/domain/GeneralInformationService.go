// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "graph-analyzer/api/graph-property/domain"

	mock "github.com/stretchr/testify/mock"
)

// GeneralInformationService is an autogenerated mock type for the GeneralInformationService type
type GeneralInformationService struct {
	mock.Mock
}

type GeneralInformationService_Expecter struct {
	mock *mock.Mock
}

func (_m *GeneralInformationService) EXPECT() *GeneralInformationService_Expecter {
	return &GeneralInformationService_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields:
func (_m *GeneralInformationService) Invoke() domain.GeneralInformationDTO {
	ret := _m.Called()

	var r0 domain.GeneralInformationDTO
	if rf, ok := ret.Get(0).(func() domain.GeneralInformationDTO); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.GeneralInformationDTO)
	}

	return r0
}

// GeneralInformationService_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type GeneralInformationService_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
func (_e *GeneralInformationService_Expecter) Invoke() *GeneralInformationService_Invoke_Call {
	return &GeneralInformationService_Invoke_Call{Call: _e.mock.On("Invoke")}
}

func (_c *GeneralInformationService_Invoke_Call) Run(run func()) *GeneralInformationService_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GeneralInformationService_Invoke_Call) Return(_a0 domain.GeneralInformationDTO) *GeneralInformationService_Invoke_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GeneralInformationService_Invoke_Call) RunAndReturn(run func() domain.GeneralInformationDTO) *GeneralInformationService_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGeneralInformationService interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeneralInformationService creates a new instance of GeneralInformationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeneralInformationService(t mockConstructorTestingTNewGeneralInformationService) *GeneralInformationService {
	mock := &GeneralInformationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
