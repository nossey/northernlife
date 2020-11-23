package application

import (
	"time"

	"github.com/google/uuid"
	"github.com/nossey/northernlife/dataaccessor"
	"github.com/nossey/northernlife/infrastructure"
)

// GetTags get all tags
func GetTags() (tags []string) {
	tags = []string{}
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

// CreateTag creates a tag
func CreateTag(tagName string, userID string) (err error) {
	db := infrastructure.Db
	tag := dataaccessor.Tag{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        uuid.New(),
		TagName:   tagName,
		UserID:    userID,
	}
	err = db.Create(&tag).Error
	return
}
