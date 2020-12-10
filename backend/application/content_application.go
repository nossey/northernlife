package application

import (
	"bytes"
	"encoding/base64"
	"errors"
	"mime"
	"regexp"
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

// ParseDataURIBase64 parse uri base64 encoded data
func ParseDataURIBase64(base64Encoded string) (fileExtension string, rawBase64Data string, err error) {
	fileExtension = ""
	rawBase64Data = base64Encoded

	match, err := regexp.MatchString(`data:\w+\/\w+;base64,+`, base64Encoded)
	if !match {
		err = errors.New("Doesn't match data uri base64 format (e.g. data:application/json;base64,XXXXX...)")
		return
	}
	colonIdx := strings.Index(base64Encoded, ":")
	semicolonIdx := strings.Index(base64Encoded, ";")
	mimeType := base64Encoded[colonIdx+1 : semicolonIdx]
	extensions, err := mime.ExtensionsByType(mimeType)
	if extensions == nil {
		err = errors.New("Invalid mime type")
		return
	}

	// SPEC:0番目のextensionを採用する
	fileExtension = extensions[0]
	commaIdx := strings.Index(base64Encoded, ",")
	rawBase64Data = rawBase64Data[commaIdx+1:]
	return
}

// Upload uploads a file to s3 bucket
func (app *ContentApplication) Upload(base64EncodedImage string) (url string, result UploadResult) {
	url = ""

	extension, raw, err := ParseDataURIBase64(base64EncodedImage)
	if err != nil {
		result = InvalidEncodedImage
		return
	}
	dec, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		result = InvalidEncodedImage
		return
	}

	r := bytes.NewReader([]byte(dec))

	id := strings.Replace(uuid.New().String(), "-", "", -1)
	key := id + extension
	_, err = app.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("northernlife-content"),
		Key:    aws.String(key),
		Body:   r,
	})
	if err != nil {
		url = ""
		result = S3UploadFailed
		return
	}

	url = "https://northernlife-content.net/" + key
	result = Success
	return
}
