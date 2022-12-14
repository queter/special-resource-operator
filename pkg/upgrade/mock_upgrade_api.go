// Code generated by MockGen. DO NOT EDIT.
// Source: upgrade.go

// Package upgrade is a generated GoMock package.
package upgrade

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockClusterInfo is a mock of ClusterInfo interface.
type MockClusterInfo struct {
	ctrl     *gomock.Controller
	recorder *MockClusterInfoMockRecorder
}

// MockClusterInfoMockRecorder is the mock recorder for MockClusterInfo.
type MockClusterInfoMockRecorder struct {
	mock *MockClusterInfo
}

// NewMockClusterInfo creates a new mock instance.
func NewMockClusterInfo(ctrl *gomock.Controller) *MockClusterInfo {
	mock := &MockClusterInfo{ctrl: ctrl}
	mock.recorder = &MockClusterInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterInfo) EXPECT() *MockClusterInfoMockRecorder {
	return m.recorder
}

// GetClusterInfo mocks base method.
func (m *MockClusterInfo) GetClusterInfo(arg0 context.Context, arg1 *v1.NodeList) (map[string]NodeVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInfo", arg0, arg1)
	ret0, _ := ret[0].(map[string]NodeVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInfo indicates an expected call of GetClusterInfo.
func (mr *MockClusterInfoMockRecorder) GetClusterInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInfo", reflect.TypeOf((*MockClusterInfo)(nil).GetClusterInfo), arg0, arg1)
}
