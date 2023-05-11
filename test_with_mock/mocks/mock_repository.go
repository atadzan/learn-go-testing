// Code generated by MockGen. DO NOT EDIT.
// Source: test_with_mock/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	models "github.com/atadzan/learn-go-testing/test_with_mock/models"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
)

// MockOrderRepo is a mock of OrderRepo interface.
type MockOrderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepoMockRecorder
}

// MockOrderRepoMockRecorder is the mock recorder for MockOrderRepo.
type MockOrderRepoMockRecorder struct {
	mock *MockOrderRepo
}

// NewMockOrderRepo creates a new mock instance.
func NewMockOrderRepo(ctrl *gomock.Controller) *MockOrderRepo {
	mock := &MockOrderRepo{ctrl: ctrl}
	mock.recorder = &MockOrderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepo) EXPECT() *MockOrderRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockOrderRepo) Get(ctx context.Context, log *logrus.Logger, ids []uint) (map[uint]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, log, ids)
	ret0, _ := ret[0].(map[uint]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrderRepoMockRecorder) Get(ctx, log, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrderRepo)(nil).Get), ctx, log, ids)
}

// Save mocks base method.
func (m *MockOrderRepo) Save(ctx context.Context, log *logrus.Logger, order *models.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, log, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOrderRepoMockRecorder) Save(ctx, log, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOrderRepo)(nil).Save), ctx, log, order)
}