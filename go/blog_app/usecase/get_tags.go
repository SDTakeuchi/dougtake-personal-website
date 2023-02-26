package usecase

import (
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
)

type (
	GetTags interface {
		Execute(ctx context.Context, input GetTagsInput) (*GetTagsOutput, error)
	}
	GetTagsInput struct {
		IDs []uint64
	}
	GetTagsOutput struct {
		Tags []model.Tag
	}
	getTagsImpl struct {
		tagRepo repository.Tag
	}
)

func NewGetTags(
	tagRepo repository.Tag,
) GetTags {
	return &getTagsImpl{
		tagRepo: tagRepo,
	}
}

func (u *getTagsImpl) Execute(ctx context.Context, input GetTagsInput) (*GetTagsOutput, error) {
	tags, err := u.tagRepo.Find(ctx, input.IDs)
	if err != nil {
		return nil, err
	}
	return &GetTagsOutput{tags}, nil
}
