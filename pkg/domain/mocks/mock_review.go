// Code generated by MockGen. DO NOT EDIT.
// Source: review.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "snakealive/m/pkg/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockReviewStorage is a mock of ReviewStorage interface.
type MockReviewStorage struct {
	ctrl     *gomock.Controller
	recorder *MockReviewStorageMockRecorder
}

// MockReviewStorageMockRecorder is the mock recorder for MockReviewStorage.
type MockReviewStorageMockRecorder struct {
	mock *MockReviewStorage
}

// NewMockReviewStorage creates a new mock instance.
func NewMockReviewStorage(ctrl *gomock.Controller) *MockReviewStorage {
	mock := &MockReviewStorage{ctrl: ctrl}
	mock.recorder = &MockReviewStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewStorage) EXPECT() *MockReviewStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockReviewStorage) Add(value domain.Review, userId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", value, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockReviewStorageMockRecorder) Add(value, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockReviewStorage)(nil).Add), value, userId)
}

// Delete mocks base method.
func (m *MockReviewStorage) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockReviewStorageMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReviewStorage)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockReviewStorage) Get(id int) (domain.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(domain.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReviewStorageMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReviewStorage)(nil).Get), id)
}

// GetListByPlace mocks base method.
func (m *MockReviewStorage) GetListByPlace(id, limit, skip int) (domain.ReviewsNoPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListByPlace", id, limit, skip)
	ret0, _ := ret[0].(domain.ReviewsNoPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListByPlace indicates an expected call of GetListByPlace.
func (mr *MockReviewStorageMockRecorder) GetListByPlace(id, limit, skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListByPlace", reflect.TypeOf((*MockReviewStorage)(nil).GetListByPlace), id, limit, skip)
}

// GetReviewAuthor mocks base method.
func (m *MockReviewStorage) GetReviewAuthor(id int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewAuthor", id)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetReviewAuthor indicates an expected call of GetReviewAuthor.
func (mr *MockReviewStorageMockRecorder) GetReviewAuthor(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewAuthor", reflect.TypeOf((*MockReviewStorage)(nil).GetReviewAuthor), id)
}

// MockReviewUseCase is a mock of ReviewUseCase interface.
type MockReviewUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockReviewUseCaseMockRecorder
}

// MockReviewUseCaseMockRecorder is the mock recorder for MockReviewUseCase.
type MockReviewUseCaseMockRecorder struct {
	mock *MockReviewUseCase
}

// NewMockReviewUseCase creates a new mock instance.
func NewMockReviewUseCase(ctrl *gomock.Controller) *MockReviewUseCase {
	mock := &MockReviewUseCase{ctrl: ctrl}
	mock.recorder = &MockReviewUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewUseCase) EXPECT() *MockReviewUseCaseMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockReviewUseCase) Add(review domain.Review, user domain.User) (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", review, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add.
func (mr *MockReviewUseCaseMockRecorder) Add(review, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockReviewUseCase)(nil).Add), review, user)
}

// CheckAuthor mocks base method.
func (m *MockReviewUseCase) CheckAuthor(user domain.User, id int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuthor", user, id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckAuthor indicates an expected call of CheckAuthor.
func (mr *MockReviewUseCaseMockRecorder) CheckAuthor(user, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuthor", reflect.TypeOf((*MockReviewUseCase)(nil).CheckAuthor), user, id)
}

// Delete mocks base method.
func (m *MockReviewUseCase) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockReviewUseCaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReviewUseCase)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockReviewUseCase) Get(id int) (domain.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(domain.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReviewUseCaseMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReviewUseCase)(nil).Get), id)
}

// GetReviewsListByPlaceId mocks base method.
func (m *MockReviewUseCase) GetReviewsListByPlaceId(id int, user domain.User, limit, skip int) (int, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewsListByPlaceId", id, user, limit, skip)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// GetReviewsListByPlaceId indicates an expected call of GetReviewsListByPlaceId.
func (mr *MockReviewUseCaseMockRecorder) GetReviewsListByPlaceId(id, user, limit, skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewsListByPlaceId", reflect.TypeOf((*MockReviewUseCase)(nil).GetReviewsListByPlaceId), id, user, limit, skip)
}

// SanitizeReview mocks base method.
func (m *MockReviewUseCase) SanitizeReview(review domain.Review) domain.Review {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SanitizeReview", review)
	ret0, _ := ret[0].(domain.Review)
	return ret0
}

// SanitizeReview indicates an expected call of SanitizeReview.
func (mr *MockReviewUseCaseMockRecorder) SanitizeReview(review interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SanitizeReview", reflect.TypeOf((*MockReviewUseCase)(nil).SanitizeReview), review)
}
