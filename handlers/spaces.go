package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// SpaceManager is digital oceans s3 bucket
type SpaceManager struct {
	Endpoint string
	Region   string
	Bucket   string
}

// NewSpaceManager creates an instance of SpaceManager
func NewSpaceManager(endpoint, region string) SpaceManager {
	return SpaceManager{
		Endpoint: endpoint,
		Region:   region,
		Bucket:   os.Getenv("DO_SPACE_NAME"),
	}
}

// Upload Any file to Space Manager
func (s SpaceManager) Upload(filename string) error {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &s.Endpoint,
		Region:   &s.Region,
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil
}

// DownloadFile can get Any file from Space Manager
func (s SpaceManager) DownloadFile(filename, destPath string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &s.Endpoint,
		Region:   &s.Region,
	}))
	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file %q, %v", destPath, err)
	}
	defer file.Close()

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	return nil
}

// GetFiles gets all files from a path
func (s SpaceManager) GetFiles(path string) ([]string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &s.Endpoint,
		Region:   &s.Region,
	}))
	svc := s3.New(sess)

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(path),
	}

	result, err := svc.ListObjectsV2(input)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects, %v", err)
	}

	var files []string
	for _, item := range result.Contents {
		files = append(files, *item.Key)
	}

	return files, nil
}

// DownloadAllFiles will use GetFiles method to get the files for a path in our space so we can loop through that path and use DownloadFile method
func (s SpaceManager) DownloadAllFiles(path, localDir string) error {
	// Step 1: List all files
	files, err := s.GetFiles(path)
	if err != nil {
		return fmt.Errorf("error listing files: %v", err)
	}

	// Step 2: Download each file
	for _, file := range files {
		destPath := filepath.Join(localDir, filepath.Base(file))
		err := s.DownloadFile(file, destPath)
		if err != nil {
			fmt.Printf("error downloading file %s: %v\n", file, err)
			// Optionally, continue downloading other files even if one fails
		} else {
			fmt.Printf("downloaded %s to %s\n", file, destPath)
		}
	}

	return nil
}

// GetFileContents will open the file contents in a buffer
func (s SpaceManager) GetFileContents(filename string) ([]byte, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &s.Endpoint,
		Region:   &s.Region,
	}))
	downloader := s3manager.NewDownloader(sess)

	buffer := aws.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file, %v", err)
	}

	return buffer.Bytes(), nil
}
