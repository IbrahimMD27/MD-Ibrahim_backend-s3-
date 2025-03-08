package storage

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Session = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("Europe (Stockholm) eu-north-1"),
}))

func UploadFile(file multipart.File, fileName string) (string, error) {
	uploader := s3.New(s3Session)
	_, err := uploader.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("bucket-fileshare0"),
		Key:    aws.String(fileName),
		Body:   file,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://bucket-fileshare0.s3.amazonaws.com/%s", fileName), nil
}
