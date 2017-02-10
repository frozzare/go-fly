package flylocal

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

// Adapter represents a local adapter.
type Adapter struct {
	path string
}

// NewAdapter creates a new local adapter.
func NewAdapter(path string) *Adapter {
	return &Adapter{path}
}

// Copy will copy a file to new path locally.
func (a *Adapter) Copy(src string, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer srcFile.Close()

	sfi, err := srcFile.Stat()
	if err != nil {
		return err
	}

	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	dfi, err := destFile.Stat()
	if err != nil {
		return err
	}

	if !dfi.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", dst)
	}

	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}

	return nil
}

// CreateDir will create a directory.
func (a *Adapter) CreateDir(path string, args ...interface{}) error {
	perm := uint32(0777)
	if len(args) > 0 {
		perm = args[0].(uint32)
	}

	return os.MkdirAll(strings.TrimRight(path, "/")+"/", os.FileMode(perm))
}

// Delete will delete a file locally.
func (a *Adapter) Delete(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// DeleteDir will delete a directory.
func (a *Adapter) DeleteDir(path string) error {
	return a.Delete(strings.TrimRight(path, "/") + "/")
}

// Has will check whether a file exists.
func (a *Adapter) Has(path string) (bool, error) {
	_, err := os.Stat(path)
	return err == nil, nil
}

// HasDir will check whether a directory exists.
func (a *Adapter) HasDir(path string) (bool, error) {
	return a.Has(strings.TrimRight(path, "/") + "/")
}

// MimeType will return the file mime type.
func (a *Adapter) MimeType(path string) (string, error) {
	return strings.Split(mime.TypeByExtension(filepath.Ext(path)), ";")[0], nil
}

// Read will read a file locally.
func (a *Adapter) Read(path string) (string, error) {
	if _, err := a.Has(path); err != nil {
		return "", err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), err
}

// ReadAndDelete will read a file and delete it if any.
func (a *Adapter) ReadAndDelete(path string) (string, error) {
	content, err := a.Read(path)

	if err != nil {
		return "", err
	}

	if err := a.Delete(path); err == nil {
		return content, nil
	}

	return "", err
}

// Rename will rename a file to a new path locally.
func (a *Adapter) Rename(src, dst string) error {
	if err := a.Copy(src, dst); err != nil {
		return err
	}

	return a.Delete(src)
}

// Write will write a a new file locally.
func (a *Adapter) Write(path, content string, args ...interface{}) error {
	permission := uint32(0644)
	if len(args) > 0 {
		permission = args[0].(uint32)
	}

	if err := ioutil.WriteFile(path, []byte(content), os.FileMode(permission)); err != nil {
		return err
	}

	return nil
}
