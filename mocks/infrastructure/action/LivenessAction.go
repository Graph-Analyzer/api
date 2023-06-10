// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// LivenessAction is an autogenerated mock type for the LivenessAction type
type LivenessAction struct {
	mock.Mock
}

type LivenessAction_Expecter struct {
	mock *mock.Mock
}

func (_m *LivenessAction) EXPECT() *LivenessAction_Expecter {
	return &LivenessAction_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields: context
func (_m *LivenessAction) Invoke(context *gin.Context) {
	_m.Called(context)
}

// LivenessAction_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type LivenessAction_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
//   - context *gin.Context
func (_e *LivenessAction_Expecter) Invoke(context interface{}) *LivenessAction_Invoke_Call {
	return &LivenessAction_Invoke_Call{Call: _e.mock.On("Invoke", context)}
}

func (_c *LivenessAction_Invoke_Call) Run(run func(context *gin.Context)) *LivenessAction_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *LivenessAction_Invoke_Call) Return() *LivenessAction_Invoke_Call {
	_c.Call.Return()
	return _c
}

func (_c *LivenessAction_Invoke_Call) RunAndReturn(run func(*gin.Context)) *LivenessAction_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewLivenessAction interface {
	mock.TestingT
	Cleanup(func())
}

// NewLivenessAction creates a new instance of LivenessAction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLivenessAction(t mockConstructorTestingTNewLivenessAction) *LivenessAction {
	mock := &LivenessAction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}