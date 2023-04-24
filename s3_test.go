package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

func main() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	log.Printf(" list items in bucket ")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	},
	)

	// Create S3 service client
	svc := s3.New(sess)

	firstImage := "/pubg-account-data/source-data/bunny_grass.gif"
	//secondImage := "/pubg-account-data/source-data/bunny_anim.gif"

	svc.DownloadFileFromS3("bunny_grass.gif", firstImage)
	if err != nil {
		return errors.Wrap(err, "paintjobs.ImageComposite - error downloading first image")
	}

	return
}
