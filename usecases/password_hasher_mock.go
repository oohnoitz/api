// Code generated by MockGen. DO NOT EDIT.
// Source: password_hasher.go

// Package usecases is a generated GoMock package.
package usecases

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/tadoku/api/domain"
	reflect "reflect"
)

// MockPasswordHasher is a mock of PasswordHasher interface
type MockPasswordHasher struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHasherMockRecorder
}

// MockPasswordHasherMockRecorder is the mock recorder for MockPasswordHasher
type MockPasswordHasherMockRecorder struct {
	mock *MockPasswordHasher
}

// NewMockPasswordHasher creates a new mock instance
func NewMockPasswordHasher(ctrl *gomock.Controller) *MockPasswordHasher {
	mock := &MockPasswordHasher{ctrl: ctrl}
	mock.recorder = &MockPasswordHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPasswordHasher) EXPECT() *MockPasswordHasherMockRecorder {
	return m.recorder
}

// Hash mocks base method
func (m *MockPasswordHasher) Hash(unhashed domain.Password) (domain.Password, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", unhashed)
	ret0, _ := ret[0].(domain.Password)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash
func (mr *MockPasswordHasherMockRecorder) Hash(unhashed interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockPasswordHasher)(nil).Hash), unhashed)
}

// Compare mocks base method
func (m *MockPasswordHasher) Compare(hash domain.Password, original string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", hash, original)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Compare indicates an expected call of Compare
func (mr *MockPasswordHasherMockRecorder) Compare(hash, original interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockPasswordHasher)(nil).Compare), hash, original)
}
