package handler

import (
	"net/http"

	"blog_app/usecase"

	"github.com/gin-gonic/gin"
)

type (
	CommentHandler interface {
		CreateComment(c *gin.Context)
		// dbConn            *gorm.DB
		// findPosstsUsecase usecase.FindPosts
		// tagRepository     repository.Post
	}

	commentHandler struct {
		createCommentUsecase usecase.CreateComment
		// updateCommentUsecase usecase.UpdateComment
		// deleteCommentUsecase usecase.DeleteComment
	}

	CreateCommentRequest struct {
		PostID uint64 `json:"post_id"`
		Body   string `json:"body"`
	}

	createCommentResponse struct {
		commentID uint64
	}
)

func NewCommentHandler(createCommentUsecase usecase.CreateComment) CommentHandler {
	return &commentHandler{createCommentUsecase}
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	params := CreateCommentRequest{}
	if err := c.BindQuery(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	output, err := h.createCommentUsecase.Execute(c, usecase.CreateCommentInput{
		PostID: params.PostID,
		Body:   params.Body,
	})
	if err != nil {
		createErrResponse(c, err)
		return
	}
	createJSONResponse(
		c,
		http.StatusOK,
		createCommentResponse{
			output.Comment.ID(),
		},
	)
}
