// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dapperlabs/flow-go/protocol (interfaces: Snapshot)

// Package mocks is a generated GoMock package.
package mocks

import (
	flow "github.com/dapperlabs/flow-go/model/flow"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSnapshot is a mock of Snapshot interface
type MockSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotMockRecorder
}

// MockSnapshotMockRecorder is the mock recorder for MockSnapshot
type MockSnapshotMockRecorder struct {
	mock *MockSnapshot
}

// NewMockSnapshot creates a new mock instance
func NewMockSnapshot(ctrl *gomock.Controller) *MockSnapshot {
	mock := &MockSnapshot{ctrl: ctrl}
	mock.recorder = &MockSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshot) EXPECT() *MockSnapshotMockRecorder {
	return m.recorder
}

// Clusters mocks base method
func (m *MockSnapshot) Clusters() (*flow.ClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clusters")
	ret0, _ := ret[0].(*flow.ClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Clusters indicates an expected call of Clusters
func (mr *MockSnapshotMockRecorder) Clusters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clusters", reflect.TypeOf((*MockSnapshot)(nil).Clusters))
}

// Commit mocks base method
func (m *MockSnapshot) Commit() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit
func (mr *MockSnapshotMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockSnapshot)(nil).Commit))
}

// Head mocks base method
func (m *MockSnapshot) Head() (*flow.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head")
	ret0, _ := ret[0].(*flow.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head
func (mr *MockSnapshotMockRecorder) Head() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockSnapshot)(nil).Head))
}

// Identities mocks base method
func (m *MockSnapshot) Identities(arg0 ...flow.IdentityFilter) (flow.IdentityList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Identities", varargs...)
	ret0, _ := ret[0].(flow.IdentityList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Identities indicates an expected call of Identities
func (mr *MockSnapshotMockRecorder) Identities(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identities", reflect.TypeOf((*MockSnapshot)(nil).Identities), arg0...)
}

// Identity mocks base method
func (m *MockSnapshot) Identity(arg0 flow.Identifier) (*flow.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identity", arg0)
	ret0, _ := ret[0].(*flow.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Identity indicates an expected call of Identity
func (mr *MockSnapshotMockRecorder) Identity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockSnapshot)(nil).Identity), arg0)
}
