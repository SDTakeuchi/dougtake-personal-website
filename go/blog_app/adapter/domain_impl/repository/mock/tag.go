// Code generated by MockGen. DO NOT EDIT.
// Source: tag.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "blog_app/domain/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTag is a mock of Tag interface.
type MockTag struct {
	ctrl     *gomock.Controller
	recorder *MockTagMockRecorder
}

// MockTagMockRecorder is the mock recorder for MockTag.
type MockTagMockRecorder struct {
	mock *MockTag
}

// NewMockTag creates a new mock instance.
func NewMockTag(ctrl *gomock.Controller) *MockTag {
	mock := &MockTag{ctrl: ctrl}
	mock.recorder = &MockTagMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTag) EXPECT() *MockTagMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTag) Create(ctx context.Context, tag model.Tag) (model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tag)
	ret0, _ := ret[0].(model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTagMockRecorder) Create(ctx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTag)(nil).Create), ctx, tag)
}

// Delete mocks base method.
func (m *MockTag) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTagMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTag)(nil).Delete), ctx, id)
}

// Find mocks base method.
func (m *MockTag) Find(ctx context.Context, ids []uint64) ([]model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, ids)
	ret0, _ := ret[0].([]model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockTagMockRecorder) Find(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTag)(nil).Find), ctx, ids)
}

// Update mocks base method.
func (m *MockTag) Update(ctx context.Context, tag model.Tag) (model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tag)
	ret0, _ := ret[0].(model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTagMockRecorder) Update(ctx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTag)(nil).Update), ctx, tag)
}
