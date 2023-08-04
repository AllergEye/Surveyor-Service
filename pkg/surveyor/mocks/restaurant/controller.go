// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/surveyor/restaurant/controller.go

// Package mock_surveyor_restaurant is a generated GoMock package.
package mock_surveyor_restaurant

import (
	context "context"
	reflect "reflect"

	restaurant "github.com/allergeye/surveyor-service/internal/domain/restaurant"
	surveyor_restaurant "github.com/allergeye/surveyor-service/pkg/surveyor/restaurant"
	gomock "go.uber.org/mock/gomock"
)

// MockRestaurantController is a mock of RestaurantController interface.
type MockRestaurantController struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantControllerMockRecorder
}

// MockRestaurantControllerMockRecorder is the mock recorder for MockRestaurantController.
type MockRestaurantControllerMockRecorder struct {
	mock *MockRestaurantController
}

// NewMockRestaurantController creates a new mock instance.
func NewMockRestaurantController(ctrl *gomock.Controller) *MockRestaurantController {
	mock := &MockRestaurantController{ctrl: ctrl}
	mock.recorder = &MockRestaurantControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantController) EXPECT() *MockRestaurantControllerMockRecorder {
	return m.recorder
}

// AddRestaurant mocks base method.
func (m *MockRestaurantController) AddRestaurant(ctx context.Context, requestBody surveyor_restaurant.AddRestaurantRequestBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRestaurant", ctx, requestBody)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRestaurant indicates an expected call of AddRestaurant.
func (mr *MockRestaurantControllerMockRecorder) AddRestaurant(ctx, requestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRestaurant", reflect.TypeOf((*MockRestaurantController)(nil).AddRestaurant), ctx, requestBody)
}

// GetAllRestaurants mocks base method.
func (m *MockRestaurantController) GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRestaurants", ctx)
	ret0, _ := ret[0].([]restaurant.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRestaurants indicates an expected call of GetAllRestaurants.
func (mr *MockRestaurantControllerMockRecorder) GetAllRestaurants(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRestaurants", reflect.TypeOf((*MockRestaurantController)(nil).GetAllRestaurants), ctx)
}
