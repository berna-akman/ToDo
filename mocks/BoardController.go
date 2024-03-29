// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// BoardController is an autogenerated mock type for the BoardController type
type BoardController struct {
	mock.Mock
}

// CreateBoard provides a mock function with given fields: e
func (_m *BoardController) CreateBoard(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBoard provides a mock function with given fields: e
func (_m *BoardController) DeleteBoard(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBoards provides a mock function with given fields: e
func (_m *BoardController) GetAllBoards(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBoardByID provides a mock function with given fields: e
func (_m *BoardController) GetBoardByID(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBoard provides a mock function with given fields: e
func (_m *BoardController) UpdateBoard(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBoardController interface {
	mock.TestingT
	Cleanup(func())
}

// NewBoardController creates a new instance of BoardController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBoardController(t mockConstructorTestingTNewBoardController) *BoardController {
	mock := &BoardController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
