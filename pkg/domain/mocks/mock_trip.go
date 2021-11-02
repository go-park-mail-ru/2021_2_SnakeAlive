// Code generated by MockGen. DO NOT EDIT.
// Source: trip.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "snakealive/m/pkg/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockTripStorage is a mock of TripStorage interface.
type MockTripStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTripStorageMockRecorder
}

// MockTripStorageMockRecorder is the mock recorder for MockTripStorage.
type MockTripStorageMockRecorder struct {
	mock *MockTripStorage
}

// NewMockTripStorage creates a new mock instance.
func NewMockTripStorage(ctrl *gomock.Controller) *MockTripStorage {
	mock := &MockTripStorage{ctrl: ctrl}
	mock.recorder = &MockTripStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTripStorage) EXPECT() *MockTripStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTripStorage) Add(value domain.Trip, user domain.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", value, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockTripStorageMockRecorder) Add(value, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTripStorage)(nil).Add), value, user)
}

// Delete mocks base method.
func (m *MockTripStorage) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTripStorageMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTripStorage)(nil).Delete), id)
}

// GetById mocks base method.
func (m *MockTripStorage) GetById(id int) (domain.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(domain.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTripStorageMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTripStorage)(nil).GetById), id)
}

// GetTripAuthor mocks base method.
func (m *MockTripStorage) GetTripAuthor(id int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTripAuthor", id)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetTripAuthor indicates an expected call of GetTripAuthor.
func (mr *MockTripStorageMockRecorder) GetTripAuthor(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripAuthor", reflect.TypeOf((*MockTripStorage)(nil).GetTripAuthor), id)
}

// Update mocks base method.
func (m *MockTripStorage) Update(id int, value domain.Trip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTripStorageMockRecorder) Update(id, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTripStorage)(nil).Update), id, value)
}

// MockTripUseCase is a mock of TripUseCase interface.
type MockTripUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockTripUseCaseMockRecorder
}

// MockTripUseCaseMockRecorder is the mock recorder for MockTripUseCase.
type MockTripUseCaseMockRecorder struct {
	mock *MockTripUseCase
}

// NewMockTripUseCase creates a new mock instance.
func NewMockTripUseCase(ctrl *gomock.Controller) *MockTripUseCase {
	mock := &MockTripUseCase{ctrl: ctrl}
	mock.recorder = &MockTripUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTripUseCase) EXPECT() *MockTripUseCaseMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTripUseCase) Add(value domain.Trip, user domain.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", value, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockTripUseCaseMockRecorder) Add(value, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTripUseCase)(nil).Add), value, user)
}

// CheckAuthor mocks base method.
func (m *MockTripUseCase) CheckAuthor(user domain.User, id int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuthor", user, id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckAuthor indicates an expected call of CheckAuthor.
func (mr *MockTripUseCaseMockRecorder) CheckAuthor(user, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuthor", reflect.TypeOf((*MockTripUseCase)(nil).CheckAuthor), user, id)
}

// Delete mocks base method.
func (m *MockTripUseCase) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTripUseCaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTripUseCase)(nil).Delete), id)
}

// GetById mocks base method.
func (m *MockTripUseCase) GetById(id int) (int, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTripUseCaseMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTripUseCase)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockTripUseCase) Update(id int, updatedTrip domain.Trip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, updatedTrip)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTripUseCaseMockRecorder) Update(id, updatedTrip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTripUseCase)(nil).Update), id, updatedTrip)
}
