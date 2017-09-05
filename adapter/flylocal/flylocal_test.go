package flylocal

import (
	"strings"
	"testing"

	"github.com/frozzare/go-assert"
)

func TestDirectory(t *testing.T) {
	fs := NewAdapter("/tmp/flylocal")

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
	fs := NewAdapter("/tmp/flylocal")

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
	assert.Nil(t, err)
}

func TestFileMimeType(t *testing.T) {
	fs := NewAdapter("/tmp/flylocal")

	err := fs.Write("test/hello.txt", "Hello, world!")
	assert.Nil(t, err)

	typ, err := fs.MimeType("test/hello.txt")
	assert.Equal(t, "text/plain", typ)
}
