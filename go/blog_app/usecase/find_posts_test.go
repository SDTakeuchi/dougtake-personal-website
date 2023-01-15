package usecase

import (
	"blog_app/domain/repository"
	"context"
	"reflect"
	"testing"
)

func TestNewFindPosts(t *testing.T) {
	type args struct {
		postRepo    repository.Post
		tagRepo     repository.Tag
		commentRepo repository.Comment
	}
	tests := []struct {
		name string
		args args
		want FindPosts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFindPosts(tt.args.postRepo, tt.args.tagRepo, tt.args.commentRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFindPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPostsImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input FindPostsInput
	}

	ctx := context.Background()

	tests := []struct {
		name    string
		u       *findPostsImpl
		args    args
		want    *FindPostsOutput
		wantErr bool
	}{
		{
			"success/ get 1 post",
			&findPostsImpl{},
			args{
				ctx,
				FindPostsInput{},
			},
			&FindPostsOutput{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPostsImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPostsImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
