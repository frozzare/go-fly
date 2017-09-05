package flys3

import (
	"bytes"
	"errors"
	"mime"
	"path/filepath"
	"strings"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// Adapter represents a AWS S3 adapter.
type Adapter struct {
	bucket string
	s3     s3iface.S3API
}

// NewAdapter creates a new AWS S3 adapter.
func NewAdapter(client s3iface.S3API, bucket string) *Adapter {
	return &Adapter{bucket, client}
}

// Copy will copy a file to a new path on AWS S3.
func (a *Adapter) Copy(src, dst string) error {
	_, err := a.s3.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(a.bucket),
		Key:        aws.String(dst),
		CopySource: aws.String(a.bucket + "/" + src),
	})

	return err
}

// CreateDir will create a directory.
func (a *Adapter) CreateDir(path string, args ...interface{}) error {
	return a.Write(strings.TrimRight(path, "/")+"/", "")
}

// Delete will delete a file on AWS S3.
func (a *Adapter) Delete(path string) error {
	_, err := a.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(path),
	})

	return err
}

// DeleteDir will delete a directory.
func (a *Adapter) DeleteDir(path string) error {
	return a.Delete(strings.TrimRight(path, "/") + "/")
}

// Has will check whether a file exists.
func (a *Adapter) Has(path string) (bool, error) {
	_, err := a.s3.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// HasDir will check whether a directory exists.
func (a *Adapter) HasDir(path string) (bool, error) {
	return a.Has(strings.TrimRight(path, "/") + "/")
}

// MimeType will return the file mime type.
func (a *Adapter) MimeType(path string) (string, error) {
	res, err := a.s3.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		return "", err
	}

	return string(*res.ContentType), nil
}

// Read will read a file on AWS S3.
func (a *Adapter) Read(path string) (string, error) {
	res, err := a.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(buf), nil
}

// ReadAndDelete will read a file and delete it if any.
func (a *Adapter) ReadAndDelete(path string) (string, error) {
	content, err := a.Read(path)

	if err != nil {
		return "", err
	}

	return content, a.Delete(path)
}

// Rename will rename a file to a new path on AWS S3.
func (a *Adapter) Rename(src string, dst string) error {
	if err := a.Copy(src, dst); err != nil {
		return err
	}

	return a.Delete(src)
}

// Write will write a a new file AWS S3.
func (a *Adapter) Write(path, content string, args ...interface{}) error {
	res, err := a.s3.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(a.bucket),
		Key:           aws.String(path),
		Body:          bytes.NewReader([]byte(content)),
		ContentLength: aws.Int64(int64(len(content))),
		ContentType:   aws.String(mime.TypeByExtension(filepath.Ext(path))),
	})

	if err != nil {
		return err
	}

	if len(*res.ETag) == 0 {
		return errors.New("No ETag created for path")
	}

	return nil
}
