package usecase

import (
	"blog_app/domain/repository"
	"context"
	"reflect"
	"testing"
)

func Test_deleteCommentImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input DeleteCommentInput
	}
	tests := []struct {
		name    string
		u       *deleteCommentImpl
		args    args
		want    *DeleteCommentOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteCommentImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteCommentImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
