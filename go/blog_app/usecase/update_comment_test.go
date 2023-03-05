package usecase

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/domain/model"
	mockrepo "blog_app/domain/repository/mock"
	testutil "blog_app/util/test"
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_updateCommentImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input UpdateCommentInput
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

	newCommentBodies := []string{
		testutil.GenRandomChars(100),
		testutil.GenRandomChars(150),
	}
	now := time.Now()

	c1, _ := modelimpl.NewComment(
		randomComments[0].ID(),
		newCommentBodies[0],
		randomComments[0].PostID(),
		randomComments[0].CreatedAt(),
		now,
	)
	c2, _ := modelimpl.NewComment(
		randomComments[1].ID(),
		newCommentBodies[1],
		randomComments[1].PostID(),
		randomComments[1].CreatedAt(),
		now,
	)
	commentsExpected := []model.Comment{c1, c2}

	tests := []struct {
		name              string
		args              args
		buildStubsComment func(mockComment *mockrepo.MockComment)
		want              *UpdateCommentOutput
		wantErr           bool
	}{
		{
			"success",
			args{
				context.Background(),
				UpdateCommentInput{
					randomComments[0].ID(),
					newCommentBodies[0],
				},
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					Get(gomock.Any(), randomComments[0].ID()).
					Times(1).
					Return(
						randomComments[0],
						nil,
					)
				mockComment.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Times(1).
					Return(
						commentsExpected[0],
						nil,
					)
			},
			&UpdateCommentOutput{commentsExpected[0]},
			false,
		},
		{
			"fail/no-body",
			args{
				context.Background(),
				UpdateCommentInput{
					randomPosts[0].ID(),
					"",
				},
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					Get(gomock.Any(), randomComments[0].ID()).
					Times(1).
					Return(
						randomComments[0],
						nil,
					)
				mockComment.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Times(0)
			},
			nil,
			true,
		},
		{
			"fail/post-id-not-found",
			args{
				context.Background(),
				UpdateCommentInput{
					uint64(1000000),
					randomPosts[1].Body(),
				},
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					Get(gomock.Any(), uint64(1000000)).
					Times(1).
					Return(
						nil,
						gorm.ErrRecordNotFound,
					)
				mockComment.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Times(0)
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

			updateComment := NewUpdateComment(
				mockComment,
			)
			got, err := updateComment.Execute(tt.args.ctx, tt.args.input)
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
