// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "graph-analyzer/api/upload/domain"

	mock "github.com/stretchr/testify/mock"
)

// UploadStatusService is an autogenerated mock type for the UploadStatusService type
type UploadStatusService struct {
	mock.Mock
}

type UploadStatusService_Expecter struct {
	mock *mock.Mock
}

func (_m *UploadStatusService) EXPECT() *UploadStatusService_Expecter {
	return &UploadStatusService_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields:
func (_m *UploadStatusService) Invoke() (domain.UploadStatusDTO, error) {
	ret := _m.Called()

	var r0 domain.UploadStatusDTO
	var r1 error
	if rf, ok := ret.Get(0).(func() (domain.UploadStatusDTO, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() domain.UploadStatusDTO); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.UploadStatusDTO)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadStatusService_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type UploadStatusService_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
func (_e *UploadStatusService_Expecter) Invoke() *UploadStatusService_Invoke_Call {
	return &UploadStatusService_Invoke_Call{Call: _e.mock.On("Invoke")}
}

func (_c *UploadStatusService_Invoke_Call) Run(run func()) *UploadStatusService_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UploadStatusService_Invoke_Call) Return(_a0 domain.UploadStatusDTO, _a1 error) *UploadStatusService_Invoke_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UploadStatusService_Invoke_Call) RunAndReturn(run func() (domain.UploadStatusDTO, error)) *UploadStatusService_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUploadStatusService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUploadStatusService creates a new instance of UploadStatusService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUploadStatusService(t mockConstructorTestingTNewUploadStatusService) *UploadStatusService {
	mock := &UploadStatusService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}