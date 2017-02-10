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

func (a *Adapter) appendPath(path string) string {
	return filepath.Join(a.path, path)
}

// Copy will copy a file to new path locally.
func (a *Adapter) Copy(src string, dst string) error {
	srcFile, err := os.Open(a.appendPath(src))
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

	destFile, err := os.Create(a.appendPath(dst))
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

	return os.MkdirAll(strings.TrimRight(a.appendPath(path), "/")+"/", os.FileMode(perm))
}

// Delete will delete a file locally.
func (a *Adapter) Delete(path string) error {
	if err := os.Remove(a.appendPath(path)); err != nil {
		return err
	}

	return nil
}

// DeleteDir will delete a directory.
func (a *Adapter) DeleteDir(path string) error {
	return a.Delete(strings.TrimRight(a.appendPath(path), "/") + "/")
}

// Has will check whether a file exists.
func (a *Adapter) Has(path string) (bool, error) {
	_, err := os.Stat(a.appendPath(path))
	return err == nil, nil
}

// HasDir will check whether a directory exists.
func (a *Adapter) HasDir(path string) (bool, error) {
	return a.Has(strings.TrimRight(a.appendPath(path), "/") + "/")
}

// MimeType will return the file mime type.
func (a *Adapter) MimeType(path string) (string, error) {
	return strings.Split(mime.TypeByExtension(filepath.Ext(a.appendPath(path))), ";")[0], nil
}

// Read will read a file locally.
func (a *Adapter) Read(path string) (string, error) {
	path = a.appendPath(path)

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
	path = a.appendPath(path)
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
	perm := uint32(0644)
	if len(args) > 0 {
		perm = args[0].(uint32)
	}

	a.CreateDir(filepath.Dir(path))

	if err := ioutil.WriteFile(a.appendPath(path), []byte(content), os.FileMode(perm)); err != nil {
		return err
	}

	return nil
}
