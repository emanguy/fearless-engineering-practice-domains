// Code generated by MockGen. DO NOT EDIT.
// Source: example.com/sample/commonlib/database (interfaces: Connection)
//
// Generated by this command:
//
//	mockgen -destination ./context_mocks.go -package database . Connection
//
// Package database is a generated GoMock package.
package database

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	sqlx "github.com/jmoiron/sqlx"
	gomock "go.uber.org/mock/gomock"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockConnection) Exec(arg0 string, arg1 ...any) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockConnectionMockRecorder) Exec(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockConnection)(nil).Exec), varargs...)
}

// ExecContext mocks base method.
func (m *MockConnection) ExecContext(arg0 context.Context, arg1 string, arg2 ...any) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockConnectionMockRecorder) ExecContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockConnection)(nil).ExecContext), varargs...)
}

// Get mocks base method.
func (m *MockConnection) Get(arg0 any, arg1 string, arg2 ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockConnectionMockRecorder) Get(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConnection)(nil).Get), varargs...)
}

// GetContext mocks base method.
func (m *MockConnection) GetContext(arg0 context.Context, arg1 any, arg2 string, arg3 ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockConnectionMockRecorder) GetContext(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockConnection)(nil).GetContext), varargs...)
}

// NamedExec mocks base method.
func (m *MockConnection) NamedExec(arg0 string, arg1 any) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedExec", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExec indicates an expected call of NamedExec.
func (mr *MockConnectionMockRecorder) NamedExec(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExec", reflect.TypeOf((*MockConnection)(nil).NamedExec), arg0, arg1)
}

// NamedExecContext mocks base method.
func (m *MockConnection) NamedExecContext(arg0 context.Context, arg1 string, arg2 any) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedExecContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExecContext indicates an expected call of NamedExecContext.
func (mr *MockConnectionMockRecorder) NamedExecContext(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExecContext", reflect.TypeOf((*MockConnection)(nil).NamedExecContext), arg0, arg1, arg2)
}

// NamedQuery mocks base method.
func (m *MockConnection) NamedQuery(arg0 string, arg1 any) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedQuery", arg0, arg1)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedQuery indicates an expected call of NamedQuery.
func (mr *MockConnectionMockRecorder) NamedQuery(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedQuery", reflect.TypeOf((*MockConnection)(nil).NamedQuery), arg0, arg1)
}

// PrepareNamed mocks base method.
func (m *MockConnection) PrepareNamed(arg0 string) (*sqlx.NamedStmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareNamed", arg0)
	ret0, _ := ret[0].(*sqlx.NamedStmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareNamed indicates an expected call of PrepareNamed.
func (mr *MockConnectionMockRecorder) PrepareNamed(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareNamed", reflect.TypeOf((*MockConnection)(nil).PrepareNamed), arg0)
}

// PrepareNamedContext mocks base method.
func (m *MockConnection) PrepareNamedContext(arg0 context.Context, arg1 string) (*sqlx.NamedStmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareNamedContext", arg0, arg1)
	ret0, _ := ret[0].(*sqlx.NamedStmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareNamedContext indicates an expected call of PrepareNamedContext.
func (mr *MockConnectionMockRecorder) PrepareNamedContext(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareNamedContext", reflect.TypeOf((*MockConnection)(nil).PrepareNamedContext), arg0, arg1)
}

// Preparex mocks base method.
func (m *MockConnection) Preparex(arg0 string) (*sqlx.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Preparex", arg0)
	ret0, _ := ret[0].(*sqlx.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Preparex indicates an expected call of Preparex.
func (mr *MockConnectionMockRecorder) Preparex(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preparex", reflect.TypeOf((*MockConnection)(nil).Preparex), arg0)
}

// PreparexContext mocks base method.
func (m *MockConnection) PreparexContext(arg0 context.Context, arg1 string) (*sqlx.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreparexContext", arg0, arg1)
	ret0, _ := ret[0].(*sqlx.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreparexContext indicates an expected call of PreparexContext.
func (mr *MockConnectionMockRecorder) PreparexContext(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreparexContext", reflect.TypeOf((*MockConnection)(nil).PreparexContext), arg0, arg1)
}

// Queryx mocks base method.
func (m *MockConnection) Queryx(arg0 string, arg1 ...any) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Queryx", varargs...)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Queryx indicates an expected call of Queryx.
func (mr *MockConnectionMockRecorder) Queryx(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Queryx", reflect.TypeOf((*MockConnection)(nil).Queryx), varargs...)
}

// QueryxContext mocks base method.
func (m *MockConnection) QueryxContext(arg0 context.Context, arg1 string, arg2 ...any) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryxContext", varargs...)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryxContext indicates an expected call of QueryxContext.
func (mr *MockConnectionMockRecorder) QueryxContext(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryxContext", reflect.TypeOf((*MockConnection)(nil).QueryxContext), varargs...)
}

// Rebind mocks base method.
func (m *MockConnection) Rebind(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebind", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Rebind indicates an expected call of Rebind.
func (mr *MockConnectionMockRecorder) Rebind(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebind", reflect.TypeOf((*MockConnection)(nil).Rebind), arg0)
}

// Select mocks base method.
func (m *MockConnection) Select(arg0 any, arg1 string, arg2 ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockConnectionMockRecorder) Select(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockConnection)(nil).Select), varargs...)
}

// SelectContext mocks base method.
func (m *MockConnection) SelectContext(arg0 context.Context, arg1 any, arg2 string, arg3 ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectContext indicates an expected call of SelectContext.
func (mr *MockConnectionMockRecorder) SelectContext(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockConnection)(nil).SelectContext), varargs...)
}
