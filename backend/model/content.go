package model

// FileImageUploadModel is the form of image upload
type FileImageUploadModel struct {
	Image string `json:"image:"required"`
}

// FileImageUploadSuccessResultModel is the result of upload
type FileImageUploadSuccessResultModel struct {
	URL string `json:"url"`
}
