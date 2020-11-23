package model

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
