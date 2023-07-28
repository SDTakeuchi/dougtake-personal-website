package usecase

import (
	"blog_app/domain/model/uuid"
	mockrepo "blog_app/domain/repository/mock"
	testutil "blog_app/util/test"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_createPostImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input CreatePostInput
	}

	randomTags := testutil.GenRandomTags(2)

	tagIDs := func() []uint64 {
		var ids []uint64
		for _, t := range randomTags {
			ids = append(ids, t.ID())
		}
		return ids
	}()
	randomUsers := testutil.GenRandomUsers(2)
	randomPosts := testutil.GenRandomPosts(2, tagIDs)

	tests := []struct {
		name           string
		args           args
		buildStubsUser func(mockPost *mockrepo.MockUser)
		buildStubsTag  func(mockPost *mockrepo.MockTag)
		buildStubsPost func(mockPost *mockrepo.MockPost)
		want           *CreatePostOutput
		wantErr        bool
	}{
		{
			name: "success/1",
			args: args{
				ctx: context.Background(),
				input: CreatePostInput{
					UserID: randomUsers[0].ID().String(),
					Title:  randomPosts[0].Title(),
					Body:   randomPosts[0].Body(),
					TagIDs: randomPosts[0].TagIDs(),
				},
			},
			buildStubsUser: func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					GetByID(
						gomock.Any(),
						randomUsers[0].ID(),
					).
					Times(1).
					Return(randomUsers[0], nil)
			},
			buildStubsTag: func(mockUser *mockrepo.MockTag) {
				mockUser.EXPECT().
					Find(
						gomock.Any(),
						gomock.Any(),
					).
					Times(1).
					Return(randomTags, nil)
			},
			buildStubsPost: func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(
						randomPosts[0],
						nil,
					)
			},
			want:    &CreatePostOutput{randomPosts[0]},
			wantErr: false,
		},
		{
			name: "success/2",
			args: args{
				ctx: context.Background(),
				input: CreatePostInput{
					UserID: randomUsers[1].ID().String(),
					Title:  randomPosts[1].Title(),
					Body:   randomPosts[1].Body(),
					TagIDs: randomPosts[1].TagIDs(),
				},
			},
			buildStubsUser: func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					GetByID(
						gomock.Any(),
						randomUsers[1].ID(),
					).
					Times(1).
					Return(randomUsers[1], nil)
			},
			buildStubsTag: func(mockUser *mockrepo.MockTag) {
				mockUser.EXPECT().
					Find(
						gomock.Any(),
						gomock.Any(),
					).
					Times(1).
					Return(randomTags, nil)
			},
			buildStubsPost: func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(
						randomPosts[1],
						nil,
					)
			},
			want:    &CreatePostOutput{Post: randomPosts[1]},
			wantErr: false,
		},
		{
			name: "fail/tag_id_not_found",
			args: args{
				ctx: context.Background(),
				input: CreatePostInput{
					UserID: randomUsers[1].ID().String(),
					Title:  randomPosts[1].Title(),
					Body:   randomPosts[1].Body(),
					TagIDs: []uint64{uint64(100000000000000)},
				},
			},
			buildStubsUser: func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					GetByID(
						gomock.Any(),
						randomUsers[1].ID(),
					).
					Times(1).
					Return(randomUsers[1], nil)
			},
			buildStubsTag: func(mockUser *mockrepo.MockTag) {
				mockUser.EXPECT().
					Find(
						gomock.Any(),
						gomock.Any(),
					).
					Times(1).
					Return(nil, gorm.ErrRecordNotFound)
			},
			buildStubsPost: func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail/user_id_not_found",
			args: args{
				ctx: context.Background(),
				input: CreatePostInput{
					UserID: uuid.New().String(),
					Title:  randomPosts[0].Title(),
					Body:   randomPosts[0].Body(),
					TagIDs: randomPosts[0].TagIDs(),
				},
			},
			buildStubsUser: func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					GetByID(
						gomock.Any(),
						gomock.Any(),
					).
					Times(1).
					Return(nil, gorm.ErrRecordNotFound)
			},
			buildStubsTag: func(mockUser *mockrepo.MockTag) {
				mockUser.EXPECT().
					Find(
						gomock.Any(),
						gomock.Any(),
					).
					Times(0)
			},
			buildStubsPost: func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail/empty_title",
			args: args{
				ctx: context.Background(),
				input: CreatePostInput{
					UserID: randomUsers[0].ID().String(),
					Title:  "",
					Body:   randomPosts[0].Body(),
					TagIDs: randomPosts[0].TagIDs(),
				},
			},
			buildStubsUser: func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					GetByID(
						gomock.Any(),
						gomock.Any(),
					).
					Times(1).
					Return(nil, gorm.ErrRecordNotFound)
			},
			buildStubsTag: func(mockUser *mockrepo.MockTag) {
				mockUser.EXPECT().
					Find(
						gomock.Any(),
						gomock.Any(),
					).
					Times(0)
			},
			buildStubsPost: func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockUser := mockrepo.NewMockUser(ctrl)
			tt.buildStubsUser(mockUser)
			mockTag := mockrepo.NewMockTag(ctrl)
			tt.buildStubsTag(mockTag)
			mockPost := mockrepo.NewMockPost(ctrl)
			tt.buildStubsPost(mockPost)

			createPost := NewCreatePost(
				mockUser,
				mockTag,
				mockPost,
			)
			got, err := createPost.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("createPostImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPostImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
