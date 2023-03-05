package usecase

import (
	"blog_app/domain/model"
	modeltime "blog_app/domain/model/time"
	mockrepo "blog_app/domain/repository/mock"
	testutil "blog_app/util/test"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_findPostsImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input FindPostsInput
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

	tests := []struct {
		name              string
		args              args
		buildStubsPost    func(mockPost *mockrepo.MockPost)
		buildStubsTag     func(mockTag *mockrepo.MockTag)
		buildStubsComment func(mockComment *mockrepo.MockComment)
		want              *FindPostsOutput
		wantErr           bool
	}{
		{
			"success/get-1-post-by-id",
			args{
				context.Background(),
				FindPostsInput{ID: postIDs[1]},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Get(gomock.Any(), gomock.Eq(randomPosts[1].ID())).
					Times(1).
					Return(randomPosts[1], nil)
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Eq(randomPosts[1].TagIDs())).
					Times(1).
					Return(randomTags, nil)
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					FindByPostID(gomock.Any(), gomock.Eq(randomPosts[1].ID())).
					Times(1).
					Return(
						[]model.Comment{randomComments[2]},
						nil,
					)
			},
			&FindPostsOutput{
				[]postOutput{
					{
						randomPosts[1].ID(),
						randomPosts[1].Title(),
						randomPosts[1].Body(),
						modeltime.Time{randomPosts[1].CreatedAt()},
						modeltime.Time{randomPosts[1].UpdatedAt()},
						[]tagOutput{
							{
								randomTags[0].ID(),
								randomTags[0].Name(),
							},
							{
								randomTags[1].ID(),
								randomTags[1].Name(),
							},
						},
						[]commentOutput{
							{
								randomComments[2].ID(),
								randomComments[2].Body(),
								modeltime.Time{randomComments[2].CreatedAt()},
								modeltime.Time{randomComments[2].UpdatedAt()},
							},
						},
					},
				},
				// searching by id do not return next post index.
				0,
			},
			false,
		},
		{
			"success/get-1-post-by-search",
			args{
				context.Background(),
				FindPostsInput{SearchChar: randomPosts[0].Body()[0:100]},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Find(
						gomock.Any(),
						gomock.Any(),
						gomock.Eq(randomPosts[0].Body()[0:100]),
						gomock.Any(),
						gomock.Any(),
					).
					Times(1).
					Return([]model.Post{randomPosts[0]}, nil)
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Eq(randomPosts[0].TagIDs())).
					Times(1).
					Return(randomTags, nil)
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					FindByPostID(gomock.Any(), gomock.Eq(randomPosts[0].ID())).
					Times(1).
					Return(
						[]model.Comment{randomComments[0]},
						nil,
					)
			},
			&FindPostsOutput{
				[]postOutput{
					{
						randomPosts[0].ID(),
						randomPosts[0].Title(),
						randomPosts[0].Body(),
						modeltime.Time{randomPosts[0].CreatedAt()},
						modeltime.Time{randomPosts[0].CreatedAt()},
						[]tagOutput{
							{
								randomTags[0].ID(),
								randomTags[0].Name(),
							},
							{
								randomTags[1].ID(),
								randomTags[1].Name(),
							},
						},
						[]commentOutput{
							{
								randomComments[0].ID(),
								randomComments[0].Body(),
								modeltime.Time{randomComments[0].CreatedAt()},
								modeltime.Time{randomComments[0].CreatedAt()},
							},
						},
					},
				},
				0,
			},
			false,
		},
		{
			"success/get-1-post-by-tag-and-with-limit",
			args{
				context.Background(),
				FindPostsInput{TagID: randomTags[0].ID(), Limit: 1},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Find(
						gomock.Any(),
						randomTags[0].ID(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).
					Times(1).
					Return([]model.Post{randomPosts[0], randomPosts[1]}, nil)
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Eq(randomPosts[0].TagIDs())).
					Times(1).
					Return(randomTags, nil)
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					FindByPostID(gomock.Any(), gomock.Eq(randomPosts[0].ID())).
					Times(1).
					Return(
						[]model.Comment{randomComments[0]},
						nil,
					)
			},
			&FindPostsOutput{
				[]postOutput{
					{
						randomPosts[0].ID(),
						randomPosts[0].Title(),
						randomPosts[0].Body(),
						modeltime.Time{randomPosts[0].CreatedAt()},
						modeltime.Time{randomPosts[0].CreatedAt()},
						[]tagOutput{
							{
								randomTags[0].ID(),
								randomTags[0].Name(),
							},
							{
								randomTags[1].ID(),
								randomTags[1].Name(),
							},
						},
						[]commentOutput{
							{
								randomComments[0].ID(),
								randomComments[0].Body(),
								modeltime.Time{randomComments[0].CreatedAt()},
								modeltime.Time{randomComments[0].CreatedAt()},
							},
						},
					},
				},
				randomPosts[1].ID(),
			},
			false,
		},
		{
			"success/get-1-post-by-tag-and-with-offset",
			args{
				context.Background(),
				FindPostsInput{TagID: randomTags[0].ID(), Offset: 1, Limit: 1},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Find(
						gomock.Any(),
						randomTags[0].ID(),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
					).
					Times(1).
					Return([]model.Post{randomPosts[1]}, nil)
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Eq(randomPosts[1].TagIDs())).
					Times(1).
					Return(randomTags, nil)
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					FindByPostID(gomock.Any(), gomock.Eq(randomPosts[1].ID())).
					Times(1).
					Return(
						[]model.Comment{randomComments[2]},
						nil,
					)
			},
			&FindPostsOutput{
				[]postOutput{
					{
						randomPosts[1].ID(),
						randomPosts[1].Title(),
						randomPosts[1].Body(),
						modeltime.Time{randomPosts[1].CreatedAt()},
						modeltime.Time{randomPosts[1].CreatedAt()},
						[]tagOutput{
							{
								randomTags[0].ID(),
								randomTags[0].Name(),
							},
							{
								randomTags[1].ID(),
								randomTags[1].Name(),
							},
						},
						[]commentOutput{
							{
								randomComments[2].ID(),
								randomComments[2].Body(),
								modeltime.Time{randomComments[2].CreatedAt()},
								modeltime.Time{randomComments[2].CreatedAt()},
							},
						},
					},
				},
				0,
			},
			false,
		},
		{
			"fail/get-no-post-by-id",
			args{
				context.Background(),
				FindPostsInput{ID: uint64(100000)},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, gorm.ErrRecordNotFound)
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Any()).
					Times(0).
					Return(nil, nil)
			},
			func(mockComment *mockrepo.MockComment) {
				mockComment.EXPECT().
					FindByPostID(gomock.Any(), gomock.Any()).
					Times(0).
					Return(nil, nil)
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
			mockTag := mockrepo.NewMockTag(ctrl)
			tt.buildStubsTag(mockTag)
			mockComment := mockrepo.NewMockComment(ctrl)
			tt.buildStubsComment(mockComment)

			findPosts := NewFindPosts(
				mockPost,
				mockTag,
				mockComment,
			)

			got, err := findPosts.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPostsImpl.Execute() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPostsImpl.Execute() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
