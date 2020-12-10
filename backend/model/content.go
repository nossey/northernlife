package model

// FileImageUploadModel is the form of image upload
type FileImageUploadModel struct {
	Image string `json:"image" binding:"required"`
}

// FileImageUploadSuccessResultModel is the result of upload
type FileImageUploadSuccessResultModel struct {
	URL string `json:"url"`
}
