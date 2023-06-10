// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "graph-analyzer/api/upload/domain"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// UploadService is an autogenerated mock type for the UploadService type
type UploadService struct {
	mock.Mock
}

type UploadService_Expecter struct {
	mock *mock.Mock
}

func (_m *UploadService) EXPECT() *UploadService_Expecter {
	return &UploadService_Expecter{mock: &_m.Mock}
}

// Invoke provides a mock function with given fields: file
func (_m *UploadService) Invoke(file *multipart.FileHeader) (domain.UploadDTO, error) {
	ret := _m.Called(file)

	var r0 domain.UploadDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader) (domain.UploadDTO, error)); ok {
		return rf(file)
	}
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader) domain.UploadDTO); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(domain.UploadDTO)
	}

	if rf, ok := ret.Get(1).(func(*multipart.FileHeader) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadService_Invoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Invoke'
type UploadService_Invoke_Call struct {
	*mock.Call
}

// Invoke is a helper method to define mock.On call
//   - file *multipart.FileHeader
func (_e *UploadService_Expecter) Invoke(file interface{}) *UploadService_Invoke_Call {
	return &UploadService_Invoke_Call{Call: _e.mock.On("Invoke", file)}
}

func (_c *UploadService_Invoke_Call) Run(run func(file *multipart.FileHeader)) *UploadService_Invoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*multipart.FileHeader))
	})
	return _c
}

func (_c *UploadService_Invoke_Call) Return(_a0 domain.UploadDTO, _a1 error) *UploadService_Invoke_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UploadService_Invoke_Call) RunAndReturn(run func(*multipart.FileHeader) (domain.UploadDTO, error)) *UploadService_Invoke_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUploadService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUploadService creates a new instance of UploadService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUploadService(t mockConstructorTestingTNewUploadService) *UploadService {
	mock := &UploadService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}