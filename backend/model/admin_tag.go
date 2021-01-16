package model

// AdminTagsGetModel is result of tag retrieve
type AdminTagsGetModel struct {
	Tags []string `json:"tags"`
}

// AdminTagsCreateModel is structure to create a tag
type AdminTagsCreateModel struct {
	TagName string `json:"tagName" binding:"required"`
}

// AdminTagsCreateResult is result of tag creation
type AdminTagsCreateResult struct {
	TagName string `json:"tagName"`
}
