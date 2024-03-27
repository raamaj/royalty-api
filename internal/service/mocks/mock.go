// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mocks/mock.go
//
// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	domain "royalty-api/internal/domain"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserService) Delete(id uint) (*domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(*domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserServiceMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserService)(nil).Delete), id)
}

// Insert mocks base method.
func (m *MockUserService) Insert(input *domain.UserRequest) (*domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", input)
	ret0, _ := ret[0].(*domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockUserServiceMockRecorder) Insert(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserService)(nil).Insert), input)
}

// List mocks base method.
func (m *MockUserService) List() (*[]domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(*[]domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserService)(nil).List))
}

// Update mocks base method.
func (m *MockUserService) Update(input *domain.UserRequest) (*domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(*domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserServiceMockRecorder) Update(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserService)(nil).Update), input)
}

// View mocks base method.
func (m *MockUserService) View(id uint) (*domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "View", id)
	ret0, _ := ret[0].(*domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// View indicates an expected call of View.
func (mr *MockUserServiceMockRecorder) View(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "View", reflect.TypeOf((*MockUserService)(nil).View), id)
}

// MockVoucherService is a mock of VoucherService interface.
type MockVoucherService struct {
	ctrl     *gomock.Controller
	recorder *MockVoucherServiceMockRecorder
}

// MockVoucherServiceMockRecorder is the mock recorder for MockVoucherService.
type MockVoucherServiceMockRecorder struct {
	mock *MockVoucherService
}

// NewMockVoucherService creates a new mock instance.
func NewMockVoucherService(ctrl *gomock.Controller) *MockVoucherService {
	mock := &MockVoucherService{ctrl: ctrl}
	mock.recorder = &MockVoucherServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVoucherService) EXPECT() *MockVoucherServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVoucherService) Create(invoice string, timeNow time.Time) (*domain.VoucherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", invoice, timeNow)
	ret0, _ := ret[0].(*domain.VoucherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVoucherServiceMockRecorder) Create(invoice, timeNow any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVoucherService)(nil).Create), invoice, timeNow)
}

// List mocks base method.
func (m *MockVoucherService) List(userID uint) (*[]domain.VoucherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", userID)
	ret0, _ := ret[0].(*[]domain.VoucherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVoucherServiceMockRecorder) List(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVoucherService)(nil).List), userID)
}

// Redeem mocks base method.
func (m *MockVoucherService) Redeem(code string, timeNow time.Time) (*domain.VoucherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Redeem", code, timeNow)
	ret0, _ := ret[0].(*domain.VoucherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Redeem indicates an expected call of Redeem.
func (mr *MockVoucherServiceMockRecorder) Redeem(code, timeNow any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redeem", reflect.TypeOf((*MockVoucherService)(nil).Redeem), code, timeNow)
}

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransactionService) Create(request *domain.TransactionRequest, timeNow time.Time) (*domain.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", request, timeNow)
	ret0, _ := ret[0].(*domain.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTransactionServiceMockRecorder) Create(request, timeNow any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionService)(nil).Create), request, timeNow)
}

// List mocks base method.
func (m *MockTransactionService) List(userId uint) (*[]domain.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", userId)
	ret0, _ := ret[0].(*[]domain.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTransactionServiceMockRecorder) List(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTransactionService)(nil).List), userId)
}
