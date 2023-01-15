package handler

import (
	"net/http"

	"blog_app/usecase"

	"github.com/gin-gonic/gin"
)

type (
	PostHandler interface {
		GetPosts(c *gin.Context)
		// dbConn            *gorm.DB
		// findPosstsUsecase usecase.FindPosts
		// tagRepository     repository.Post
	}

	postHandler struct {
		findPostsUsecase usecase.FindPosts
		// createPostsUsecase usecase.CreatePosts
		// updatePostsUsecase usecase.UpdatePosts
		// deletePostsUsecase usecase.DeletePosts
	}

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

func NewPostHandler(findPostsUsecase usecase.FindPosts) PostHandler {
	return &postHandler{findPostsUsecase}
}

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
