// Code generated by MockGen. DO NOT EDIT.
// Source: ./message_encoder.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	message "github.com/topfreegames/pitaya/conn/message"
	reflect "reflect"
)

// MockEncoder is a mock of Encoder interface
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return m.recorder
}

// IsCompressionEnabled mocks base method
func (m *MockEncoder) IsCompressionEnabled() bool {
	ret := m.ctrl.Call(m, "IsCompressionEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCompressionEnabled indicates an expected call of IsCompressionEnabled
func (mr *MockEncoderMockRecorder) IsCompressionEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCompressionEnabled", reflect.TypeOf((*MockEncoder)(nil).IsCompressionEnabled))
}

// Encode mocks base method
func (m *MockEncoder) Encode(message *message.Message) ([]byte, error) {
	ret := m.ctrl.Call(m, "Encode", message)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode
func (mr *MockEncoderMockRecorder) Encode(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoder)(nil).Encode), message)
}
