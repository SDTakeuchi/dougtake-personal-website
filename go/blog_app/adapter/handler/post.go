package handler

import (
	"blog_app/usecase"

	"github.com/gin-gonic/gin"
)

type (
	PostHandler interface {
		GetPosts(c *gin.Context)
		CreatePost(c *gin.Context)
	}
	postHandler struct {
		createPostUsecase usecase.CreatePost
		findPostsUsecase  usecase.FindPosts
	}

	createPostRequest struct {
		UserID string   `json:"user_id"`
		Title  string   `json:"title"`
		Body   string   `json:"body"`
		TagIDs []uint64 `json:"tag_ids"`
	}
	createPostResponse struct {
		ID uint64 `json:"post_id"`
	}

	getPostsRequest struct {
		ID         uint64 `form:"id"`
		TagID      uint64 `form:"tag_id"`
		SearchChar string `form:"search_char"`
		PageIndex  uint64 `form:"page_index"`
		PageSize   uint64 `form:"page_size"`
	}
	getPostsResponse struct {
		Posts         []post `json:"posts"`
		NextPostIndex uint64 `json:"next_post_index"`
	}
)

func NewPostHandler(
	findPostsUsecase usecase.FindPosts,
	createPostUsecase usecase.CreatePost,
) PostHandler {
	return &postHandler{
		createPostUsecase: createPostUsecase,
		findPostsUsecase:  findPostsUsecase,
	}
}

func (h *postHandler) CreatePost(c *gin.Context) {
	params := createPostRequest{}
	if err := c.Bind(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}

	output, err := h.createPostUsecase.Execute(
		c,
		usecase.CreatePostInput{
			UserID: params.UserID,
			Title:  params.Title,
			Body:   params.Body,
			TagIDs: params.TagIDs,
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	createJSONResponse(
		c,
		createPostResponse{
			output.Post.ID(),
		},
	)
}

func (h *postHandler) GetPosts(c *gin.Context) {
	var params getPostsRequest
	if err := c.BindQuery(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}

	output, err := h.findPostsUsecase.Execute(
		c,
		usecase.FindPostsInput{
			ID:         params.ID,
			TagID:      params.TagID,
			SearchChar: params.SearchChar,
			Offset:     params.PageIndex,
			Limit:      params.PageSize,
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	resp := getPostsResponse{
		NextPostIndex: (*output).NextPostIndex,
	}
	for _, p := range (*output).Posts {

		var tags []tag
		for _, t := range p.Tags {
			tags = append(
				tags,
				tag{
					ID:   t.ID,
					Name: t.Name,
				},
			)
		}

		var comments []comment
		for _, c := range p.Comments {
			comments = append(
				comments,
				comment{
					ID:        c.ID,
					Body:      c.Body,
					CreatedAt: c.CreatedAt.StringHour(),
					UpdatedAt: c.UpdatedAt.StringHour(),
				},
			)
		}

		resp.Posts = append(resp.Posts, post{
			ID:        p.ID,
			Title:     p.Title,
			Body:      p.Body,
			CreatedAt: p.CreatedAt.StringDay(),
			UpdatedAt: p.UpdatedAt.StringDay(),
			Tags:      tags,
			Comments:  comments,
		})
	}
	createJSONResponse(
		c,
		resp,
	)
}
