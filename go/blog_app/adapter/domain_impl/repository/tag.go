package repository

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) repository.Tag {
	return &tagRepository{db: db}
}

func (r *tagRepository) Create(ctx context.Context, tag model.Tag) (model.Tag, error) {
	record := modelimpl.TagToRecord(tag)
	if err := r.db.Create(&record).Error; err != nil {
		return nil, err
	}
	return modelimpl.TagFromRecord(record), nil
}

func (r *tagRepository) Find(ctx context.Context, ids []uint64) ([]model.Tag, error) {
	records := []postgres.Tag{}

	q := r.db.Order("name")

	if len(ids) > 0 {
		// TIPS: surround "?" when using IN query
		q = q.Where("id IN (?)", ids)
	}

	if err := q.Find(&records).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	models := make([]model.Tag, len(records))
	for i, v := range records {
		models[i] = modelimpl.TagFromRecord(v)
	}
	return models, nil
}

func (r *tagRepository) Update(ctx context.Context, tag model.Tag) (model.Tag, error) {
	record := modelimpl.TagToRecord(tag)
	if err := r.db.Updates(&record).Error; err != nil {
		return nil, err
	}
	return modelimpl.TagFromRecord(record), nil
}

func (r *tagRepository) Delete(ctx context.Context, id uint64) error {
	record := postgres.Tag{}
	if err := r.db.Take(&record, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&record).Error
}
