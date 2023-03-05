package test

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/model/password"
	"blog_app/domain/model/uuid"
	"fmt"
	"math/rand"
	"time"
)

func GenRandomPosts(wantCount int, tagIDs []uint64) []model.Post {
	var posts []model.Post
	for i := 1; i < wantCount+1; i++ {
		p := postgres.Post{
			ID:        uint64(i),
			Title:     GenRandomChars(30),
			Body:      GenRandomChars(2000),
			UserID:    uuid.New(),
			TagIDs:    tagIDs,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		posts = append(posts, modelimpl.PostFromRecord(p))
	}
	return posts
}

func GenRandomTags(wantCount int) []model.Tag {
	var tags []model.Tag
	for i := 1; i < wantCount+1; i++ {
		t := postgres.Tag{
			ID:   uint64(i),
			Name: GenRandomChars(10),
		}
		tags = append(tags, modelimpl.TagFromRecord(t))
	}
	return tags
}

var SamplePassword = "a9#jL0s8hbFSiolk"

func GenRandomUsers(wantCount int) []model.User {
	var users []model.User
	samplePassword, _ := password.NewPassword(SamplePassword)
	hashedPassword := samplePassword.HashedPassword()
	for i := 1; i < wantCount+1; i++ {
		email := fmt.Sprintf("%sABC@email.com", GenRandomChars(20))
		u := postgres.User{
			ID:       uuid.New(),
			Name:     GenRandomChars(10),
			Email:    email,
			Password: hashedPassword,
		}
		users = append(users, modelimpl.UserFromRecord(u))
	}
	return users
}

func GenRandomComments(wantCount int, postIDs []uint64) []model.Comment {
	var comments []model.Comment
	for i := 1; i < wantCount+1; i++ {

		postID := len(postIDs)
		if i < len(postIDs) {
			postID = i
		}

		pgC := postgres.Comment{
			ID:        uint64(i),
			Body:      GenRandomChars(200),
			PostID:    uint64(postID),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		comments = append(comments, modelimpl.CommentFromRecord(pgC))
	}
	return comments
}

func GenRandomChars(count int) string {
	var res string
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", ",", ".", "あ", "ん"}
	for i := 0; i < count; i++ {
		res += chars[rand.Intn(len(chars))]
	}
	return res
}
