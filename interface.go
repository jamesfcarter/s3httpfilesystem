package s3httpfilesystem

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

// s3Interface defines an interface that performs the actual AWS API
// calls we use.
type s3Interface interface {
	ListObjects(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error)
	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
}
