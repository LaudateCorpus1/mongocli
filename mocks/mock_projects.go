// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/projects.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProjectLister is a mock of ProjectLister interface
type MockProjectLister struct {
	ctrl     *gomock.Controller
	recorder *MockProjectListerMockRecorder
}

// MockProjectListerMockRecorder is the mock recorder for MockProjectLister
type MockProjectListerMockRecorder struct {
	mock *MockProjectLister
}

// NewMockProjectLister creates a new mock instance
func NewMockProjectLister(ctrl *gomock.Controller) *MockProjectLister {
	mock := &MockProjectLister{ctrl: ctrl}
	mock.recorder = &MockProjectListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectLister) EXPECT() *MockProjectListerMockRecorder {
	return m.recorder
}

// GetAllProjects mocks base method
func (m *MockProjectLister) GetAllProjects() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProjects")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProjects indicates an expected call of GetAllProjects
func (mr *MockProjectListerMockRecorder) GetAllProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProjects", reflect.TypeOf((*MockProjectLister)(nil).GetAllProjects))
}

// MockProjectCreator is a mock of ProjectCreator interface
type MockProjectCreator struct {
	ctrl     *gomock.Controller
	recorder *MockProjectCreatorMockRecorder
}

// MockProjectCreatorMockRecorder is the mock recorder for MockProjectCreator
type MockProjectCreatorMockRecorder struct {
	mock *MockProjectCreator
}

// NewMockProjectCreator creates a new mock instance
func NewMockProjectCreator(ctrl *gomock.Controller) *MockProjectCreator {
	mock := &MockProjectCreator{ctrl: ctrl}
	mock.recorder = &MockProjectCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectCreator) EXPECT() *MockProjectCreatorMockRecorder {
	return m.recorder
}

// CreateProject mocks base method
func (m *MockProjectCreator) CreateProject(arg0, arg1 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject
func (mr *MockProjectCreatorMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectCreator)(nil).CreateProject), arg0, arg1)
}

// MockProjectStore is a mock of ProjectStore interface
type MockProjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockProjectStoreMockRecorder
}

// MockProjectStoreMockRecorder is the mock recorder for MockProjectStore
type MockProjectStoreMockRecorder struct {
	mock *MockProjectStore
}

// NewMockProjectStore creates a new mock instance
func NewMockProjectStore(ctrl *gomock.Controller) *MockProjectStore {
	mock := &MockProjectStore{ctrl: ctrl}
	mock.recorder = &MockProjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectStore) EXPECT() *MockProjectStoreMockRecorder {
	return m.recorder
}

// GetAllProjects mocks base method
func (m *MockProjectStore) GetAllProjects() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProjects")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProjects indicates an expected call of GetAllProjects
func (mr *MockProjectStoreMockRecorder) GetAllProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProjects", reflect.TypeOf((*MockProjectStore)(nil).GetAllProjects))
}

// CreateProject mocks base method
func (m *MockProjectStore) CreateProject(arg0, arg1 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject
func (mr *MockProjectStoreMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectStore)(nil).CreateProject), arg0, arg1)
}
