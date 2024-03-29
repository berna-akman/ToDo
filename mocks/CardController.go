// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	controller "to-do-api/pg/application/api/controller"
	card "to-do-api/pg/domain/card"

	mock "github.com/stretchr/testify/mock"
)

// CardController is an autogenerated mock type for the CardController type
type CardController struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: req, ctx
func (_m *CardController) GetAll(req *card.GetCardRequest, ctx controller.CardContext) (*card.Card, error) {
	ret := _m.Called(req, ctx)

	var r0 *card.Card
	var r1 error
	if rf, ok := ret.Get(0).(func(*card.GetCardRequest, controller.CardContext) (*card.Card, error)); ok {
		return rf(req, ctx)
	}
	if rf, ok := ret.Get(0).(func(*card.GetCardRequest, controller.CardContext) *card.Card); ok {
		r0 = rf(req, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*card.Card)
		}
	}

	if rf, ok := ret.Get(1).(func(*card.GetCardRequest, controller.CardContext) error); ok {
		r1 = rf(req, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCardController interface {
	mock.TestingT
	Cleanup(func())
}

// NewCardController creates a new instance of CardController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCardController(t mockConstructorTestingTNewCardController) *CardController {
	mock := &CardController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
