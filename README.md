[![Go Report Card](https://goreportcard.com/badge/github.com/jamesfcarter/s3httpfilesystem)](https://goreportcard.com/report/github.com/jamesfcarter/s3httpfilesystem)
[![Documentation](https://godoc.org/github.com/jamesfcarter/s3httpfilesystem?status.svg)](http://godoc.org/github.com/jamesfcarter/s3httpfilesystem)

s3httpfilesystem implements the `http.FileSystem` interface allowing an S3
bucket to be easily served over http.

Example:
```
package main

import (
	"net/http"

	s3 "github.com/jamesfcarter/s3httpfilesystem"
)

func main() {
	endpoint := "s3-us-west-2.amazonaws.com"
	bucket := "example"
	region := "us-west-2"
	// Credentials come from the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
	// environment variables.
	s3FileSystem := s3.New(endpoint, region, bucket)
	http.ListenAndServe(":80", http.FileServer(s3FileSystem))
}
```
