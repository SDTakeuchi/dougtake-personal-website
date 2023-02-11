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

	getPostsResponse struct {
		posts usecase.FindPostsOutput
	}
)

func NewPostHandler(findPostsUsecase usecase.FindPosts) PostHandler {
	return &postHandler{findPostsUsecase}
}

func (h *postHandler) GetPosts(c *gin.Context) {
	params := GetPostsRequest{}
	if err := c.BindQuery(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	output, err := h.findPostsUsecase.Execute(c, usecase.FindPostsInput{
		ID:         params.ID,
		TagID:      params.TagID,
		SearchChar: params.SearchChar,
		Offset:     params.PageIndex,
		Limit:      params.PageSize,
	})
	if err != nil {
		createErrResponse(c, err)
		return
	}
	createJSONResponse(
		c,
		http.StatusOK,
		getPostsResponse{
			*output,
		},
	)
}
