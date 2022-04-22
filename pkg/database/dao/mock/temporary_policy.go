// Code generated by MockGen. DO NOT EDIT.
// Source: temporary_policy.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	dao "iam/pkg/database/dao"
	reflect "reflect"
)

// MockTemporaryPolicyManager is a mock of TemporaryPolicyManager interface
type MockTemporaryPolicyManager struct {
	ctrl     *gomock.Controller
	recorder *MockTemporaryPolicyManagerMockRecorder
}

// MockTemporaryPolicyManagerMockRecorder is the mock recorder for MockTemporaryPolicyManager
type MockTemporaryPolicyManagerMockRecorder struct {
	mock *MockTemporaryPolicyManager
}

// NewMockTemporaryPolicyManager creates a new mock instance
func NewMockTemporaryPolicyManager(ctrl *gomock.Controller) *MockTemporaryPolicyManager {
	mock := &MockTemporaryPolicyManager{ctrl: ctrl}
	mock.recorder = &MockTemporaryPolicyManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTemporaryPolicyManager) EXPECT() *MockTemporaryPolicyManagerMockRecorder {
	return m.recorder
}

// ListThinBySubjectAction mocks base method
func (m *MockTemporaryPolicyManager) ListThinBySubjectAction(subjectPK, actionPK, expiredAt int64) ([]dao.ThinTemporaryPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListThinBySubjectAction", subjectPK, actionPK, expiredAt)
	ret0, _ := ret[0].([]dao.ThinTemporaryPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListThinBySubjectAction indicates an expected call of ListThinBySubjectAction
func (mr *MockTemporaryPolicyManagerMockRecorder) ListThinBySubjectAction(subjectPK, actionPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThinBySubjectAction", reflect.TypeOf((*MockTemporaryPolicyManager)(nil).ListThinBySubjectAction), subjectPK, actionPK, expiredAt)
}

// ListByPKs mocks base method
func (m *MockTemporaryPolicyManager) ListByPKs(pks []int64) ([]dao.TemporaryPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPKs", pks)
	ret0, _ := ret[0].([]dao.TemporaryPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByPKs indicates an expected call of ListByPKs
func (mr *MockTemporaryPolicyManagerMockRecorder) ListByPKs(pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPKs", reflect.TypeOf((*MockTemporaryPolicyManager)(nil).ListByPKs), pks)
}

// BulkCreateWithTx mocks base method
func (m *MockTemporaryPolicyManager) BulkCreateWithTx(tx *sqlx.Tx, policies []dao.TemporaryPolicy) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateWithTx", tx, policies)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkCreateWithTx indicates an expected call of BulkCreateWithTx
func (mr *MockTemporaryPolicyManagerMockRecorder) BulkCreateWithTx(tx, policies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateWithTx", reflect.TypeOf((*MockTemporaryPolicyManager)(nil).BulkCreateWithTx), tx, policies)
}

// BulkDeleteByPKs mocks base method
func (m *MockTemporaryPolicyManager) BulkDeleteByPKs(subjectPK int64, pks []int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteByPKs", subjectPK, pks)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteByPKs indicates an expected call of BulkDeleteByPKs
func (mr *MockTemporaryPolicyManagerMockRecorder) BulkDeleteByPKs(subjectPK, pks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteByPKs", reflect.TypeOf((*MockTemporaryPolicyManager)(nil).BulkDeleteByPKs), subjectPK, pks)
}

// BulkDeleteBeforeExpiredAtWithTx mocks base method
func (m *MockTemporaryPolicyManager) BulkDeleteBeforeExpiredAtWithTx(tx *sqlx.Tx, expiredAt, limit int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBeforeExpiredAtWithTx", tx, expiredAt, limit)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteBeforeExpiredAtWithTx indicates an expected call of BulkDeleteBeforeExpiredAtWithTx
func (mr *MockTemporaryPolicyManagerMockRecorder) BulkDeleteBeforeExpiredAtWithTx(tx, expiredAt, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBeforeExpiredAtWithTx", reflect.TypeOf((*MockTemporaryPolicyManager)(nil).BulkDeleteBeforeExpiredAtWithTx), tx, expiredAt, limit)
}