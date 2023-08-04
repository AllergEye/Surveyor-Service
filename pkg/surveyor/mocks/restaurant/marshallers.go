// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/surveyor/restaurant/marshallers.go

// Package mock_surveyor_restaurant is a generated GoMock package.
package mock_surveyor_restaurant

import (
	reflect "reflect"

	dish "github.com/allergeye/surveyor-service/internal/domain/dish"
	restaurant "github.com/allergeye/surveyor-service/internal/domain/restaurant"
	surveyor_restaurant "github.com/allergeye/surveyor-service/pkg/surveyor/restaurant"
	gomock "go.uber.org/mock/gomock"
)

// MockMarshallers is a mock of Marshallers interface.
type MockMarshallers struct {
	ctrl     *gomock.Controller
	recorder *MockMarshallersMockRecorder
}

// MockMarshallersMockRecorder is the mock recorder for MockMarshallers.
type MockMarshallersMockRecorder struct {
	mock *MockMarshallers
}

// NewMockMarshallers creates a new mock instance.
func NewMockMarshallers(ctrl *gomock.Controller) *MockMarshallers {
	mock := &MockMarshallers{ctrl: ctrl}
	mock.recorder = &MockMarshallersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarshallers) EXPECT() *MockMarshallersMockRecorder {
	return m.recorder
}

// MarshalRestaurantDishAllergenRequestBody mocks base method.
func (m *MockMarshallers) MarshalRestaurantDishAllergenRequestBody(allergensRequestBody []surveyor_restaurant.AddRestaurantDishAllergenRequestBody) ([]dish.Allergen, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalRestaurantDishAllergenRequestBody", allergensRequestBody)
	ret0, _ := ret[0].([]dish.Allergen)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalRestaurantDishAllergenRequestBody indicates an expected call of MarshalRestaurantDishAllergenRequestBody.
func (mr *MockMarshallersMockRecorder) MarshalRestaurantDishAllergenRequestBody(allergensRequestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalRestaurantDishAllergenRequestBody", reflect.TypeOf((*MockMarshallers)(nil).MarshalRestaurantDishAllergenRequestBody), allergensRequestBody)
}

// MarshalRestaurantDishRequestBody mocks base method.
func (m *MockMarshallers) MarshalRestaurantDishRequestBody(dishesRequestBody []surveyor_restaurant.AddRestaurantDishRequestBody) ([]dish.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalRestaurantDishRequestBody", dishesRequestBody)
	ret0, _ := ret[0].([]dish.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalRestaurantDishRequestBody indicates an expected call of MarshalRestaurantDishRequestBody.
func (mr *MockMarshallersMockRecorder) MarshalRestaurantDishRequestBody(dishesRequestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalRestaurantDishRequestBody", reflect.TypeOf((*MockMarshallers)(nil).MarshalRestaurantDishRequestBody), dishesRequestBody)
}

// MarshalRestaurantLocationRequestBody mocks base method.
func (m *MockMarshallers) MarshalRestaurantLocationRequestBody(locationsRequestBody []surveyor_restaurant.AddRestaurantLocationRequestBody) ([]restaurant.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalRestaurantLocationRequestBody", locationsRequestBody)
	ret0, _ := ret[0].([]restaurant.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalRestaurantLocationRequestBody indicates an expected call of MarshalRestaurantLocationRequestBody.
func (mr *MockMarshallersMockRecorder) MarshalRestaurantLocationRequestBody(locationsRequestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalRestaurantLocationRequestBody", reflect.TypeOf((*MockMarshallers)(nil).MarshalRestaurantLocationRequestBody), locationsRequestBody)
}

// MarshalRestaurantRequestBody mocks base method.
func (m *MockMarshallers) MarshalRestaurantRequestBody(restaurantRequestBody surveyor_restaurant.AddRestaurantRequestBody) (*restaurant.Restaurant, []dish.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalRestaurantRequestBody", restaurantRequestBody)
	ret0, _ := ret[0].(*restaurant.Restaurant)
	ret1, _ := ret[1].([]dish.Dish)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MarshalRestaurantRequestBody indicates an expected call of MarshalRestaurantRequestBody.
func (mr *MockMarshallersMockRecorder) MarshalRestaurantRequestBody(restaurantRequestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalRestaurantRequestBody", reflect.TypeOf((*MockMarshallers)(nil).MarshalRestaurantRequestBody), restaurantRequestBody)
}
