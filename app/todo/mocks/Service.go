// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	domains "technical-test-be-abc/domains"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: _a0, userId
func (_m *Service) CreateTodo(_a0 domains.TodoReq, userId string) error {
	ret := _m.Called(_a0, userId)

	if len(ret) == 0 {
		panic("no return value specified for CreateTodo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domains.TodoReq, string) error); ok {
		r0 = rf(_a0, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTodo provides a mock function with given fields: id, userId
func (_m *Service) DeleteTodo(id string, userId string) error {
	ret := _m.Called(id, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTodo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTodo provides a mock function with given fields: UserId
func (_m *Service) GetAllTodo(UserId string) ([]domains.Todo, error) {
	ret := _m.Called(UserId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTodo")
	}

	var r0 []domains.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domains.Todo, error)); ok {
		return rf(UserId)
	}
	if rf, ok := ret.Get(0).(func(string) []domains.Todo); ok {
		r0 = rf(UserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domains.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(UserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoById provides a mock function with given fields: id, userId
func (_m *Service) GetTodoById(id string, userId string) (domains.Todo, error) {
	ret := _m.Called(id, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetTodoById")
	}

	var r0 domains.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (domains.Todo, error)); ok {
		return rf(id, userId)
	}
	if rf, ok := ret.Get(0).(func(string, string) domains.Todo); ok {
		r0 = rf(id, userId)
	} else {
		r0 = ret.Get(0).(domains.Todo)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusTodo provides a mock function with given fields: id, status, userId
func (_m *Service) UpdateStatusTodo(id string, status bool, userId string) error {
	ret := _m.Called(id, status, userId)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatusTodo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool, string) error); ok {
		r0 = rf(id, status, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTodoById provides a mock function with given fields: id, userId, todoReq
func (_m *Service) UpdateTodoById(id string, userId string, todoReq domains.TodoReq) error {
	ret := _m.Called(id, userId, todoReq)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTodoById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, domains.TodoReq) error); ok {
		r0 = rf(id, userId, todoReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
