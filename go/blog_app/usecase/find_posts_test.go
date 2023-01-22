package usecase

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/model/uuid"
	mockrepo "blog_app/domain/repository/mock"
	"context"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

// func TestNewFindPosts(t *testing.T) {
// 	type args struct {
// 		postRepo    repository.Post
// 		tagRepo     repository.Tag
// 		commentRepo repository.Comment
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want FindPosts
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewFindPosts(tt.args.postRepo, tt.args.tagRepo, tt.args.commentRepo); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewFindPosts() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_findPostsImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input FindPostsInput
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
		name string
		args              args
		buildStubsPost    func(mockPost *mockrepo.MockPost)
		buildStubsTag     func(mockTag *mockrepo.MockTag)
		buildStubsComment func(mockComment *mockrepo.MockComment)
		want              *FindPostsOutput
		wantErr           bool
	}{
		{
			"success/ get 1 post",
			args{
				context.Background(),
				FindPostsInput{ID: postIDs[0]},
			},
			func(mockPost *mockrepo.MockPost) {
				mockPost.EXPECT().
					Get(gomock.Any(), gomock.Eq(randomPosts[0].ID())).
					Times(1).
					Return(randomPosts[0], nil)
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
						randomPosts[0].CreatedAt(),
						randomPosts[0].UpdatedAt(),
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
								randomComments[0].CreatedAt(),
								randomComments[0].UpdatedAt(),
							},
						},
					},
				},
			},
			false,
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
			usecaseImpl, _ := findPosts.(*findPostsImpl)

			got, err := usecaseImpl.Execute(tt.args.ctx, tt.args.input)
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

func genRandomPosts(wantCount int, tagIDs []uint64) []model.Post {
	var posts []model.Post
	for i := 1; i < wantCount+1; i++ {
		p := postgres.Post{
			uint64(i),
			genRandomChars(30),
			genRandomChars(400),
			uuid.New(),
			tagIDs,
			time.Now(),
			time.Now(),
		}
		posts = append(posts, modelimpl.PostFromRecord(p))
	}
	return posts
}

func genRandomTags(wantCount int) []model.Tag {
	var tags []model.Tag
	for i := 1; i < wantCount+1; i++ {
		t := postgres.Tag{
			uint64(i),
			genRandomChars(10),
		}
		tags = append(tags, modelimpl.TagFromRecord(t))
	}
	return tags
}

func genRandomComments(wantCount int, postIDs []uint64) []model.Comment {
	var comments []model.Comment
	for i := 1; i < wantCount+1; i++ {

		postID := len(postIDs)
		if i < len(postIDs) {
			postID = i
		}

		c := postgres.Comment{
			uint64(i),
			genRandomChars(150),
			uint64(postID),
			time.Now(),
			time.Now(),
		}
		comments = append(comments, modelimpl.CommentFromRecord(c))
	}
	return comments
}

func genRandomChars(count int) string {
	var res string
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", ",", ".", "あ", "ん"}
	for i := 0; i < count; i++ {
		res += chars[rand.Intn(len(chars))]
	}
	return res
}
