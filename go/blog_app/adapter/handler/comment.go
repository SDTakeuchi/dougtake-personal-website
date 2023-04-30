package handler

import (
	"blog_app/usecase"

	"github.com/gin-gonic/gin"
)

type (
	CommentHandler interface {
		CreateComment(c *gin.Context)
		UpdateComment(c *gin.Context)
		DeleteComment(c *gin.Context)
	}
	commentHandler struct {
		createCommentUsecase usecase.CreateComment
		updateCommentUsecase usecase.UpdateComment
		deleteCommentUsecase usecase.DeleteComment
	}

	createCommentRequest struct {
		PostID uint64 `form:"post_id" json:"post_id"`
		Body   string `form:"body" json:"body"`
	}
	createCommentResponse struct {
		ID uint64 `json:"id"`
	}

	updateCommentRequest struct {
		ID   uint64 `form:"id" json:"id"`
		Body string `form:"body" json:"body"`
	}
	updateCommentResponse struct {
		ID uint64 `json:"id"`
	}

	deleteCommentRequest struct {
		ID uint64 `form:"id" json:"id"`
	}
	deleteCommentResponse struct {
	}
)

func NewCommentHandler(
	createCommentUsecase usecase.CreateComment,
	updateCommentResponse usecase.UpdateComment,
	deleteCommentResponse usecase.DeleteComment,
) CommentHandler {
	return &commentHandler{
		createCommentUsecase,
		updateCommentResponse,
		deleteCommentResponse,
	}
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	params := createCommentRequest{}
	if err := c.Bind(&params); err != nil {
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
		createCommentResponse{
			(*output).Comment.ID(),
		},
	)
}

func (h *commentHandler) UpdateComment(c *gin.Context) {
	params := updateCommentRequest{}
	if err := c.Bind(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	output, err := h.updateCommentUsecase.Execute(
		c,
		usecase.UpdateCommentInput{
			ID:   params.ID,
			Body: params.Body,
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	createJSONResponse(
		c,
		updateCommentResponse{
			(*output).Comment.ID(),
		},
	)
}

func (h *commentHandler) DeleteComment(c *gin.Context) {
	params := deleteCommentRequest{}
	if err := c.Bind(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	_, err := h.deleteCommentUsecase.Execute(
		c,
		usecase.DeleteCommentInput{
			ID: params.ID,
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	createJSONResponse(
		c,
		deleteCommentResponse{},
	)
}
