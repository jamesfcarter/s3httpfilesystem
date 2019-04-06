// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_s3httpfilesystem is a generated GoMock package.
package mock_s3httpfilesystem

import (
	s3 "github.com/aws/aws-sdk-go/service/s3"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mocks3Interface is a mock of s3Interface interface
type Mocks3Interface struct {
	ctrl     *gomock.Controller
	recorder *Mocks3InterfaceMockRecorder
}

// Mocks3InterfaceMockRecorder is the mock recorder for Mocks3Interface
type Mocks3InterfaceMockRecorder struct {
	mock *Mocks3Interface
}

// NewMocks3Interface creates a new mock instance
func NewMocks3Interface(ctrl *gomock.Controller) *Mocks3Interface {
	mock := &Mocks3Interface{ctrl: ctrl}
	mock.recorder = &Mocks3InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mocks3Interface) EXPECT() *Mocks3InterfaceMockRecorder {
	return m.recorder
}

// ListObjects mocks base method
func (m *Mocks3Interface) ListObjects(arg0 *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjects", arg0)
	ret0, _ := ret[0].(*s3.ListObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjects indicates an expected call of ListObjects
func (mr *Mocks3InterfaceMockRecorder) ListObjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*Mocks3Interface)(nil).ListObjects), arg0)
}

// GetObject mocks base method
func (m *Mocks3Interface) GetObject(arg0 *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", arg0)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *Mocks3InterfaceMockRecorder) GetObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*Mocks3Interface)(nil).GetObject), arg0)
}
