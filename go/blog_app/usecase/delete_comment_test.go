package usecase

import (
	mockrepo "blog_app/domain/repository/mock"
	testutil "blog_app/util/test"
	"context"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func Test_deleteCommentImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input DeleteCommentInput
	}

	randomTags := testutil.GenRandomTags(2)

	tagIDs := func() []uint64 {
		var ids []uint64
		for _, t := range randomTags {
			ids = append(ids, t.ID())
		}
		return ids
	}()
	randomPosts := testutil.GenRandomPosts(2, tagIDs)
	postIDs := func() []uint64 {
		var ids []uint64
		for _, p := range randomPosts {
			ids = append(ids, p.ID())
		}
		return ids
	}()
	randomComments := testutil.GenRandomComments(4, postIDs)

	fakeID := uint64(100000000000)

	tests := []struct {
		name              string
		args              args
		buildStubsComment func(mockComment *mockrepo.MockComment)
		want              *DeleteCommentOutput
		wantErr           bool
	}{
		{
			"success",
			args{
				context.Background(),
				DeleteCommentInput{
					randomComments[0].ID(),
				},
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					Delete(gomock.Any(), randomComments[0].ID()).
					Times(1).
					Return(
						nil,
					)
			},
			&DeleteCommentOutput{},
			false,
		},
		{
			"fail/id-not-found",
			args{
				context.Background(),
				DeleteCommentInput{
					fakeID,
				},
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					Delete(gomock.Any(), fakeID).
					Times(1).
					Return(
						gorm.ErrRecordNotFound,
					)
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockComment := mockrepo.NewMockComment(ctrl)
			tt.buildStubsComment(mockComment)

			deleteComment := NewDeleteComment(
				mockComment,
			)
			got, err := deleteComment.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateCommentImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateCommentImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
