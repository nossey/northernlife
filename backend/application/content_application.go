package application

import (
	"bytes"
	"encoding/base64"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

// ContentApplication manages file uploads of northernlife
type ContentApplication struct {
	Session  *session.Session
	Uploader *s3manager.Uploader
}

// ContentApp is singleton of ContentApplication
var ContentApp *ContentApplication

func init() {
	ContentApp = &ContentApplication{}
	ContentApp.Session = session.Must(session.NewSession())
	ContentApp.Uploader = s3manager.NewUploader(ContentApp.Session)
}

// UploadResult represents result type of upload
type UploadResult int

const (
	// Success represents success of upload
	Success UploadResult = iota
	// InvalidEncodedImage represents invalid image input
	InvalidEncodedImage
	// S3UploadFailed represents failure of s3 upload
	S3UploadFailed
)

// Upload uploads a file to s3 bucket
func (app *ContentApplication) Upload(base64EncodedImage string) (id string, result UploadResult) {
	id = ""

	dec, err := base64.StdEncoding.DecodeString(base64EncodedImage)
	if err != nil {
		result = InvalidEncodedImage
		return
	}
	r := bytes.NewReader(dec)

	id = strings.Replace(uuid.New().String(), "-", "", -1)
	_, err = app.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("northernlife-content"),
		Key:    aws.String(id),
		Body:   r,
	})
	if err != nil {
		id = ""
		result = S3UploadFailed
		return
	}
	result = Success
	return
}
