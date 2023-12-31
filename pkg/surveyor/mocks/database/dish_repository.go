// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/database/dish_repository.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	dish "github.com/allergeye/surveyor-service/internal/domain/dish"
	gomock "go.uber.org/mock/gomock"
)

// MockDishRepository is a mock of DishRepository interface.
type MockDishRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDishRepositoryMockRecorder
}

// MockDishRepositoryMockRecorder is the mock recorder for MockDishRepository.
type MockDishRepositoryMockRecorder struct {
	mock *MockDishRepository
}

// NewMockDishRepository creates a new mock instance.
func NewMockDishRepository(ctrl *gomock.Controller) *MockDishRepository {
	mock := &MockDishRepository{ctrl: ctrl}
	mock.recorder = &MockDishRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDishRepository) EXPECT() *MockDishRepositoryMockRecorder {
	return m.recorder
}

// AddDishes mocks base method.
func (m *MockDishRepository) AddDishes(ctx context.Context, dishes []dish.Dish) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDishes", ctx, dishes)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDishes indicates an expected call of AddDishes.
func (mr *MockDishRepositoryMockRecorder) AddDishes(ctx, dishes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDishes", reflect.TypeOf((*MockDishRepository)(nil).AddDishes), ctx, dishes)
}

// GetDishById mocks base method.
func (m *MockDishRepository) GetDishById(ctx context.Context, dishId string) (*dish.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDishById", ctx, dishId)
	ret0, _ := ret[0].(*dish.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDishById indicates an expected call of GetDishById.
func (mr *MockDishRepositoryMockRecorder) GetDishById(ctx, dishId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDishById", reflect.TypeOf((*MockDishRepository)(nil).GetDishById), ctx, dishId)
}
