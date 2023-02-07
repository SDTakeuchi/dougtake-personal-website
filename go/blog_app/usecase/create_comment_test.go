package usecase

import (
	mockrepo "blog_app/domain/repository/mock"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_createCommentImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input CreateCommentInput
	}

	randomTags := genRandomTags(2)

	tagIDs := func() []uint64 {
		var ids []uint64
		for _, t := range randomTags {
			ids = append(ids, t.ID())
		}
		return ids
	}()
	randomPosts := genRandomPosts(2, tagIDs)
	postIDs := func() []uint64 {
		var ids []uint64
		for _, p := range randomPosts {
			ids = append(ids, p.ID())
		}
		return ids
	}()
	randomComments := genRandomComments(4, postIDs)

	tests := []struct {
		name              string
		args              args
		buildStubsPost    func(mockPost *mockrepo.MockPost)
		buildStubsComment func(mockComment *mockrepo.MockComment)
		want              *CreateCommentOutput
		wantErr           bool
	}{
		{
			"success",
			args{
				context.Background(),
				CreateCommentInput{
					randomPosts[0].ID(),
					randomComments[0].Body(),
				},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Get(
						gomock.Any(),
						randomPosts[0].ID(),
					).
					Times(1).
					Return(randomPosts[0], nil)
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(
						randomComments[0],
						nil,
					)
			},
			&CreateCommentOutput{randomComments[0]},
			false,
		},
		{
			"fail/no-body",
			args{
				context.Background(),
				CreateCommentInput{
					randomPosts[0].ID(),
					"",
				},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Get(
						gomock.Any(),
						randomPosts[0].ID(),
					).
					Times(1).
					Return(randomPosts[0], nil)
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			nil,
			true,
		},
		{
			"fail/post-id-not-found",
			args{
				context.Background(),
				CreateCommentInput{
					uint64(1000000),
					randomPosts[1].Body(),
				},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Get(
						gomock.Any(),
						uint64(1000000),
					).
					Times(1).
					Return(nil, gorm.ErrRecordNotFound)
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockPost := mockrepo.NewMockPost(ctrl)
			tt.buildStubsPost(mockPost)
			mockComment := mockrepo.NewMockComment(ctrl)
			tt.buildStubsComment(mockComment)

			createComment := NewCreateComment(
				mockPost,
				mockComment,
			)
			usecaseImpl, _ := createComment.(*createCommentImpl)
			got, err := usecaseImpl.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("createCommentImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createCommentImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
