package handler

import (
	"blog_app/usecase"

	"github.com/gin-gonic/gin"
)

type (
	CommentHandler interface {
		CreateComment(c *gin.Context)
	}
	commentHandler struct {
		createCommentUsecase usecase.CreateComment
	}
	createCommentRequest struct {
		PostID uint64 `json:"post_id"`
		Body   string `json:"body"`
	}
)

func NewCommentHandler(createCommentUsecase usecase.CreateComment) CommentHandler {
	return &commentHandler{createCommentUsecase}
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	params := createCommentRequest{}
	if err := c.BindQuery(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	output, err := h.createCommentUsecase.Execute(
		c,
		usecase.CreateCommentInput{
			PostID: params.PostID,
			Body:   params.Body,
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	createJSONResponse(
		c,
		(*output).Comment.ID(),
	)
}
