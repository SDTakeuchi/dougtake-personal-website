package usecase

import (
	"blog_app/domain/model"
	"context"
)

type (
	FindTagOutput struct {
		Tags []model.Tag
		Err  error
	}
)

func Execute(ctx context.Context) FindTagOutput {

}
