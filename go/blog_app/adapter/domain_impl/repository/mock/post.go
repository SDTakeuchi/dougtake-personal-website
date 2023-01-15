// Code generated by MockGen. DO NOT EDIT.
// Source: post.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "blog_app/domain/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPost is a mock of Post interface.
type MockPost struct {
	ctrl     *gomock.Controller
	recorder *MockPostMockRecorder
}

// MockPostMockRecorder is the mock recorder for MockPost.
type MockPostMockRecorder struct {
	mock *MockPost
}

// NewMockPost creates a new mock instance.
func NewMockPost(ctrl *gomock.Controller) *MockPost {
	mock := &MockPost{ctrl: ctrl}
	mock.recorder = &MockPostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPost) EXPECT() *MockPostMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPost) Create(ctx context.Context, post model.Post) (model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, post)
	ret0, _ := ret[0].(model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostMockRecorder) Create(ctx, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPost)(nil).Create), ctx, post)
}

// Delete mocks base method.
func (m *MockPost) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPost)(nil).Delete), ctx, id)
}

// Find mocks base method.
func (m *MockPost) Find(ctx context.Context, tagID uint64, searchChar string, offset, limit uint64) ([]model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, tagID, searchChar, offset, limit)
	ret0, _ := ret[0].([]model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockPostMockRecorder) Find(ctx, tagID, searchChar, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPost)(nil).Find), ctx, tagID, searchChar, offset, limit)
}

// Get mocks base method.
func (m *MockPost) Get(ctx context.Context, id uint64) (model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPostMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPost)(nil).Get), ctx, id)
}

// Update mocks base method.
func (m *MockPost) Update(ctx context.Context, post model.Post) (model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, post)
	ret0, _ := ret[0].(model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPostMockRecorder) Update(ctx, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPost)(nil).Update), ctx, post)
}
