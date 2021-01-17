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

// TagCreate is structure to create a tag
type TagCreate struct {
	TagName string
	UserID  string
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

// GetAttachedTags get tags attached to posts
func (accessor *TagDataAccessor) GetAttachedTags() (tags []string) {
	db := infrastructure.Db
	tags = []string{}
	sql := `
with attached_tag_ids as (
	select distinct 
		tag_id
	from
		tags_posts_attachments tpa
)
select 
	t.tag_name
from
	tags t
inner join
	attached_tag_ids
	on t.id = attached_tag_ids.tag_id;	
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

// CreateTag create a tag
func (accessor *TagDataAccessor) CreateTag(create TagCreate) (err error) {
	db := infrastructure.Db
	createSQL := `
insert into
	tags
(
	created_at, 
	updated_at, 
	id,
	tag_name, 
	user_id
)
values
(
	current_timestamp,
	current_timestamp,
	uuid_generate_v4(),
	$1,
	$2
)
`
	err = db.Exec(createSQL, create.TagName, create.UserID).Error
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
