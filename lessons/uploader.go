package lessons

import (
	"cloud.google.com/go/storage"
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"
)

type GoogleUploader interface {
	UploadFile(file multipart.File, object string, key string) (*UploadFileResp, error)
}

type GoogleUploaderConfig struct {
	GoogleAppCred string
	BucketName    string
	ProjectId     string
	UploadPath    string
}

type UploadFileResp struct {
	FileUrl string `json:"file_url"`
}

func NewGoogleUploader(config GoogleUploaderConfig) GoogleUploader {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.GoogleAppCred)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader := &googleUploader{
		cl:         client,
		bucketName: config.BucketName,
		projectID:  config.ProjectId,
		uploadPath: config.UploadPath,
	}
	return uploader
}

type googleUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func (c *googleUploader) UploadFile(file multipart.File, object string, key string) (*UploadFileResp, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	fileUploadPath := c.uploadPath + key + "_" + object
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + key + "_" + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return nil, err
	}
	if err := wc.Close(); err != nil {
		return nil, err
	}
	return &UploadFileResp{FileUrl: "https://storage.cloud.google.com/" + c.bucketName + "/" + fileUploadPath}, nil
}
