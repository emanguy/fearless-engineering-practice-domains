// Code generated by MockGen. DO NOT EDIT.
// Source: example.com/sample/commonlib/sharedfeatures/loglevel (interfaces: Core)
//
// Generated by this command:
//
//	mockgen -destination ./adjust_log_leveL_mocks.go -package loglevel . Core
//
// Package loglevel is a generated GoMock package.
package loglevel

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	zapcore "go.uber.org/zap/zapcore"
)

// MockCore is a mock of Core interface.
type MockCore struct {
	ctrl     *gomock.Controller
	recorder *MockCoreMockRecorder
}

// MockCoreMockRecorder is the mock recorder for MockCore.
type MockCoreMockRecorder struct {
	mock *MockCore
}

// NewMockCore creates a new mock instance.
func NewMockCore(ctrl *gomock.Controller) *MockCore {
	mock := &MockCore{ctrl: ctrl}
	mock.recorder = &MockCoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCore) EXPECT() *MockCoreMockRecorder {
	return m.recorder
}

// SetLogLevel mocks base method.
func (m *MockCore) SetLogLevel(arg0 zapcore.Level) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogLevel", arg0)
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MockCoreMockRecorder) SetLogLevel(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockCore)(nil).SetLogLevel), arg0)
}
