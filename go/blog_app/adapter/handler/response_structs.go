package handler

type (
	tag struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	comment struct {
		ID        uint64 `json:"id"`
		Body      string `json:"body"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	post struct {
		ID        uint64    `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
		Tags      []tag     `json:"tags"`
		Comments  []comment `json:"comments"`
	}
)
