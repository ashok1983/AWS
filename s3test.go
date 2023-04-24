package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	log.Printf(" list items in bucket ")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	},
	)
	firstImage := "/pubg-account-data/source-data/bunny_grass.gif"

	downloader := s3manager.NewDownloader(sess)
	// Create a file to write the S3 Object contents to.
	f, err := os.Create("test-file")
	if err != nil {
		fmt.Println(err)
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String(firstImage),
	})
	if err != nil {
		fmt.Println(err, n)
	}

	return
}
