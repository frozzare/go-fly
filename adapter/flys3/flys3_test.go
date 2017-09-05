package flys3

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/frozzare/go-assert"
)

var (
	ErrNoSuchBucket  = errors.New("NoSuchBucket: The specified bucket does not exist")
	ErrBucketExists  = errors.New("Bucket already exists")
	ErrBucketHasKeys = errors.New("Bucket has keys so cannot be deleted")
	ErrMisingKey     = errors.New("Missing key")
	client           = &MockS3{data: map[string]MockBucket{
		"/tmp": MockBucket{},
	}}
)

func TestDirectory(t *testing.T) {
	fs := NewAdapter(client, "/tmp")

	err := fs.CreateDir("test/folder")
	assert.Nil(t, err)

	has, err := fs.HasDir("test/folder")
	assert.True(t, has)
	assert.Nil(t, err)

	err = fs.DeleteDir("test/folder")
	assert.Nil(t, err)

	has, err = fs.HasDir("test/folder")
	assert.False(t, has)
	assert.Nil(t, err)
}

func TestFile(t *testing.T) {
	fs := NewAdapter(client, "/tmp")

	err := fs.Write("test/hello.txt", "Hello, world!")
	assert.Nil(t, err)

	err = fs.Copy("test/hello.txt", "test/hello-copy.txt")
	assert.Nil(t, err)

	has, err := fs.Has("test/hello.txt")
	assert.True(t, has)
	assert.Nil(t, err)

	content, err := fs.ReadAndDelete("test/hello-copy.txt")
	assert.Nil(t, err)
	assert.Equal(t, "Hello, world!", strings.TrimSpace(content))

	has, err = fs.Has("test/hello-copy.txt")
	assert.False(t, has)
	assert.NotNil(t, err)
}

func TestFileMimeType(t *testing.T) {
	fs := NewAdapter(client, "/tmp")

	err := fs.Write("test/hello.txt", "Hello, world!")
	assert.Nil(t, err)

	typ, err := fs.MimeType("test/hello.txt")
	assert.Equal(t, "text/plain", typ)
}

type MockBucket map[string][]byte
type MockS3 struct {
	s3iface.S3API
	sync.RWMutex
	data map[string]MockBucket
}

func (s *MockS3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	s.Lock()
	defer s.Unlock()
	content, _ := ioutil.ReadAll(input.Body)
	if bucket, ok := s.data[*input.Bucket]; ok {
		bucket[*input.Key] = content
	} else {
		return nil, ErrNoSuchBucket
	}
	return &s3.PutObjectOutput{
		ETag: input.Key,
	}, nil
}

func (s *MockS3) CopyObject(input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	s.Lock()
	defer s.Unlock()
	bucket := s.data[*input.Bucket]
	p := strings.Split(*input.CopySource, "/")
	src, ok := bucket[strings.Join(p[2:], "/")]
	if !ok {
		return nil, ErrMisingKey
	}

	bucket[*input.Key] = src

	return &s3.CopyObjectOutput{}, nil
}

func (s *MockS3) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	s.Lock()
	defer s.Unlock()
	bucket := s.data[*input.Bucket]
	delete(bucket, *input.Key)
	return &s3.DeleteObjectOutput{}, nil
}

func (s *MockS3) HeadObject(input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	s.Lock()
	defer s.Unlock()
	bucket := s.data[*input.Bucket]
	if _, ok := bucket[*input.Key]; !ok {
		return nil, ErrMisingKey
	}
	var c string
	if strings.HasSuffix(*input.Key, "txt") {
		c = "text/plain"
	}
	return &s3.HeadObjectOutput{
		ContentType: &c,
	}, nil
}

func (s *MockS3) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	s.RLock()
	defer s.RUnlock()
	bucket := s.data[*input.Bucket]
	if object, ok := bucket[*input.Key]; ok {
		body := ioutil.NopCloser(bytes.NewReader(object))
		output := s3.GetObjectOutput{
			Body: body,
		}
		return &output, nil
	}
	return nil, ErrMisingKey
}
