package server

import (
	"blog_app/domain/repository"
	"blog_app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	GetPosts(c *gin.Context)
	// dbConn            *gorm.DB
	// findPosstsUsecase usecase.FindPosts
	// tagRepository     repository.Post
}

type postHandler struct {
	findPostsUsecase usecase.FindPosts
	// createPostsUsecase usecase.CreatePosts
	// updatePostsUsecase usecase.UpdatePosts
	// deletePostsUsecase usecase.DeletePosts
}

func NewPostHandler(
	postRepo repository.Post,
	tagRepo repository.Tag,
	commentRepo repository.Comment,
) PostHandler {
	return &postHandler{
		findPostsUsecase: usecase.NewFindPosts(postRepo, tagRepo, commentRepo),
	}
}

type (
	GetPostsRequest struct {
		ID         uint64 `json:"id"`
		TagID      uint64 `json:"tag_id"`
		SearchChar string `json:"search_char"`
		PageIndex  uint64 `json:"page_index"`
		PageSize   uint64 `json:"page_size"`
	}

	GetPostsResponse struct {
		Posts usecase.FindPostsOutput `json:"posts"`
	}
)

func (p *postHandler) GetPosts(c *gin.Context) {
	param := GetPostsRequest{}
	// TODO: fail this to know what kind of error we get, and add it to createErrResponse's switch-cases
	if err := c.BindQuery(&param); err != nil {
		createErrResponse(c, err)
	}
	output, err := p.findPostsUsecase.Execute(c, usecase.FindPostsInput{
		ID:         param.ID,
		TagID:      param.TagID,
		SearchChar: param.SearchChar,
		Offset:     param.PageIndex,
		Limit:      param.PageSize,
	})
	if err != nil {
		createErrResponse(c, err)
	}
	createJSONResponse(c, http.StatusOK, *output)
}
