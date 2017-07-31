package download

import (
	"testing"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/docker/docker/pkg/testutil/assert"
	"io"
	"github.com/aws/aws-sdk-go/service/s3"
	"errors"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FakeDownloader struct {
	fail bool
}

func (f *FakeDownloader) Download (w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (int64, error) {
	if f.fail {
		return 0, errors.New("Access Denied, check your credentials")
	} else {
		return 1, nil
	}
}

func TestS3Get(t *testing.T) {
	fd := &os.File{}
	downloader := &FakeDownloader{fail: false}

	err := S3Get(fd, aws.String("toto"), aws.String("titi"), downloader)
	assert.NilError(t, err)
}

func TestS3GetFail(t *testing.T) {
	fd := &os.File{}
	downloader := &FakeDownloader{fail: true}

	err := S3Get(fd, aws.String("toto"), aws.String("titi"), downloader)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "failed to download file, Access Denied, check your credentials")
}
