// Code generated by MockGen. DO NOT EDIT.
// Source: comment.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "blog_app/domain/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockComment is a mock of Comment interface.
type MockComment struct {
	ctrl     *gomock.Controller
	recorder *MockCommentMockRecorder
}

// MockCommentMockRecorder is the mock recorder for MockComment.
type MockCommentMockRecorder struct {
	mock *MockComment
}

// NewMockComment creates a new mock instance.
func NewMockComment(ctrl *gomock.Controller) *MockComment {
	mock := &MockComment{ctrl: ctrl}
	mock.recorder = &MockCommentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComment) EXPECT() *MockCommentMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockComment) Create(ctx context.Context, comment model.Comment) (model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, comment)
	ret0, _ := ret[0].(model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCommentMockRecorder) Create(ctx, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockComment)(nil).Create), ctx, comment)
}

// Delete mocks base method.
func (m *MockComment) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockComment)(nil).Delete), ctx, id)
}

// FindByPostID mocks base method.
func (m *MockComment) FindByPostID(ctx context.Context, postID uint64) ([]model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPostID", ctx, postID)
	ret0, _ := ret[0].([]model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPostID indicates an expected call of FindByPostID.
func (mr *MockCommentMockRecorder) FindByPostID(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPostID", reflect.TypeOf((*MockComment)(nil).FindByPostID), ctx, postID)
}

// Update mocks base method.
func (m *MockComment) Update(ctx context.Context, comment model.Comment) (model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, comment)
	ret0, _ := ret[0].(model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCommentMockRecorder) Update(ctx, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockComment)(nil).Update), ctx, comment)
}
