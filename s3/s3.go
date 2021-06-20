package s3

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// https://dev.classmethod.jp/articles/golang-write-s3/
func UploadObject() {
	// sessionの作成
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("REGION")),
		},
	))

	targetFilePath := "./sample.txt"
	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bucketName := os.Getenv("BUCKET_NAME")
	// Keyはパスを示す
	// {バケット名}/sample/sample.txt にしたいなら sample/sample.txt とする
	objectKey := os.Getenv("OBJECT_KEY")

	uploader := s3manager.NewUploader(sess)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println(output)
}

func DownloadObject() {
	// sessionの作成
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("REGION")),
		},
	))

	bucketName := os.Getenv("BUCKET_NAME")
	// Keyはパスを示す
	// {バケット名}/sample/sample.txt にしたいなら sample/sample.txt とする
	objectKey := os.Getenv("OBJECT_KEY")

	downloader := s3manager.NewDownloader(sess)
	buf := aws.NewWriteAtBuffer([]byte{})
	output, err := downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println(output)
}
