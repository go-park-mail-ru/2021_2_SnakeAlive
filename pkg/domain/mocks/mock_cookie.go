// Code generated by MockGen. DO NOT EDIT.
// Source: cookie.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "snakealive/m/pkg/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockCookieStorage is a mock of CookieStorage interface.
type MockCookieStorage struct {
	ctrl     *gomock.Controller
	recorder *MockCookieStorageMockRecorder
}

// MockCookieStorageMockRecorder is the mock recorder for MockCookieStorage.
type MockCookieStorageMockRecorder struct {
	mock *MockCookieStorage
}

// NewMockCookieStorage creates a new mock instance.
func NewMockCookieStorage(ctrl *gomock.Controller) *MockCookieStorage {
	mock := &MockCookieStorage{ctrl: ctrl}
	mock.recorder = &MockCookieStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCookieStorage) EXPECT() *MockCookieStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCookieStorage) Add(key string, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", key, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCookieStorageMockRecorder) Add(key, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCookieStorage)(nil).Add), key, userId)
}

// Delete mocks base method.
func (m *MockCookieStorage) Delete(value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCookieStorageMockRecorder) Delete(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCookieStorage)(nil).Delete), value)
}

// Get mocks base method.
func (m *MockCookieStorage) Get(value string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", value)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCookieStorageMockRecorder) Get(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCookieStorage)(nil).Get), value)
}

// MockCookieUseCase is a mock of CookieUseCase interface.
type MockCookieUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCookieUseCaseMockRecorder
}

// MockCookieUseCaseMockRecorder is the mock recorder for MockCookieUseCase.
type MockCookieUseCaseMockRecorder struct {
	mock *MockCookieUseCase
}

// NewMockCookieUseCase creates a new mock instance.
func NewMockCookieUseCase(ctrl *gomock.Controller) *MockCookieUseCase {
	mock := &MockCookieUseCase{ctrl: ctrl}
	mock.recorder = &MockCookieUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCookieUseCase) EXPECT() *MockCookieUseCaseMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCookieUseCase) Add(key string, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", key, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCookieUseCaseMockRecorder) Add(key, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCookieUseCase)(nil).Add), key, userId)
}

// Delete mocks base method.
func (m *MockCookieUseCase) Delete(value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCookieUseCaseMockRecorder) Delete(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCookieUseCase)(nil).Delete), value)
}

// Get mocks base method.
func (m *MockCookieUseCase) Get(value string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", value)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCookieUseCaseMockRecorder) Get(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCookieUseCase)(nil).Get), value)
}
