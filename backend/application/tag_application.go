package application

import (
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
