package storage

import "github.com/spf13/afero"

// Storage storage
type Storage struct {
	afero.Fs
}

// New init
func New() *Storage {
	var storage = &Storage{
		Fs: afero.NewOsFs(),
	}
	return storage
}
