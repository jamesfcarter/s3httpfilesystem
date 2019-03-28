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
