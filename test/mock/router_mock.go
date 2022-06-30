// Code generated by MockGen. DO NOT EDIT.
// Source: router.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockRouterRegisterInterface is a mock of RouterRegisterInterface interface.
type MockRouterRegisterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRouterRegisterInterfaceMockRecorder
}

// MockRouterRegisterInterfaceMockRecorder is the mock recorder for MockRouterRegisterInterface.
type MockRouterRegisterInterfaceMockRecorder struct {
	mock *MockRouterRegisterInterface
}

// NewMockRouterRegisterInterface creates a new mock instance.
func NewMockRouterRegisterInterface(ctrl *gomock.Controller) *MockRouterRegisterInterface {
	mock := &MockRouterRegisterInterface{ctrl: ctrl}
	mock.recorder = &MockRouterRegisterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouterRegisterInterface) EXPECT() *MockRouterRegisterInterfaceMockRecorder {
	return m.recorder
}

// TrackingRouter mocks base method.
func (m *MockRouterRegisterInterface) TrackingRouter(router *gin.RouterGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TrackingRouter", router)
}

// TrackingRouter indicates an expected call of TrackingRouter.
func (mr *MockRouterRegisterInterfaceMockRecorder) TrackingRouter(router interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackingRouter", reflect.TypeOf((*MockRouterRegisterInterface)(nil).TrackingRouter), router)
}
