// Code generated by MockGen. DO NOT EDIT.
// Source: api/repository/course_repository.go

// Package mockdb is a generated GoMock package.
package mockdb

import (
	reflect "reflect"
	models "github.com/Sugaml/mrc-api/api/models"

	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockICourse is a mock of ICourse interface.
type MockICourse struct {
	ctrl     *gomock.Controller
	recorder *MockICourseMockRecorder
}

// MockICourseMockRecorder is the mock recorder for MockICourse.
type MockICourseMockRecorder struct {
	mock *MockICourse
}

// NewMockICourse creates a new mock instance.
func NewMockICourse(ctrl *gomock.Controller) *MockICourse {
	mock := &MockICourse{ctrl: ctrl}
	mock.recorder = &MockICourseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICourse) EXPECT() *MockICourseMockRecorder {
	return m.recorder
}

// DeleteCourse mocks base method.
func (m *MockICourse) DeleteCourse(db *gorm.DB, cid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCourse", db, cid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCourse indicates an expected call of DeleteCourse.
func (mr *MockICourseMockRecorder) DeleteCourse(db, cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCourse", reflect.TypeOf((*MockICourse)(nil).DeleteCourse), db, cid)
}

// FindAllCourse mocks base method.
func (m *MockICourse) FindAllCourse(db *gorm.DB) (*[]models.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllCourse", db)
	ret0, _ := ret[0].(*[]models.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllCourse indicates an expected call of FindAllCourse.
func (mr *MockICourseMockRecorder) FindAllCourse(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllCourse", reflect.TypeOf((*MockICourse)(nil).FindAllCourse), db)
}

// FindbyId mocks base method.
func (m *MockICourse) FindbyId(db *gorm.DB, cid uint) (*models.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindbyId", db, cid)
	ret0, _ := ret[0].(*models.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindbyId indicates an expected call of FindbyId.
func (mr *MockICourseMockRecorder) FindbyId(db, cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindbyId", reflect.TypeOf((*MockICourse)(nil).FindbyId), db, cid)
}

// SaveCourse mocks base method.
func (m *MockICourse) SaveCourse(db *gorm.DB, course *models.Course) (*models.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCourse", db, course)
	ret0, _ := ret[0].(*models.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveCourse indicates an expected call of SaveCourse.
func (mr *MockICourseMockRecorder) SaveCourse(db, course interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCourse", reflect.TypeOf((*MockICourse)(nil).SaveCourse), db, course)
}

// UpdateCourse mocks base method.
func (m *MockICourse) UpdateCourse(db *gorm.DB, course *models.Course, cid uint) (*models.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourse", db, course, cid)
	ret0, _ := ret[0].(*models.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCourse indicates an expected call of UpdateCourse.
func (mr *MockICourseMockRecorder) UpdateCourse(db, course, cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourse", reflect.TypeOf((*MockICourse)(nil).UpdateCourse), db, course, cid)
}
