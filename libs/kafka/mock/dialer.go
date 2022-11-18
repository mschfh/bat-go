// Code generated by MockGen. DO NOT EDIT.
// Source: ./kafka/dialer.go

// Package mockdialer is a generated GoMock package.
package mockdialer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kafka "github.com/segmentio/kafka-go"
)

// MockConsumer is a mock of Consumer interface.
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer.
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance.
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// CommitMessages mocks base method.
func (m *MockConsumer) CommitMessages(ctx context.Context, messages ...kafka.Message) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range messages {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommitMessages", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitMessages indicates an expected call of CommitMessages.
func (mr *MockConsumerMockRecorder) CommitMessages(ctx interface{}, messages ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, messages...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitMessages", reflect.TypeOf((*MockConsumer)(nil).CommitMessages), varargs...)
}

// FetchMessage mocks base method.
func (m *MockConsumer) FetchMessage(ctx context.Context) (kafka.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMessage", ctx)
	ret0, _ := ret[0].(kafka.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMessage indicates an expected call of FetchMessage.
func (mr *MockConsumerMockRecorder) FetchMessage(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMessage", reflect.TypeOf((*MockConsumer)(nil).FetchMessage), ctx)
}

// ReadMessage mocks base method.
func (m *MockConsumer) ReadMessage(ctx context.Context) (kafka.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage", ctx)
	ret0, _ := ret[0].(kafka.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMessage indicates an expected call of ReadMessage.
func (mr *MockConsumerMockRecorder) ReadMessage(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockConsumer)(nil).ReadMessage), ctx)
}
