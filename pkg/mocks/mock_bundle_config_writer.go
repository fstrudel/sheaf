// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bryanl/sheaf/pkg/sheaf (interfaces: BundleConfigWriter)

// Package mocks is a generated GoMock package.
package mocks

import (
	sheaf "github.com/bryanl/sheaf/pkg/sheaf"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBundleConfigWriter is a mock of BundleConfigWriter interface
type MockBundleConfigWriter struct {
	ctrl     *gomock.Controller
	recorder *MockBundleConfigWriterMockRecorder
}

// MockBundleConfigWriterMockRecorder is the mock recorder for MockBundleConfigWriter
type MockBundleConfigWriterMockRecorder struct {
	mock *MockBundleConfigWriter
}

// NewMockBundleConfigWriter creates a new mock instance
func NewMockBundleConfigWriter(ctrl *gomock.Controller) *MockBundleConfigWriter {
	mock := &MockBundleConfigWriter{ctrl: ctrl}
	mock.recorder = &MockBundleConfigWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBundleConfigWriter) EXPECT() *MockBundleConfigWriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockBundleConfigWriter) Write(arg0 sheaf.Bundle, arg1 sheaf.BundleConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockBundleConfigWriterMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockBundleConfigWriter)(nil).Write), arg0, arg1)
}