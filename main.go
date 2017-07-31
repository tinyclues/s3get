package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/tinyclues/s3get/download"
)

func main() {
	var filename string
	var bucket string
	var key string
	var awsRegion string

	awsConfig := aws.NewConfig()

	flag.StringVar(&filename, "filename", "", "Where s3Get will download the file")
	flag.StringVar(&bucket, "bucket", "", "The bucket name where it stored the s3 key")
	flag.StringVar(&key, "key", "", "The S3 key to download")
	flag.StringVar(&awsRegion, "region", "", "AWS region")
	flag.Parse()

	if awsRegion != "" {
		awsConfig.Region = aws.String(awsRegion)
	}

	// The session the S3 Downloader will use
	sess := session.Must(session.NewSession(awsConfig))
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	fileDescriptor, err := os.Create(filename)

	if err != nil {
		fmt.Errorf("failed to create file %q, %v", filename, err)
	}

	if err := download.S3Get(fileDescriptor, &bucket, &key, downloader); err != nil {
		fmt.Printf("Can't download the file from s3 %s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
