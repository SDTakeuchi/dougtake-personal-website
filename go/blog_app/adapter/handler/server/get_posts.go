package server

import (
	"blog_app/domain/repository"
	"blog_app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type GetPostsProvider struct {
	dbConn            *gorm.DB
	findPosstsUsecase usecase.FindPosts
	tagRepository     repository.Post
}

func NewGetPostsProvider(
	dbConn *gorm.DB,
	findPosstsUsecase usecase.FindPosts,
	tagRepository repository.Post,
) *GetPostsProvider {
	return &GetPostsProvider{
		dbConn:            dbConn,
		findPosstsUsecase: findPosstsUsecase,
		tagRepository:     tagRepository,
	}
}

type (
	GetPostsRequest struct {
		TagId      uint64 `json:"tag_id"`
		searchChar string `json:"search_char"`
		pageIndex  uint64 `json:"page_index"`
		pageSize   uint64 `json:"page_size"`
	}

	GetPostsResponse struct {
		Posts usecase.FindPostsOutput `json:"posts"`
	}
)

func (p *GetPostsProvider) GetPosts(c *gin.Context) {

}
