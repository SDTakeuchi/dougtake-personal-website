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
	posts := make([]model.Post, wantCount)
	for i := 0; i < wantCount; i++ {
		p := postgres.Post{
			ID:        uint64(i + 1),
			Title:     GenRandomChars(30),
			Body:      GenRandomChars(2000),
			UserID:    uuid.New(),
			TagIDs:    tagIDs,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		posts[i] = modelimpl.PostFromRecord(p)
	}
	return posts
}

func GenRandomTags(wantCount int) []model.Tag {
	tags := make([]model.Tag, wantCount)
	for i := 0; i < wantCount; i++ {
		t := postgres.Tag{
			ID:   uint64(i + 1),
			Name: GenRandomChars(10),
		}
		tags[i] = modelimpl.TagFromRecord(t)
	}
	return tags
}

var SamplePassword = "a9#jL0s8hbFSiolk"

func GenRandomUsers(wantCount int) []model.User {
	users := make([]model.User, wantCount)
	samplePassword, _ := password.NewPassword(SamplePassword)
	hashedPassword := samplePassword.HashedPassword()
	for i := 0; i < wantCount; i++ {
		email := fmt.Sprintf("%sABC@email.com", GenRandomChars(20))
		u := postgres.User{
			ID:       uuid.New(),
			Name:     GenRandomChars(10),
			Email:    email,
			Password: hashedPassword,
		}
		users[i] = modelimpl.UserFromRecord(u)
	}
	return users
}

func GenRandomComments(wantCount int, postIDs []uint64) []model.Comment {
	comments := make([]model.Comment, wantCount)
	for i := 0; i < wantCount; i++ {

		postID := len(postIDs)
		if i < len(postIDs) {
			postID = i+1
		}

		pgC := postgres.Comment{
			ID:        uint64(i+1),
			Body:      GenRandomChars(200),
			PostID:    uint64(postID),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		comments[i] = modelimpl.CommentFromRecord(pgC)
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
