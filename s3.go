// snippet-comment:[These are tags for the AWS doc team's sample catalog. Do not remove.]
// snippet-sourceauthor:[Doug-AWS]
// snippet-sourcedescription:[Lists all of your S3 buckets.]
// snippet-keyword:[Amazon Simple Storage Service]
// snippet-keyword:[Amazon S3]
// snippet-keyword:[ListBuckets function]
// snippet-keyword:[Go]
// snippet-sourcesyntax:[go]
// snippet-service:[s3]
// snippet-keyword:[Code Sample]
// snippet-sourcetype:[full-example]
// snippet-sourcedate:[2018-03-16]
/*
   Copyright 2010-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at
    http://aws.amazon.com/apache2.0/
   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/
package main

import (
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
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

    // result, err := svc.ListBuckets(&s3.ListBucketsInput{})
    // if err != nil {
    //     log.Println("Failed to list buckets", err)
    //     return
    // }

    // log.Println("Buckets:")

    // for _, bucket := range result.Buckets {
    //     log.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
    // }

    resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String("pubg-account-data"), Prefix: aws.String("source-data")})
    if err != nil {
        log.Printf("Unable to list items in bucket  %v", err)
    }
    for _, item := range resp.Contents {
        log.Printf("Name:         %s", *item.Key)
        log.Printf("Last modified:%s", *item.LastModified)
        log.Printf("Size:         %d", *item.Size)
        log.Printf("Storage class: %s", *item.StorageClass)
        log.Printf("\n")

        if *item.Key == "source-data/" {
            continue
        }
        srcKey := "/" + "pubg-account-data" + "/" + *item.Key
        destKey := "destination-data/" + *item.Key
        fmt.Printf("Item  copying from bucket %q to bucket %q \n", srcKey, destKey)

        _, err = svc.CopyObject(&s3.CopyObjectInput{
            Bucket:     aws.String("pubg-account-data"),
            CopySource: aws.String(srcKey),
            Key:        aws.String(fmt.Sprintf("%v/%v", "destination-data", *item.Key))})
        if err != nil {
            fmt.Printf("Unable to copy item  from bucket %q to bucket %q, %v", srcKey, *item.Key, err)
            return
        }

        // Wait to see if the item got copied
        err = svc.WaitUntilObjectExists(&s3.HeadObjectInput{Bucket: aws.String("pubg-account-data"), Key: item.Key})
        if err != nil {
            fmt.Printf("Error occurred while waiting for item to be copied to bucket %q, %v", destKey, err)
            return
        }
        fmt.Printf("Item successfully copied from bucket %q to bucket %q\n", srcKey, destKey)
        // Delete item from source
        _, err = svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String("pubg-account-data"), Key: item.Key})
        if err != nil {
            fmt.Printf("Unable to delete object %q from bucket %q, %v", *item.Key, "pubg-account-data", err)
            return
        }

        err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
            Bucket: aws.String("pubg-account-data"),
            Key:    item.Key,
        })
        if err != nil {
            fmt.Printf("Error occurred while waiting for object %q to be deleted, %v", *item.Key, err)
            return
        }

    }
    return
}
