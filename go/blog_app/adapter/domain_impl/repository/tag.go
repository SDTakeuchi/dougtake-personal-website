package repository

import (
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"blog_app/adapter/persistance/database/postgres"
	modelimpl "blog_app/adapter/domain_impl/model"
	"context"
)

type tagImpl struct{}

func NewTagImpl() repository.Tag {
	return &tagImpl{}
}

func (r *tagImpl) Create(ctx context.Context, tag model.Tag) (model.Tag, error) {
	var record = modelimpl.TagToRecord(tag)
	return modelimpl.TagFromRecord(record), nil
}

func (r *tagImpl) Find(ctx context.Context) ([]model.Tag, error) {
	var records = []postgres.Tag{}

	var models = make([]model.Tag, len(records))
	for i, v := range records {
		models[i] = modelimpl.TagFromRecord(v)
	}
	return models, nil
}

func (r *tagImpl) Update(ctx context.Context, tag model.Tag) (model.Tag, error) {
	var record = modelimpl.TagToRecord(tag)
	return modelimpl.TagFromRecord(record), nil
}

func (r *tagImpl) Delete(ctx context.Context, id uint64) error {
	return nil
}
