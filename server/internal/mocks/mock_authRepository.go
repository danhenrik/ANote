// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\julio\OneDrive\Documentos\GitHub\pds\ANote\server\internal\interfaces\repositories\authRepository.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	domain "anote/internal/domain"
	errors "anote/internal/errors"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// DeleteToken mocks base method.
func (m *MockAuthRepository) DeleteToken(token string) *errors.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteToken", token)
	ret0, _ := ret[0].(*errors.AppError)
	return ret0
}

// DeleteToken indicates an expected call of DeleteToken.
func (mr *MockAuthRepositoryMockRecorder) DeleteToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToken", reflect.TypeOf((*MockAuthRepository)(nil).DeleteToken), token)
}

// RetrieveToken mocks base method.
func (m *MockAuthRepository) RetrieveToken(token string) (*domain.User, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveToken", token)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// RetrieveToken indicates an expected call of RetrieveToken.
func (mr *MockAuthRepositoryMockRecorder) RetrieveToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveToken", reflect.TypeOf((*MockAuthRepository)(nil).RetrieveToken), token)
}

// SaveToken mocks base method.
func (m *MockAuthRepository) SaveToken(token, userId string) *errors.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveToken", token, userId)
	ret0, _ := ret[0].(*errors.AppError)
	return ret0
}

// SaveToken indicates an expected call of SaveToken.
func (mr *MockAuthRepositoryMockRecorder) SaveToken(token, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveToken", reflect.TypeOf((*MockAuthRepository)(nil).SaveToken), token, userId)
}
