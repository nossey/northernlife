package dataaccessor

import (
	"time"

	"github.com/google/uuid"
	"github.com/nossey/northernlife/infrastructure"
)

// TagDataAccessor provides functions to access tags
type TagDataAccessor struct {
}

// TagAccessor is singleton of TagDataAccessor
var TagAccessor *TagDataAccessor

func init() {
	TagAccessor = &TagDataAccessor{}
}

// GetTags get tags
func (accessor *TagDataAccessor) GetTags() (tags []string) {
	db := infrastructure.Db

	sql := `
select
	tag_name
from
	tags
`
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		tags = append(tags, tag)
	}
	return
}

// Tag is column for table tags
type Tag struct {
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	ID        uuid.UUID `gorm:"id"`
	TagName   string    `gorm:"tag_name"`
	UserID    string    `gorm:"user_id"`
}
