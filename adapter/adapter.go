package adapter

// Adapter represents a Fly adapter.
type Adapter interface {
	CreateDir(string, ...interface{}) error
	Copy(string, string) error
	Delete(string) error
	DeleteDir(string) error
	Has(string) (bool, error)
	HasDir(string) (bool, error)
	MimeType(string) (string, error)
	Read(string) (string, error)
	ReadAndDelete(string) (string, error)
	Rename(string, string) error
	Write(string, string, ...interface{}) error
}
