// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dityuiri/go-adapter/kafka/producer (interfaces: IProducer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	kafka "github.com/dityuiri/go-adapter/kafka"
	gomock "github.com/golang/mock/gomock"
)

// MockIProducer is a mock of IProducer interface.
type MockIProducer struct {
	ctrl     *gomock.Controller
	recorder *MockIProducerMockRecorder
}

// MockIProducerMockRecorder is the mock recorder for MockIProducer.
type MockIProducerMockRecorder struct {
	mock *MockIProducer
}

// NewMockIProducer creates a new mock instance.
func NewMockIProducer(ctrl *gomock.Controller) *MockIProducer {
	mock := &MockIProducer{ctrl: ctrl}
	mock.recorder = &MockIProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProducer) EXPECT() *MockIProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIProducer)(nil).Close))
}

// Produce mocks base method.
func (m *MockIProducer) Produce(arg0 context.Context, arg1 string, arg2 ...*kafka.Message) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Produce", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockIProducerMockRecorder) Produce(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockIProducer)(nil).Produce), varargs...)
}