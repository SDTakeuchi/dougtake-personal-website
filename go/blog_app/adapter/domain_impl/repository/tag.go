package repository

import (
	"blog_app/domain/model"
	postgres "blog_app/adapter/domain_impl/postgres/model"
	"blog_app/domain/repository"
	"context"
	"fmt"
)

type tagImpl struct{}

func NewTagImpl() repository.Tag {
	return &tagImpl{}
}

func (r *tagImpl) Create(ctx context.Context) ([]model.Tag, error) {}

func (r *tagImpl) Find(ctx context.Context) ([]model.Tag, error) {}

func (r *tagImpl) Update(ctx context.Context) (model.Tag, error) {
	var tag postgres.Tag
	fmt.Println(Ugh)
	return tag, nil
}

func (r *tagImpl) Delete(ctx context.Context, id uint64) error {
	return nil
}
