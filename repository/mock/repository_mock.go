// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/isaacRevan24/fitness-tracking/model"
)

// MockTrackingRepository is a mock of TrackingRepository interface.
type MockTrackingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTrackingRepositoryMockRecorder
}

// MockTrackingRepositoryMockRecorder is the mock recorder for MockTrackingRepository.
type MockTrackingRepositoryMockRecorder struct {
	mock *MockTrackingRepository
}

// NewMockTrackingRepository creates a new mock instance.
func NewMockTrackingRepository(ctrl *gomock.Controller) *MockTrackingRepository {
	mock := &MockTrackingRepository{ctrl: ctrl}
	mock.recorder = &MockTrackingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackingRepository) EXPECT() *MockTrackingRepositoryMockRecorder {
	return m.recorder
}

// AddWeightRegister mocks base method.
func (m *MockTrackingRepository) AddWeightRegister(request model.AddWeightRegisterReq) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWeightRegister", request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddWeightRegister indicates an expected call of AddWeightRegister.
func (mr *MockTrackingRepositoryMockRecorder) AddWeightRegister(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWeightRegister", reflect.TypeOf((*MockTrackingRepository)(nil).AddWeightRegister), request)
}

// DeleteWeightRegister mocks base method.
func (m *MockTrackingRepository) DeleteWeightRegister(request model.DeleteWeightRegisterReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWeightRegister", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWeightRegister indicates an expected call of DeleteWeightRegister.
func (mr *MockTrackingRepositoryMockRecorder) DeleteWeightRegister(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWeightRegister", reflect.TypeOf((*MockTrackingRepository)(nil).DeleteWeightRegister), request)
}

// GetWeightRegister mocks base method.
func (m *MockTrackingRepository) GetWeightRegister(clientId, weightId string) (model.GetWeightRegisterRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeightRegister", clientId, weightId)
	ret0, _ := ret[0].(model.GetWeightRegisterRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeightRegister indicates an expected call of GetWeightRegister.
func (mr *MockTrackingRepositoryMockRecorder) GetWeightRegister(clientId, weightId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeightRegister", reflect.TypeOf((*MockTrackingRepository)(nil).GetWeightRegister), clientId, weightId)
}

// UpdateWeightRegister mocks base method.
func (m *MockTrackingRepository) UpdateWeightRegister(request model.UpdateWeightRegisterReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWeightRegister", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWeightRegister indicates an expected call of UpdateWeightRegister.
func (mr *MockTrackingRepositoryMockRecorder) UpdateWeightRegister(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWeightRegister", reflect.TypeOf((*MockTrackingRepository)(nil).UpdateWeightRegister), request)
}
