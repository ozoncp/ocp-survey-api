// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-survey-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-survey-api/internal/models"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddSurvey mocks base method.
func (m *MockRepo) AddSurvey(arg0 context.Context, arg1 []models.Survey) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSurvey", arg0, arg1)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSurvey indicates an expected call of AddSurvey.
func (mr *MockRepoMockRecorder) AddSurvey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSurvey", reflect.TypeOf((*MockRepo)(nil).AddSurvey), arg0, arg1)
}

// DescribeSurvey mocks base method.
func (m *MockRepo) DescribeSurvey(arg0 context.Context, arg1 uint64) (*models.Survey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSurvey", arg0, arg1)
	ret0, _ := ret[0].(*models.Survey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSurvey indicates an expected call of DescribeSurvey.
func (mr *MockRepoMockRecorder) DescribeSurvey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSurvey", reflect.TypeOf((*MockRepo)(nil).DescribeSurvey), arg0, arg1)
}

// ListSurveys mocks base method.
func (m *MockRepo) ListSurveys(arg0 context.Context, arg1, arg2 uint64) ([]models.Survey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSurveys", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Survey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSurveys indicates an expected call of ListSurveys.
func (mr *MockRepoMockRecorder) ListSurveys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSurveys", reflect.TypeOf((*MockRepo)(nil).ListSurveys), arg0, arg1, arg2)
}

// RemoveSurvey mocks base method.
func (m *MockRepo) RemoveSurvey(arg0 context.Context, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSurvey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSurvey indicates an expected call of RemoveSurvey.
func (mr *MockRepoMockRecorder) RemoveSurvey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSurvey", reflect.TypeOf((*MockRepo)(nil).RemoveSurvey), arg0, arg1)
}

// UpdateSurvey mocks base method.
func (m *MockRepo) UpdateSurvey(arg0 context.Context, arg1 models.Survey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSurvey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSurvey indicates an expected call of UpdateSurvey.
func (mr *MockRepoMockRecorder) UpdateSurvey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSurvey", reflect.TypeOf((*MockRepo)(nil).UpdateSurvey), arg0, arg1)
}
