package repository

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
)

type tagRepository struct {
	db *postgres.DB
}

func NewTagRepository(db *postgres.DB) repository.Tag {
	return &tagRepository{db: db}
}

func (r *tagRepository) Create(ctx context.Context, tag model.Tag) (model.Tag, error) {
	record := modelimpl.TagToRecord(tag)
	if err := r.db.Conn.WithContext(ctx).Create(&record).Error; err != nil {
		return nil, err
	}
	return modelimpl.TagFromRecord(record), nil
}

func (r *tagRepository) Find(ctx context.Context, ids []uint64) ([]model.Tag, error) {
	pTags := []postgres.Tag{}

	q := r.db.Conn.WithContext(ctx).Order("name")

	if len(ids) > 0 {
		// TIPS: surround "?" with patentheses when using IN query
		q = q.Where("id IN (?)", ids)
	}

	if err := q.Find(&pTags).Error; err != nil {
		return nil, err
	}

	mTags := make([]model.Tag, len(pTags))
	for i, v := range pTags {
		mTags[i] = modelimpl.TagFromRecord(v)
	}
	return mTags, nil
}

func (r *tagRepository) Update(ctx context.Context, tag model.Tag) (model.Tag, error) {
	record := modelimpl.TagToRecord(tag)
	if err := r.db.Conn.WithContext(ctx).Updates(&record).Error; err != nil {
		return nil, err
	}
	return modelimpl.TagFromRecord(record), nil
}

func (r *tagRepository) Delete(ctx context.Context, id uint64) error {
	record := postgres.Tag{}
	if err := r.db.Conn.WithContext(ctx).Take(&record, id).Error; err != nil {
		return err
	}
	return r.db.Conn.Delete(&record).Error
}
