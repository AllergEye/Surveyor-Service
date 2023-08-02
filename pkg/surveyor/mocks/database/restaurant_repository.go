// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/database/restaurant_repository.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	restaurant "github.com/allergeye/surveyor-service/internal/domain/restaurant"
	gomock "go.uber.org/mock/gomock"
)

// MockRestaurantRepository is a mock of RestaurantRepository interface.
type MockRestaurantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantRepositoryMockRecorder
}

// MockRestaurantRepositoryMockRecorder is the mock recorder for MockRestaurantRepository.
type MockRestaurantRepositoryMockRecorder struct {
	mock *MockRestaurantRepository
}

// NewMockRestaurantRepository creates a new mock instance.
func NewMockRestaurantRepository(ctrl *gomock.Controller) *MockRestaurantRepository {
	mock := &MockRestaurantRepository{ctrl: ctrl}
	mock.recorder = &MockRestaurantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantRepository) EXPECT() *MockRestaurantRepositoryMockRecorder {
	return m.recorder
}

// AddRestaurant mocks base method.
func (m *MockRestaurantRepository) AddRestaurant(ctx context.Context, restaurant restaurant.Restaurant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRestaurant", ctx, restaurant)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRestaurant indicates an expected call of AddRestaurant.
func (mr *MockRestaurantRepositoryMockRecorder) AddRestaurant(ctx, restaurant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRestaurant", reflect.TypeOf((*MockRestaurantRepository)(nil).AddRestaurant), ctx, restaurant)
}

// GetAllRestaurants mocks base method.
func (m *MockRestaurantRepository) GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRestaurants", ctx)
	ret0, _ := ret[0].([]restaurant.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRestaurants indicates an expected call of GetAllRestaurants.
func (mr *MockRestaurantRepositoryMockRecorder) GetAllRestaurants(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRestaurants", reflect.TypeOf((*MockRestaurantRepository)(nil).GetAllRestaurants), ctx)
}

// GetRestaurantByName mocks base method.
func (m *MockRestaurantRepository) GetRestaurantByName(ctx context.Context, name string) (*restaurant.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantByName", ctx, name)
	ret0, _ := ret[0].(*restaurant.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantByName indicates an expected call of GetRestaurantByName.
func (mr *MockRestaurantRepositoryMockRecorder) GetRestaurantByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantByName", reflect.TypeOf((*MockRestaurantRepository)(nil).GetRestaurantByName), ctx, name)
}

// UpdateRestaurantLocations mocks base method.
func (m *MockRestaurantRepository) UpdateRestaurantLocations(ctx context.Context, restaurant restaurant.Restaurant, locations []restaurant.Location) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRestaurantLocations", ctx, restaurant, locations)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRestaurantLocations indicates an expected call of UpdateRestaurantLocations.
func (mr *MockRestaurantRepositoryMockRecorder) UpdateRestaurantLocations(ctx, restaurant, locations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRestaurantLocations", reflect.TypeOf((*MockRestaurantRepository)(nil).UpdateRestaurantLocations), ctx, restaurant, locations)
}
