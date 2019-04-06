package s3httpfilesystem

import (
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	as3 "github.com/aws/aws-sdk-go/service/s3"
)

// Filesystem implements the http.FileSystem interface for an S3 bucket.
type Filesystem struct {
	s3     s3Interface
	bucket string
}

// New builds and returns a new Filesystem accessing an S3.
// Credentials come from the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
// environment variables.
func New(endpoint, region, bucket string) *Filesystem {
	creds := credentials.NewEnvCredentials()
	config := aws.NewConfig().
		WithCredentials(creds).
		WithEndpoint(endpoint).
		WithRegion(region)
	sess := session.Must(session.NewSession(config))
	return &Filesystem{
		s3:     as3.New(sess),
		bucket: bucket,
	}
}

// NewWithS3 takes an s3Interface and is used for testing
func NewWithS3(bucket string, mock s3Interface) *Filesystem {
	return &Filesystem{
		s3:     mock,
		bucket: bucket,
	}
}

// Open returns an http.File interface for an object or simulated directory
// within the S3 bucket.
func (f *Filesystem) Open(path string) (http.File, error) {
	return &File{fs: f, path: prepPath(path)}, nil
}

func prepPath(path string) string {
	path = strings.TrimPrefix(path, "/")
	return path
}
