// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/makasim/amqpextra/consumer (interfaces: AMQPConnection,AMQPChannel)

// Package mock_consumer is a generated GoMock package.
package mock_consumer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

// MockAMQPConnection is a mock of AMQPConnection interface.
type MockAMQPConnection struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPConnectionMockRecorder
}

// MockAMQPConnectionMockRecorder is the mock recorder for MockAMQPConnection.
type MockAMQPConnectionMockRecorder struct {
	mock *MockAMQPConnection
}

// NewMockAMQPConnection creates a new mock instance.
func NewMockAMQPConnection(ctrl *gomock.Controller) *MockAMQPConnection {
	mock := &MockAMQPConnection{ctrl: ctrl}
	mock.recorder = &MockAMQPConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPConnection) EXPECT() *MockAMQPConnectionMockRecorder {
	return m.recorder
}

// MockAMQPChannel is a mock of AMQPChannel interface.
type MockAMQPChannel struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPChannelMockRecorder
}

// MockAMQPChannelMockRecorder is the mock recorder for MockAMQPChannel.
type MockAMQPChannelMockRecorder struct {
	mock *MockAMQPChannel
}

// NewMockAMQPChannel creates a new mock instance.
func NewMockAMQPChannel(ctrl *gomock.Controller) *MockAMQPChannel {
	mock := &MockAMQPChannel{ctrl: ctrl}
	mock.recorder = &MockAMQPChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPChannel) EXPECT() *MockAMQPChannelMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAMQPChannel) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPChannelMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPChannel)(nil).Close))
}

// Consume mocks base method.
func (m *MockAMQPChannel) Consume(arg0, arg1 string, arg2, arg3, arg4, arg5 bool, arg6 amqp091.Table) (<-chan amqp091.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(<-chan amqp091.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockAMQPChannelMockRecorder) Consume(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockAMQPChannel)(nil).Consume), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// NotifyCancel mocks base method.
func (m *MockAMQPChannel) NotifyCancel(arg0 chan string) chan string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyCancel", arg0)
	ret0, _ := ret[0].(chan string)
	return ret0
}

// NotifyCancel indicates an expected call of NotifyCancel.
func (mr *MockAMQPChannelMockRecorder) NotifyCancel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyCancel", reflect.TypeOf((*MockAMQPChannel)(nil).NotifyCancel), arg0)
}

// NotifyClose mocks base method.
func (m *MockAMQPChannel) NotifyClose(arg0 chan *amqp091.Error) chan *amqp091.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyClose", arg0)
	ret0, _ := ret[0].(chan *amqp091.Error)
	return ret0
}

// NotifyClose indicates an expected call of NotifyClose.
func (mr *MockAMQPChannelMockRecorder) NotifyClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyClose", reflect.TypeOf((*MockAMQPChannel)(nil).NotifyClose), arg0)
}

// Qos mocks base method.
func (m *MockAMQPChannel) Qos(arg0, arg1 int, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Qos", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Qos indicates an expected call of Qos.
func (mr *MockAMQPChannelMockRecorder) Qos(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Qos", reflect.TypeOf((*MockAMQPChannel)(nil).Qos), arg0, arg1, arg2)
}

// QueueBind mocks base method.
func (m *MockAMQPChannel) QueueBind(arg0, arg1, arg2 string, arg3 bool, arg4 amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueBind", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueBind indicates an expected call of QueueBind.
func (mr *MockAMQPChannelMockRecorder) QueueBind(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueBind", reflect.TypeOf((*MockAMQPChannel)(nil).QueueBind), arg0, arg1, arg2, arg3, arg4)
}

// QueueDeclare mocks base method.
func (m *MockAMQPChannel) QueueDeclare(arg0 string, arg1, arg2, arg3, arg4 bool, arg5 amqp091.Table) (amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDeclare", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(amqp091.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclare indicates an expected call of QueueDeclare.
func (mr *MockAMQPChannelMockRecorder) QueueDeclare(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclare", reflect.TypeOf((*MockAMQPChannel)(nil).QueueDeclare), arg0, arg1, arg2, arg3, arg4, arg5)
}
