// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	models "snakealive/m/internal/services/trip/models"

	gomock "github.com/golang/mock/gomock"
)

// MockTripRepository is a mock of TripRepository interface.
type MockTripRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTripRepositoryMockRecorder
}

// MockTripRepositoryMockRecorder is the mock recorder for MockTripRepository.
type MockTripRepositoryMockRecorder struct {
	mock *MockTripRepository
}

// NewMockTripRepository creates a new mock instance.
func NewMockTripRepository(ctrl *gomock.Controller) *MockTripRepository {
	mock := &MockTripRepository{ctrl: ctrl}
	mock.recorder = &MockTripRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTripRepository) EXPECT() *MockTripRepositoryMockRecorder {
	return m.recorder
}

// AddAlbum mocks base method.
func (m *MockTripRepository) AddAlbum(ctx context.Context, album *models.Album, userID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAlbum", ctx, album, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAlbum indicates an expected call of AddAlbum.
func (mr *MockTripRepositoryMockRecorder) AddAlbum(ctx, album, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlbum", reflect.TypeOf((*MockTripRepository)(nil).AddAlbum), ctx, album, userID)
}

// AddFilename mocks base method.
func (m *MockTripRepository) AddFilename(ctx context.Context, filename string, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilename", ctx, filename, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilename indicates an expected call of AddFilename.
func (mr *MockTripRepositoryMockRecorder) AddFilename(ctx, filename, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilename", reflect.TypeOf((*MockTripRepository)(nil).AddFilename), ctx, filename, id)
}

// AddTrip mocks base method.
func (m *MockTripRepository) AddTrip(ctx context.Context, value *models.Trip, userID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrip", ctx, value, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTrip indicates an expected call of AddTrip.
func (mr *MockTripRepositoryMockRecorder) AddTrip(ctx, value, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrip", reflect.TypeOf((*MockTripRepository)(nil).AddTrip), ctx, value, userID)
}

// DeleteAlbum mocks base method.
func (m *MockTripRepository) DeleteAlbum(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlbum", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlbum indicates an expected call of DeleteAlbum.
func (mr *MockTripRepositoryMockRecorder) DeleteAlbum(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlbum", reflect.TypeOf((*MockTripRepository)(nil).DeleteAlbum), ctx, id)
}

// DeleteTrip mocks base method.
func (m *MockTripRepository) DeleteTrip(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrip", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrip indicates an expected call of DeleteTrip.
func (mr *MockTripRepositoryMockRecorder) DeleteTrip(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrip", reflect.TypeOf((*MockTripRepository)(nil).DeleteTrip), ctx, id)
}

// GetAlbumAuthor mocks base method.
func (m *MockTripRepository) GetAlbumAuthor(ctx context.Context, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumAuthor", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbumAuthor indicates an expected call of GetAlbumAuthor.
func (mr *MockTripRepositoryMockRecorder) GetAlbumAuthor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumAuthor", reflect.TypeOf((*MockTripRepository)(nil).GetAlbumAuthor), ctx, id)
}

// GetAlbumById mocks base method.
func (m *MockTripRepository) GetAlbumById(ctx context.Context, id int) (*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumById", ctx, id)
	ret0, _ := ret[0].(*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbumById indicates an expected call of GetAlbumById.
func (mr *MockTripRepositoryMockRecorder) GetAlbumById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumById", reflect.TypeOf((*MockTripRepository)(nil).GetAlbumById), ctx, id)
}

// GetTripAuthor mocks base method.
func (m *MockTripRepository) GetTripAuthor(ctx context.Context, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTripAuthor", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTripAuthor indicates an expected call of GetTripAuthor.
func (mr *MockTripRepositoryMockRecorder) GetTripAuthor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripAuthor", reflect.TypeOf((*MockTripRepository)(nil).GetTripAuthor), ctx, id)
}

// GetTripById mocks base method.
func (m *MockTripRepository) GetTripById(ctx context.Context, id int) (*models.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTripById", ctx, id)
	ret0, _ := ret[0].(*models.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTripById indicates an expected call of GetTripById.
func (mr *MockTripRepositoryMockRecorder) GetTripById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripById", reflect.TypeOf((*MockTripRepository)(nil).GetTripById), ctx, id)
}

// SightsByTrip mocks base method.
func (m *MockTripRepository) SightsByTrip(ctx context.Context, id int) (*[]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SightsByTrip", ctx, id)
	ret0, _ := ret[0].(*[]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SightsByTrip indicates an expected call of SightsByTrip.
func (mr *MockTripRepositoryMockRecorder) SightsByTrip(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SightsByTrip", reflect.TypeOf((*MockTripRepository)(nil).SightsByTrip), ctx, id)
}

// UpdateAlbum mocks base method.
func (m *MockTripRepository) UpdateAlbum(ctx context.Context, id int, album *models.Album) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlbum", ctx, id, album)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAlbum indicates an expected call of UpdateAlbum.
func (mr *MockTripRepositoryMockRecorder) UpdateAlbum(ctx, id, album interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlbum", reflect.TypeOf((*MockTripRepository)(nil).UpdateAlbum), ctx, id, album)
}

// UpdateTrip mocks base method.
func (m *MockTripRepository) UpdateTrip(ctx context.Context, id int, value *models.Trip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrip", ctx, id, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrip indicates an expected call of UpdateTrip.
func (mr *MockTripRepositoryMockRecorder) UpdateTrip(ctx, id, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrip", reflect.TypeOf((*MockTripRepository)(nil).UpdateTrip), ctx, id, value)
}