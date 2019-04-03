package s3httpfilesystem

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	as3 "github.com/aws/aws-sdk-go/service/s3"
)

// File implements the http.File interface for a S3 object or simulated
// directory.
type File struct {
	fs      *Filesystem
	path    string
	content *bytes.Reader
}

func (f *File) download() error {
	if f.content != nil {
		return nil
	}
	log.Printf("S3 downloading %s\n", f.path)
	input := as3.GetObjectInput{
		Bucket: &f.fs.bucket,
		Key:    &f.path,
	}
	output, err := f.fs.s3.GetObject(&input)
	if err != nil {
		return err
	}
	defer output.Body.Close()
	data, err := ioutil.ReadAll(output.Body)
	if err != nil {
		return err
	}
	f.content = bytes.NewReader(data)
	return nil
}

// Read reads from a downloaded S3 object as if it were a file
func (f *File) Read(p []byte) (int, error) {
	err := f.download()
	if err != nil {
		return 0, err
	}
	return f.content.Read(p)
}

// Close closes and removes a downloaded S3 object
func (f *File) Close() error {
	f.content = nil
	return nil
}

// Seek sets the offset for the next Read of a downloaded S3 object
func (f *File) Seek(offset int64, whence int) (int64, error) {
	err := f.download()
	if err != nil {
		return 0, err
	}
	return f.content.Seek(offset, whence)
}

// Readdir returns the contents of a simulated directory of S3 objects
func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	resp, err := f.fs.s3.ListObjects((&as3.ListObjectsInput{}).
		SetBucket(f.fs.bucket).
		SetPrefix(f.path).
		SetDelimiter("/"))
	if err != nil {
		return nil, err
	}
	return listObjectsToFileInfo(resp), nil
}

// Stat returns the FileInfo structure describing file.
func (f *File) Stat() (os.FileInfo, error) {
	delimiter := "/"
	input := as3.ListObjectsInput{
		Bucket:    &f.fs.bucket,
		Prefix:    &f.path,
		Delimiter: &delimiter,
		MaxKeys:   aws.Int64(1),
	}
	resp, err := f.fs.s3.ListObjects(&input)
	if err != nil {
		return nil, err
	}
	if len(resp.Contents) == 0 {
		return nil, os.ErrNotExist
	}
	return listObjectsToFileInfo(resp)[0], nil
}

func listObjectsToFileInfo(r *as3.ListObjectsOutput) []os.FileInfo {
	objects := make([]os.FileInfo, 0,
		len(r.Contents)+len(r.CommonPrefixes))
	for _, p := range r.CommonPrefixes {
		objects = append(objects, dirInfo{p})
	}
	for _, o := range r.Contents {
		objects = append(objects, fileInfo{o})
	}
	return objects
}
