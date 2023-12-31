// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/lib/helpers.go

// Package mock_lib is a generated GoMock package.
package mock_lib

import (
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockHelpers is a mock of Helpers interface.
type MockHelpers struct {
	ctrl     *gomock.Controller
	recorder *MockHelpersMockRecorder
}

// MockHelpersMockRecorder is the mock recorder for MockHelpers.
type MockHelpersMockRecorder struct {
	mock *MockHelpers
}

// NewMockHelpers creates a new mock instance.
func NewMockHelpers(ctrl *gomock.Controller) *MockHelpers {
	mock := &MockHelpers{ctrl: ctrl}
	mock.recorder = &MockHelpersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelpers) EXPECT() *MockHelpersMockRecorder {
	return m.recorder
}

// GenerateUUID mocks base method.
func (m *MockHelpers) GenerateUUID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUUID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GenerateUUID indicates an expected call of GenerateUUID.
func (mr *MockHelpersMockRecorder) GenerateUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUUID", reflect.TypeOf((*MockHelpers)(nil).GenerateUUID))
}
