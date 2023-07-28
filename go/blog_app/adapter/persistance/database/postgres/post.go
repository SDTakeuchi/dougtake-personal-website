package postgres

import (
	"blog_app/domain/model/uuid"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        uint64
	Title     string
	Body      string
	UserID    uuid.UUID
	TagIDs    pq.Int64Array `gorm:"column:tag_ids;type:bigint[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
