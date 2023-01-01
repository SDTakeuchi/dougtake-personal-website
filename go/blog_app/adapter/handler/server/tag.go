package server

import (
	"blog_app/domain/repository"
	"github.com/jinzhu/gorm"
)

type GetTagsProvider struct {
	dbConn         *gorm.DB
	getTagsUsecase usecase.GetTags
	tagRepository  repository.Tag
}

func NewGetTagsProvider(
	dbConn *gorm.DB,
	getTagsUsecase usecase.GetTags,
	tagRepository repository.Tag,
) *GetTagsProvider {
	return &GetTagsProvider{
		dbConn:         dbConn,
		getTagsUsecase: getTagsUsecase,
		tagRepository:  tagRepository,
	}
}

func (p *GetTagsProvider) GetTags(c *gin.Context) {

}
