package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"../util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	S3_REGION = "ap-southeast-1"
	S3_BUCKET = "brain-dump"
)

func main() {

	// Create a single AWS session (we can re use this if we're uploading many files)
	s := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	Value, err := s.Config.Credentials.Get()
	if err != nil {
		log.Println("Failed at creds")
		log.Fatal(err)
	}
	fmt.Println(Value)
	for _, file := range util.SearchFolder("../build") {
		// Upload
		err = AddFileToS3(s, "../build/"+file, file)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, filePath string, destFile string) error {

	// Open the file for use
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	contentType := "text/html"
	if strings.HasSuffix(filePath, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(filePath, ".ico") {
		contentType = "image/x-icon"
	}

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(destFile),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(contentType),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}
