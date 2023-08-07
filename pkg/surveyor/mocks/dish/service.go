// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/surveyor/dish/service.go

// Package mock_surveyor_dish is a generated GoMock package.
package mock_surveyor_dish

import (
	context "context"
	reflect "reflect"

	dish "github.com/allergeye/surveyor-service/internal/domain/dish"
	gomock "go.uber.org/mock/gomock"
)

// MockDishService is a mock of DishService interface.
type MockDishService struct {
	ctrl     *gomock.Controller
	recorder *MockDishServiceMockRecorder
}

// MockDishServiceMockRecorder is the mock recorder for MockDishService.
type MockDishServiceMockRecorder struct {
	mock *MockDishService
}

// NewMockDishService creates a new mock instance.
func NewMockDishService(ctrl *gomock.Controller) *MockDishService {
	mock := &MockDishService{ctrl: ctrl}
	mock.recorder = &MockDishServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDishService) EXPECT() *MockDishServiceMockRecorder {
	return m.recorder
}

// AddDishesToRestaurant mocks base method.
func (m *MockDishService) AddDishesToRestaurant(ctx context.Context, restaurantId string, dishes []dish.Dish) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDishesToRestaurant", ctx, restaurantId, dishes)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDishesToRestaurant indicates an expected call of AddDishesToRestaurant.
func (mr *MockDishServiceMockRecorder) AddDishesToRestaurant(ctx, restaurantId, dishes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDishesToRestaurant", reflect.TypeOf((*MockDishService)(nil).AddDishesToRestaurant), ctx, restaurantId, dishes)
}

// GetDishesByRestaurantId mocks base method.
func (m *MockDishService) GetDishesByRestaurantId(ctx context.Context, restaurantId string) ([]dish.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDishesByRestaurantId", ctx, restaurantId)
	ret0, _ := ret[0].([]dish.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDishesByRestaurantId indicates an expected call of GetDishesByRestaurantId.
func (mr *MockDishServiceMockRecorder) GetDishesByRestaurantId(ctx, restaurantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDishesByRestaurantId", reflect.TypeOf((*MockDishService)(nil).GetDishesByRestaurantId), ctx, restaurantId)
}
