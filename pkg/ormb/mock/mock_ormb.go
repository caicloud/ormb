// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ormb/ormb.go

// Package mock_ormb is a generated GoMock package.
package mock_ormb

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockInterface) Login(hostname, username, password string, insecureOpt bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", hostname, username, password, insecureOpt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockInterfaceMockRecorder) Login(hostname, username, password, insecureOpt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockInterface)(nil).Login), hostname, username, password, insecureOpt)
}

// Push mocks base method
func (m *MockInterface) Push(refStr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", refStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push
func (mr *MockInterfaceMockRecorder) Push(refStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockInterface)(nil).Push), refStr)
}

// Pull mocks base method
func (m *MockInterface) Pull(refStr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", refStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull
func (mr *MockInterfaceMockRecorder) Pull(refStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockInterface)(nil).Pull), refStr)
}

// Export mocks base method
func (m *MockInterface) Export(refStr, dst string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", refStr, dst)
	ret0, _ := ret[0].(error)
	return ret0
}

// Export indicates an expected call of Export
func (mr *MockInterfaceMockRecorder) Export(refStr, dst interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockInterface)(nil).Export), refStr, dst)
}

// Save mocks base method
func (m *MockInterface) Save(src, refStr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", src, refStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockInterfaceMockRecorder) Save(src, refStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockInterface)(nil).Save), src, refStr)
}

// Tag mocks base method
func (m *MockInterface) Tag(refStr, targetStr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag", refStr, targetStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tag indicates an expected call of Tag
func (mr *MockInterfaceMockRecorder) Tag(refStr, targetStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockInterface)(nil).Tag), refStr, targetStr)
}

// Remove mocks base method
func (m *MockInterface) Remove(refStr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", refStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockInterfaceMockRecorder) Remove(refStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockInterface)(nil).Remove), refStr)
}
