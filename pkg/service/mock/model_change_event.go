// Code generated by MockGen. DO NOT EDIT.
// Source: model_change_event.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	types "iam/pkg/service/types"
	reflect "reflect"
)

// MockModelChangeEventService is a mock of ModelChangeEventService interface
type MockModelChangeEventService struct {
	ctrl     *gomock.Controller
	recorder *MockModelChangeEventServiceMockRecorder
}

// MockModelChangeEventServiceMockRecorder is the mock recorder for MockModelChangeEventService
type MockModelChangeEventServiceMockRecorder struct {
	mock *MockModelChangeEventService
}

// NewMockModelChangeEventService creates a new mock instance
func NewMockModelChangeEventService(ctrl *gomock.Controller) *MockModelChangeEventService {
	mock := &MockModelChangeEventService{ctrl: ctrl}
	mock.recorder = &MockModelChangeEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelChangeEventService) EXPECT() *MockModelChangeEventServiceMockRecorder {
	return m.recorder
}

// ListByStatus mocks base method
func (m *MockModelChangeEventService) ListByStatus(status string) ([]types.ModelChangeEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByStatus", status)
	ret0, _ := ret[0].([]types.ModelChangeEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByStatus indicates an expected call of ListByStatus
func (mr *MockModelChangeEventServiceMockRecorder) ListByStatus(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByStatus", reflect.TypeOf((*MockModelChangeEventService)(nil).ListByStatus), status)
}

// UpdateStatusByPK mocks base method
func (m *MockModelChangeEventService) UpdateStatusByPK(pk int64, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusByPK", pk, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusByPK indicates an expected call of UpdateStatusByPK
func (mr *MockModelChangeEventServiceMockRecorder) UpdateStatusByPK(pk, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusByPK", reflect.TypeOf((*MockModelChangeEventService)(nil).UpdateStatusByPK), pk, status)
}

// BulkCreate mocks base method
func (m *MockModelChangeEventService) BulkCreate(modelChangeEvents []types.ModelChangeEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", modelChangeEvents)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate
func (mr *MockModelChangeEventServiceMockRecorder) BulkCreate(modelChangeEvents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockModelChangeEventService)(nil).BulkCreate), modelChangeEvents)
}

// ExistByTypeModel mocks base method
func (m *MockModelChangeEventService) ExistByTypeModel(eventType, status, modelType string, modelPK int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistByTypeModel", eventType, status, modelType, modelPK)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistByTypeModel indicates an expected call of ExistByTypeModel
func (mr *MockModelChangeEventServiceMockRecorder) ExistByTypeModel(eventType, status, modelType, modelPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistByTypeModel", reflect.TypeOf((*MockModelChangeEventService)(nil).ExistByTypeModel), eventType, status, modelType, modelPK)
}