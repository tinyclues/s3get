package download

import (
	"os"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Downloader interface {
	Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (int64, error)
}

func S3Get(fileDescriptor *os.File, bucket *string, key *string, downloader Downloader) error {
	nbBytes, err := downloader.Download(fileDescriptor, &s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}

	fmt.Printf("file downloaded, %d bytes\n", nbBytes)
	return nil
}
