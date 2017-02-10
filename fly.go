package fly

import "github.com/frozzare/go-fly/adapter"

// Filesystem repretents a fly filesystem.
type Filesystem struct {
	adapter adapter.Adapter
}

// NewFly creates a new filesystem struct.
func NewFly(adapter adapter.Adapter) *Filesystem {
	return &Filesystem{adapter: adapter}
}

// CreateDir creates a new directory.
func (f *Filesystem) CreateDir(path string, args ...interface{}) error {
	return f.adapter.CreateDir(path)
}

// Copy will copy a file from source path to destionation path.
func (f *Filesystem) Copy(src string, dst string) error {
	return f.adapter.Copy(src, dst)
}

// Delete will delete a file from source path.
func (f *Filesystem) Delete(path string) error {
	return f.adapter.Delete(path)
}

// DeleteDir will delete a directory.
func (f *Filesystem) DeleteDir(path string) error {
	return f.adapter.DeleteDir(path)
}

// Has will check if a file exists.
func (f *Filesystem) Has(path string) (bool, error) {
	return f.adapter.Has(path)
}

// HasDir will check if a directory exists.
func (f *Filesystem) HasDir(path string) (bool, error) {
	return f.adapter.HasDir(path)
}

// MimeType will return the file mime type.
func (f *Filesystem) MimeType(path string) (string, error) {
	return f.adapter.MimeType(path)
}

// Read will read file content.
func (f *Filesystem) Read(path string) (string, error) {
	return f.adapter.Read(path)
}

// ReadAndDelete will read file content and then delete the file.
func (f *Filesystem) ReadAndDelete(path string) (string, error) {
	return f.adapter.ReadAndDelete(path)
}

// Rename will rename a file.
func (f *Filesystem) Rename(src string, dst string) error {
	return f.adapter.Rename(src, dst)
}

// Write will write content to a file.
func (f *Filesystem) Write(path, content string, args ...interface{}) error {
	return f.adapter.Write(path, content)
}
