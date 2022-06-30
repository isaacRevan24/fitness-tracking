// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/isaacRevan24/fitness-tracking/model"
)

// MockTrackingHandlerInterface is a mock of TrackingHandlerInterface interface.
type MockTrackingHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTrackingHandlerInterfaceMockRecorder
}

// MockTrackingHandlerInterfaceMockRecorder is the mock recorder for MockTrackingHandlerInterface.
type MockTrackingHandlerInterfaceMockRecorder struct {
	mock *MockTrackingHandlerInterface
}

// NewMockTrackingHandlerInterface creates a new mock instance.
func NewMockTrackingHandlerInterface(ctrl *gomock.Controller) *MockTrackingHandlerInterface {
	mock := &MockTrackingHandlerInterface{ctrl: ctrl}
	mock.recorder = &MockTrackingHandlerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackingHandlerInterface) EXPECT() *MockTrackingHandlerInterfaceMockRecorder {
	return m.recorder
}

// AddWeightRegister mocks base method.
func (m *MockTrackingHandlerInterface) AddWeightRegister(request model.AddWeightRegisterReq) model.FitnessResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWeightRegister", request)
	ret0, _ := ret[0].(model.FitnessResponse)
	return ret0
}

// AddWeightRegister indicates an expected call of AddWeightRegister.
func (mr *MockTrackingHandlerInterfaceMockRecorder) AddWeightRegister(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWeightRegister", reflect.TypeOf((*MockTrackingHandlerInterface)(nil).AddWeightRegister), request)
}
