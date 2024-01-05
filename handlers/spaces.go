package handlers

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type SpaceManager struct {
	Endpoint string
	Region   string
	Bucket   string
}

func NewSpaceManager(endpoint, region string) SpaceManager {
	return SpaceManager{
		Endpoint: endpoint,
		Region:   region,
		Bucket:   os.Getenv("DO_SPACE_NAME"),
	}
}

// Upload Any file to Space Manager
func (c SpaceManager) Upload(filename string) error {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &c.Endpoint,
		Region:   &c.Region,
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil
}
