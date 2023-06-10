// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// ReadinessAction is an autogenerated mock type for the ReadinessAction type
type ReadinessAction struct {
	mock.Mock
}

type ReadinessAction_Expecter struct {
	mock *mock.Mock
}

func (_m *ReadinessAction) EXPECT() *ReadinessAction_Expecter {
	return &ReadinessAction_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields: context
func (_m *ReadinessAction) Invoke(context *gin.Context) {
	_m.Called(context)
}

// ReadinessAction_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type ReadinessAction_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
//   - context *gin.Context
func (_e *ReadinessAction_Expecter) Invoke(context interface{}) *ReadinessAction_Invoke_Call {
	return &ReadinessAction_Invoke_Call{Call: _e.mock.On("Invoke", context)}
}

func (_c *ReadinessAction_Invoke_Call) Run(run func(context *gin.Context)) *ReadinessAction_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *ReadinessAction_Invoke_Call) Return() *ReadinessAction_Invoke_Call {
	_c.Call.Return()
	return _c
}

func (_c *ReadinessAction_Invoke_Call) RunAndReturn(run func(*gin.Context)) *ReadinessAction_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewReadinessAction interface {
	mock.TestingT
	Cleanup(func())
}

// NewReadinessAction creates a new instance of ReadinessAction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReadinessAction(t mockConstructorTestingTNewReadinessAction) *ReadinessAction {
	mock := &ReadinessAction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
