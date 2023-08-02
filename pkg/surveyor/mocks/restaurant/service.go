// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/surveyor/restaurant/service.go

// Package mock_surveyor is a generated GoMock package.
package mock_surveyor

import (
	context "context"
	reflect "reflect"

	restaurant "github.com/allergeye/surveyor-service/internal/domain/restaurant"
	gomock "go.uber.org/mock/gomock"
)

// MockRestaurantService is a mock of RestaurantService interface.
type MockRestaurantService struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantServiceMockRecorder
}

// MockRestaurantServiceMockRecorder is the mock recorder for MockRestaurantService.
type MockRestaurantServiceMockRecorder struct {
	mock *MockRestaurantService
}

// NewMockRestaurantService creates a new mock instance.
func NewMockRestaurantService(ctrl *gomock.Controller) *MockRestaurantService {
	mock := &MockRestaurantService{ctrl: ctrl}
	mock.recorder = &MockRestaurantServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantService) EXPECT() *MockRestaurantServiceMockRecorder {
	return m.recorder
}

// AddRestaurant mocks base method.
func (m *MockRestaurantService) AddRestaurant(ctx context.Context, restaurant restaurant.Restaurant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRestaurant", ctx, restaurant)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRestaurant indicates an expected call of AddRestaurant.
func (mr *MockRestaurantServiceMockRecorder) AddRestaurant(ctx, restaurant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRestaurant", reflect.TypeOf((*MockRestaurantService)(nil).AddRestaurant), ctx, restaurant)
}

// GetAllRestaurants mocks base method.
func (m *MockRestaurantService) GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRestaurants", ctx)
	ret0, _ := ret[0].([]restaurant.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRestaurants indicates an expected call of GetAllRestaurants.
func (mr *MockRestaurantServiceMockRecorder) GetAllRestaurants(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRestaurants", reflect.TypeOf((*MockRestaurantService)(nil).GetAllRestaurants), ctx)
}
