package handler

import (
	"blog_app/usecase"

	"github.com/gin-gonic/gin"
)

type (
	TagHandler interface {
		GetTags(c *gin.Context)
	}
	tagHandler struct {
		getTagsUsecase usecase.GetTags
	}
	getTagsRequest struct {
		IDs []uint64 `form:"ids"`
	}
	getTagsResponse struct {
		Tags []tag `json:"tags"`
	}
)

func NewTagHandler(getTagsUsecase usecase.GetTags) TagHandler {
	return &tagHandler{getTagsUsecase}
}

func (h *tagHandler) GetTags(c *gin.Context) {
	params := getTagsRequest{}
	if err := c.BindQuery(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	output, err := h.getTagsUsecase.Execute(
		c,
		usecase.GetTagsInput{
			IDs: params.IDs,
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	var resp getTagsResponse
	for _, t := range output.Tags {
		resp.Tags = append(resp.Tags, tag{t.ID(), t.Name()})
	}
	createJSONResponse(c, resp)
}
