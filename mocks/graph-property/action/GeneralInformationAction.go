// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// GeneralInformationAction is an autogenerated mock type for the GeneralInformationAction type
type GeneralInformationAction struct {
	mock.Mock
}

type GeneralInformationAction_Expecter struct {
	mock *mock.Mock
}

func (_m *GeneralInformationAction) EXPECT() *GeneralInformationAction_Expecter {
	return &GeneralInformationAction_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields: context
func (_m *GeneralInformationAction) Invoke(context *gin.Context) {
	_m.Called(context)
}

// GeneralInformationAction_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type GeneralInformationAction_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
//   - context *gin.Context
func (_e *GeneralInformationAction_Expecter) Invoke(context interface{}) *GeneralInformationAction_Invoke_Call {
	return &GeneralInformationAction_Invoke_Call{Call: _e.mock.On("Invoke", context)}
}

func (_c *GeneralInformationAction_Invoke_Call) Run(run func(context *gin.Context)) *GeneralInformationAction_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *GeneralInformationAction_Invoke_Call) Return() *GeneralInformationAction_Invoke_Call {
	_c.Call.Return()
	return _c
}

func (_c *GeneralInformationAction_Invoke_Call) RunAndReturn(run func(*gin.Context)) *GeneralInformationAction_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewGeneralInformationAction interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeneralInformationAction creates a new instance of GeneralInformationAction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeneralInformationAction(t mockConstructorTestingTNewGeneralInformationAction) *GeneralInformationAction {
	mock := &GeneralInformationAction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
