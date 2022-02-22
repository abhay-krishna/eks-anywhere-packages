// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/driver/packagedriver.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	api "github.com/aws/modelrocket-add-ons/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockPackageDriver is a mock of PackageDriver interface.
type MockPackageDriver struct {
	ctrl     *gomock.Controller
	recorder *MockPackageDriverMockRecorder
}

// MockPackageDriverMockRecorder is the mock recorder for MockPackageDriver.
type MockPackageDriverMockRecorder struct {
	mock *MockPackageDriver
}

// NewMockPackageDriver creates a new mock instance.
func NewMockPackageDriver(ctrl *gomock.Controller) *MockPackageDriver {
	mock := &MockPackageDriver{ctrl: ctrl}
	mock.recorder = &MockPackageDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackageDriver) EXPECT() *MockPackageDriverMockRecorder {
	return m.recorder
}

// Install mocks base method.
func (m *MockPackageDriver) Install(ctx context.Context, name, namespace string, source api.PackageOCISource, values map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ctx, name, namespace, source, values)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockPackageDriverMockRecorder) Install(ctx, name, namespace, source, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockPackageDriver)(nil).Install), ctx, name, namespace, source, values)
}

// Uninstall mocks base method.
func (m *MockPackageDriver) Uninstall(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockPackageDriverMockRecorder) Uninstall(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockPackageDriver)(nil).Uninstall), ctx, name)
}
