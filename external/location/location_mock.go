// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package location

import (
	"context"

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

// GetLocationName provides a mock function for the type MockServiceI
func (_mock *MockServiceI) GetLocationName(ctx context.Context, lat float64, lon float64) (string, error) {
	ret := _mock.Called(ctx, lat, lon)

	if len(ret) == 0 {
		panic("no return value specified for GetLocationName")
	}

	var r0 string
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, float64, float64) (string, error)); ok {
		return returnFunc(ctx, lat, lon)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, float64, float64) string); ok {
		r0 = returnFunc(ctx, lat, lon)
	} else {
		r0 = ret.Get(0).(string)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, float64, float64) error); ok {
		r1 = returnFunc(ctx, lat, lon)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockServiceI_GetLocationName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLocationName'
type MockServiceI_GetLocationName_Call struct {
	*mock.Call
}

// GetLocationName is a helper method to define mock.On call
//   - ctx
//   - lat
//   - lon
func (_e *MockServiceI_Expecter) GetLocationName(ctx interface{}, lat interface{}, lon interface{}) *MockServiceI_GetLocationName_Call {
	return &MockServiceI_GetLocationName_Call{Call: _e.mock.On("GetLocationName", ctx, lat, lon)}
}

func (_c *MockServiceI_GetLocationName_Call) Run(run func(ctx context.Context, lat float64, lon float64)) *MockServiceI_GetLocationName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(float64), args[2].(float64))
	})
	return _c
}

func (_c *MockServiceI_GetLocationName_Call) Return(s string, err error) *MockServiceI_GetLocationName_Call {
	_c.Call.Return(s, err)
	return _c
}

func (_c *MockServiceI_GetLocationName_Call) RunAndReturn(run func(ctx context.Context, lat float64, lon float64) (string, error)) *MockServiceI_GetLocationName_Call {
	_c.Call.Return(run)
	return _c
}
