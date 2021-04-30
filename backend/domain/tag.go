package domain

// GetTagsItem represents single item of get attached tags
type GetTagsItem struct {
	Id      string `gorm:"tag_id"`
	TagName string `gorm:"tag_name"`
	Count   int    `gorm:"count"`
}

// GetTagsResult represents results of get attached tags
type GetTagsResult struct {
	Items []GetTagsItem
}
