// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/redhat-developer/service-binding-operator/pkg/reconcile/pipeline/context (interfaces: K8STypeLookup)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kubernetes "github.com/redhat-developer/service-binding-operator/pkg/client/kubernetes"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// MockK8STypeLookup is a mock of K8STypeLookup interface.
type MockK8STypeLookup struct {
	ctrl     *gomock.Controller
	recorder *MockK8STypeLookupMockRecorder
}

// MockK8STypeLookupMockRecorder is the mock recorder for MockK8STypeLookup.
type MockK8STypeLookupMockRecorder struct {
	mock *MockK8STypeLookup
}

// NewMockK8STypeLookup creates a new mock instance.
func NewMockK8STypeLookup(ctrl *gomock.Controller) *MockK8STypeLookup {
	mock := &MockK8STypeLookup{ctrl: ctrl}
	mock.recorder = &MockK8STypeLookupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8STypeLookup) EXPECT() *MockK8STypeLookupMockRecorder {
	return m.recorder
}

// KindForResource mocks base method.
func (m *MockK8STypeLookup) KindForResource(arg0 schema.GroupVersionResource) (*schema.GroupVersionKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KindForResource", arg0)
	ret0, _ := ret[0].(*schema.GroupVersionKind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KindForResource indicates an expected call of KindForResource.
func (mr *MockK8STypeLookupMockRecorder) KindForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KindForResource", reflect.TypeOf((*MockK8STypeLookup)(nil).KindForResource), arg0)
}

// ResourceForKind mocks base method.
func (m *MockK8STypeLookup) ResourceForKind(arg0 schema.GroupVersionKind) (*schema.GroupVersionResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceForKind", arg0)
	ret0, _ := ret[0].(*schema.GroupVersionResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourceForKind indicates an expected call of ResourceForKind.
func (mr *MockK8STypeLookupMockRecorder) ResourceForKind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceForKind", reflect.TypeOf((*MockK8STypeLookup)(nil).ResourceForKind), arg0)
}

// ResourceForReferable mocks base method.
func (m *MockK8STypeLookup) ResourceForReferable(arg0 kubernetes.Referable) (*schema.GroupVersionResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceForReferable", arg0)
	ret0, _ := ret[0].(*schema.GroupVersionResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourceForReferable indicates an expected call of ResourceForReferable.
func (mr *MockK8STypeLookupMockRecorder) ResourceForReferable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceForReferable", reflect.TypeOf((*MockK8STypeLookup)(nil).ResourceForReferable), arg0)
}
