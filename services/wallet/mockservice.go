// Code generated by MockGen. DO NOT EDIT.
// Source: ./wallet/service.go

// Package wallet is a generated GoMock package.
package wallet

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGeoValidator is a mock of GeoValidator interface.
type MockGeoValidator struct {
	ctrl     *gomock.Controller
	recorder *MockGeoValidatorMockRecorder
}

// MockGeoValidatorMockRecorder is the mock recorder for MockGeoValidator.
type MockGeoValidatorMockRecorder struct {
	mock *MockGeoValidator
}

// NewMockGeoValidator creates a new mock instance.
func NewMockGeoValidator(ctrl *gomock.Controller) *MockGeoValidator {
	mock := &MockGeoValidator{ctrl: ctrl}
	mock.recorder = &MockGeoValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeoValidator) EXPECT() *MockGeoValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockGeoValidator) Validate(ctx context.Context, geolocation string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, geolocation)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockGeoValidatorMockRecorder) Validate(ctx, geolocation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockGeoValidator)(nil).Validate), ctx, geolocation)
}

// MockmetricSvc is a mock of metricSvc interface.
type MockmetricSvc struct {
	ctrl     *gomock.Controller
	recorder *MockmetricSvcMockRecorder
}

// MockmetricSvcMockRecorder is the mock recorder for MockmetricSvc.
type MockmetricSvcMockRecorder struct {
	mock *MockmetricSvc
}

// NewMockmetricSvc creates a new mock instance.
func NewMockmetricSvc(ctrl *gomock.Controller) *MockmetricSvc {
	mock := &MockmetricSvc{ctrl: ctrl}
	mock.recorder = &MockmetricSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmetricSvc) EXPECT() *MockmetricSvcMockRecorder {
	return m.recorder
}

// LinkFailureZP mocks base method.
func (m *MockmetricSvc) LinkFailureZP(cc string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LinkFailureZP", cc)
}

// LinkFailureZP indicates an expected call of LinkFailureZP.
func (mr *MockmetricSvcMockRecorder) LinkFailureZP(cc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkFailureZP", reflect.TypeOf((*MockmetricSvc)(nil).LinkFailureZP), cc)
}

// LinkSuccessZP mocks base method.
func (m *MockmetricSvc) LinkSuccessZP(cc string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LinkSuccessZP", cc)
}

// LinkSuccessZP indicates an expected call of LinkSuccessZP.
func (mr *MockmetricSvcMockRecorder) LinkSuccessZP(cc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSuccessZP", reflect.TypeOf((*MockmetricSvc)(nil).LinkSuccessZP), cc)
}
