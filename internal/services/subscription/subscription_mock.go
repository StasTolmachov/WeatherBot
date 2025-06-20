// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package subscription

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	mock "github.com/stretchr/testify/mock"
)

// NewMockServiceI creates a new instance of MockServiceI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServiceI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServiceI {
	mock := &MockServiceI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockServiceI is an autogenerated mock type for the ServiceI type
type MockServiceI struct {
	mock.Mock
}

type MockServiceI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServiceI) EXPECT() *MockServiceI_Expecter {
	return &MockServiceI_Expecter{mock: &_m.Mock}
}

// CreateSubscription provides a mock function for the type MockServiceI
func (_mock *MockServiceI) CreateSubscription(ctx context.Context, userID int64, localTime string) error {
	ret := _mock.Called(ctx, userID, localTime)

	if len(ret) == 0 {
		panic("no return value specified for CreateSubscription")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64, string) error); ok {
		r0 = returnFunc(ctx, userID, localTime)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockServiceI_CreateSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSubscription'
type MockServiceI_CreateSubscription_Call struct {
	*mock.Call
}

// CreateSubscription is a helper method to define mock.On call
//   - ctx
//   - userID
//   - localTime
func (_e *MockServiceI_Expecter) CreateSubscription(ctx interface{}, userID interface{}, localTime interface{}) *MockServiceI_CreateSubscription_Call {
	return &MockServiceI_CreateSubscription_Call{Call: _e.mock.On("CreateSubscription", ctx, userID, localTime)}
}

func (_c *MockServiceI_CreateSubscription_Call) Run(run func(ctx context.Context, userID int64, localTime string)) *MockServiceI_CreateSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockServiceI_CreateSubscription_Call) Return(err error) *MockServiceI_CreateSubscription_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockServiceI_CreateSubscription_Call) RunAndReturn(run func(ctx context.Context, userID int64, localTime string) error) *MockServiceI_CreateSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSubscription provides a mock function for the type MockServiceI
func (_mock *MockServiceI) DeleteSubscription(ctx context.Context, id string) error {
	ret := _mock.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSubscription")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = returnFunc(ctx, id)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockServiceI_DeleteSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSubscription'
type MockServiceI_DeleteSubscription_Call struct {
	*mock.Call
}

// DeleteSubscription is a helper method to define mock.On call
//   - ctx
//   - id
func (_e *MockServiceI_Expecter) DeleteSubscription(ctx interface{}, id interface{}) *MockServiceI_DeleteSubscription_Call {
	return &MockServiceI_DeleteSubscription_Call{Call: _e.mock.On("DeleteSubscription", ctx, id)}
}

func (_c *MockServiceI_DeleteSubscription_Call) Run(run func(ctx context.Context, id string)) *MockServiceI_DeleteSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockServiceI_DeleteSubscription_Call) Return(err error) *MockServiceI_DeleteSubscription_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockServiceI_DeleteSubscription_Call) RunAndReturn(run func(ctx context.Context, id string) error) *MockServiceI_DeleteSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSubscriptions provides a mock function for the type MockServiceI
func (_mock *MockServiceI) DeleteSubscriptions(ctx context.Context, userID int64) error {
	ret := _mock.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSubscriptions")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = returnFunc(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockServiceI_DeleteSubscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSubscriptions'
type MockServiceI_DeleteSubscriptions_Call struct {
	*mock.Call
}

// DeleteSubscriptions is a helper method to define mock.On call
//   - ctx
//   - userID
func (_e *MockServiceI_Expecter) DeleteSubscriptions(ctx interface{}, userID interface{}) *MockServiceI_DeleteSubscriptions_Call {
	return &MockServiceI_DeleteSubscriptions_Call{Call: _e.mock.On("DeleteSubscriptions", ctx, userID)}
}

func (_c *MockServiceI_DeleteSubscriptions_Call) Run(run func(ctx context.Context, userID int64)) *MockServiceI_DeleteSubscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockServiceI_DeleteSubscriptions_Call) Return(err error) *MockServiceI_DeleteSubscriptions_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockServiceI_DeleteSubscriptions_Call) RunAndReturn(run func(ctx context.Context, userID int64) error) *MockServiceI_DeleteSubscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubsByUserAndKeyboard provides a mock function for the type MockServiceI
func (_mock *MockServiceI) GetSubsByUserAndKeyboard(ctx context.Context, userID int64) (tgbotapi.InlineKeyboardMarkup, error) {
	ret := _mock.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetSubsByUserAndKeyboard")
	}

	var r0 tgbotapi.InlineKeyboardMarkup
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) (tgbotapi.InlineKeyboardMarkup, error)); ok {
		return returnFunc(ctx, userID)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) tgbotapi.InlineKeyboardMarkup); ok {
		r0 = returnFunc(ctx, userID)
	} else {
		r0 = ret.Get(0).(tgbotapi.InlineKeyboardMarkup)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = returnFunc(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockServiceI_GetSubsByUserAndKeyboard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubsByUserAndKeyboard'
type MockServiceI_GetSubsByUserAndKeyboard_Call struct {
	*mock.Call
}

// GetSubsByUserAndKeyboard is a helper method to define mock.On call
//   - ctx
//   - userID
func (_e *MockServiceI_Expecter) GetSubsByUserAndKeyboard(ctx interface{}, userID interface{}) *MockServiceI_GetSubsByUserAndKeyboard_Call {
	return &MockServiceI_GetSubsByUserAndKeyboard_Call{Call: _e.mock.On("GetSubsByUserAndKeyboard", ctx, userID)}
}

func (_c *MockServiceI_GetSubsByUserAndKeyboard_Call) Run(run func(ctx context.Context, userID int64)) *MockServiceI_GetSubsByUserAndKeyboard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockServiceI_GetSubsByUserAndKeyboard_Call) Return(inlineKeyboardMarkup tgbotapi.InlineKeyboardMarkup, err error) *MockServiceI_GetSubsByUserAndKeyboard_Call {
	_c.Call.Return(inlineKeyboardMarkup, err)
	return _c
}

func (_c *MockServiceI_GetSubsByUserAndKeyboard_Call) RunAndReturn(run func(ctx context.Context, userID int64) (tgbotapi.InlineKeyboardMarkup, error)) *MockServiceI_GetSubsByUserAndKeyboard_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserLocation provides a mock function for the type MockServiceI
func (_mock *MockServiceI) SetUserLocation(chatID int64, lat float64, lon float64) {
	_mock.Called(chatID, lat, lon)
	return
}

// MockServiceI_SetUserLocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserLocation'
type MockServiceI_SetUserLocation_Call struct {
	*mock.Call
}

// SetUserLocation is a helper method to define mock.On call
//   - chatID
//   - lat
//   - lon
func (_e *MockServiceI_Expecter) SetUserLocation(chatID interface{}, lat interface{}, lon interface{}) *MockServiceI_SetUserLocation_Call {
	return &MockServiceI_SetUserLocation_Call{Call: _e.mock.On("SetUserLocation", chatID, lat, lon)}
}

func (_c *MockServiceI_SetUserLocation_Call) Run(run func(chatID int64, lat float64, lon float64)) *MockServiceI_SetUserLocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(float64), args[2].(float64))
	})
	return _c
}

func (_c *MockServiceI_SetUserLocation_Call) Return() *MockServiceI_SetUserLocation_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockServiceI_SetUserLocation_Call) RunAndReturn(run func(chatID int64, lat float64, lon float64)) *MockServiceI_SetUserLocation_Call {
	_c.Run(run)
	return _c
}
