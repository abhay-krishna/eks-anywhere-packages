// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/bundle/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	types "k8s.io/apimachinery/pkg/types"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetActiveBundle mocks base method.
func (m *MockClient) GetActiveBundle(ctx context.Context) (*v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveBundle", ctx)
	ret0, _ := ret[0].(*v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveBundle indicates an expected call of GetActiveBundle.
func (mr *MockClientMockRecorder) GetActiveBundle(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveBundle", reflect.TypeOf((*MockClient)(nil).GetActiveBundle), ctx)
}

// GetActiveBundleNamespacedName mocks base method.
func (m *MockClient) GetActiveBundleNamespacedName(ctx context.Context) (types.NamespacedName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveBundleNamespacedName", ctx)
	ret0, _ := ret[0].(types.NamespacedName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveBundleNamespacedName indicates an expected call of GetActiveBundleNamespacedName.
func (mr *MockClientMockRecorder) GetActiveBundleNamespacedName(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveBundleNamespacedName", reflect.TypeOf((*MockClient)(nil).GetActiveBundleNamespacedName), ctx)
}

// IsActive mocks base method.
func (m *MockClient) IsActive(ctx context.Context, packageBundle *v1alpha1.PackageBundle) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive", ctx, packageBundle)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsActive indicates an expected call of IsActive.
func (mr *MockClientMockRecorder) IsActive(ctx, packageBundle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockClient)(nil).IsActive), ctx, packageBundle)
}
