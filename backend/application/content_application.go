package application

import (
	"fmt"
	"os"

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

// Upload uploads a file to s3 bucket
func (app *ContentApplication) Upload() error {
	f, err := os.Open("./main.go")
	if err != nil {
		return err
	}
	defer f.Close()

	result, err := app.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("northernlife-content"),
		Key:    aws.String(uuid.New().String()),
		Body:   f,
	})
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
