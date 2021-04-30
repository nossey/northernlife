package model

// TagsGetItem represents tag attached to posts
type TagsGetItem struct {
	Name          string `json:"name"`
	AttachedCount int    `json:"attachedCount"`
}

// TagsGetResult represents tags attached to posts
type TagsGetResult struct {
	Tags []TagsGetItem `json:"tags"`
}

// TagCreatedResult represents results of tag creation
type TagCreatedResult struct {
	TagName string `json:"tagName"`
}

// TagCreateRequest to create a tag
type TagCreateRequest struct {
	TagName string `json:"tagName"`
}

// TagCreateFailResult is returned when tag creation failed
type TagCreateFailResult struct {
	Message string `json:"message"`
}
