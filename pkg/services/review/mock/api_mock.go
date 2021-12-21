// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/review/api_grpc.pb.go

// Package mock_review_service is a generated GoMock package.
package mock_review_service

import (
	context "context"
	reflect "reflect"
	review_service "snakealive/m/pkg/services/review"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockReviewServiceClient is a mock of ReviewServiceClient interface.
type MockReviewServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockReviewServiceClientMockRecorder
}

// MockReviewServiceClientMockRecorder is the mock recorder for MockReviewServiceClient.
type MockReviewServiceClientMockRecorder struct {
	mock *MockReviewServiceClient
}

// NewMockReviewServiceClient creates a new mock instance.
func NewMockReviewServiceClient(ctrl *gomock.Controller) *MockReviewServiceClient {
	mock := &MockReviewServiceClient{ctrl: ctrl}
	mock.recorder = &MockReviewServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewServiceClient) EXPECT() *MockReviewServiceClientMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockReviewServiceClient) Add(ctx context.Context, in *review_service.AddReviewRequest, opts ...grpc.CallOption) (*review_service.Review, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*review_service.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockReviewServiceClientMockRecorder) Add(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockReviewServiceClient)(nil).Add), varargs...)
}

// Delete mocks base method.
func (m *MockReviewServiceClient) Delete(ctx context.Context, in *review_service.DeleteReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockReviewServiceClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReviewServiceClient)(nil).Delete), varargs...)
}

// ReviewByPlace mocks base method.
func (m *MockReviewServiceClient) ReviewByPlace(ctx context.Context, in *review_service.ReviewRequest, opts ...grpc.CallOption) (*review_service.Reviews, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReviewByPlace", varargs...)
	ret0, _ := ret[0].(*review_service.Reviews)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewByPlace indicates an expected call of ReviewByPlace.
func (mr *MockReviewServiceClientMockRecorder) ReviewByPlace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewByPlace", reflect.TypeOf((*MockReviewServiceClient)(nil).ReviewByPlace), varargs...)
}

// MockReviewServiceServer is a mock of ReviewServiceServer interface.
type MockReviewServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockReviewServiceServerMockRecorder
}

// MockReviewServiceServerMockRecorder is the mock recorder for MockReviewServiceServer.
type MockReviewServiceServerMockRecorder struct {
	mock *MockReviewServiceServer
}

// NewMockReviewServiceServer creates a new mock instance.
func NewMockReviewServiceServer(ctrl *gomock.Controller) *MockReviewServiceServer {
	mock := &MockReviewServiceServer{ctrl: ctrl}
	mock.recorder = &MockReviewServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewServiceServer) EXPECT() *MockReviewServiceServerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockReviewServiceServer) Add(arg0 context.Context, arg1 *review_service.AddReviewRequest) (*review_service.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(*review_service.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockReviewServiceServerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockReviewServiceServer)(nil).Add), arg0, arg1)
}

// Delete mocks base method.
func (m *MockReviewServiceServer) Delete(arg0 context.Context, arg1 *review_service.DeleteReviewRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockReviewServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReviewServiceServer)(nil).Delete), arg0, arg1)
}

// ReviewByPlace mocks base method.
func (m *MockReviewServiceServer) ReviewByPlace(arg0 context.Context, arg1 *review_service.ReviewRequest) (*review_service.Reviews, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReviewByPlace", arg0, arg1)
	ret0, _ := ret[0].(*review_service.Reviews)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewByPlace indicates an expected call of ReviewByPlace.
func (mr *MockReviewServiceServerMockRecorder) ReviewByPlace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewByPlace", reflect.TypeOf((*MockReviewServiceServer)(nil).ReviewByPlace), arg0, arg1)
}

// mustEmbedUnimplementedReviewServiceServer mocks base method.
func (m *MockReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedReviewServiceServer")
}

// mustEmbedUnimplementedReviewServiceServer indicates an expected call of mustEmbedUnimplementedReviewServiceServer.
func (mr *MockReviewServiceServerMockRecorder) mustEmbedUnimplementedReviewServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedReviewServiceServer", reflect.TypeOf((*MockReviewServiceServer)(nil).mustEmbedUnimplementedReviewServiceServer))
}

// MockUnsafeReviewServiceServer is a mock of UnsafeReviewServiceServer interface.
type MockUnsafeReviewServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeReviewServiceServerMockRecorder
}

// MockUnsafeReviewServiceServerMockRecorder is the mock recorder for MockUnsafeReviewServiceServer.
type MockUnsafeReviewServiceServerMockRecorder struct {
	mock *MockUnsafeReviewServiceServer
}

// NewMockUnsafeReviewServiceServer creates a new mock instance.
func NewMockUnsafeReviewServiceServer(ctrl *gomock.Controller) *MockUnsafeReviewServiceServer {
	mock := &MockUnsafeReviewServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeReviewServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeReviewServiceServer) EXPECT() *MockUnsafeReviewServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedReviewServiceServer mocks base method.
func (m *MockUnsafeReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedReviewServiceServer")
}

// mustEmbedUnimplementedReviewServiceServer indicates an expected call of mustEmbedUnimplementedReviewServiceServer.
func (mr *MockUnsafeReviewServiceServerMockRecorder) mustEmbedUnimplementedReviewServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedReviewServiceServer", reflect.TypeOf((*MockUnsafeReviewServiceServer)(nil).mustEmbedUnimplementedReviewServiceServer))
}
