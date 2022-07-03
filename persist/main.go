package persist

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/peterbourgon/diskv"
)

var store *diskv.Diskv

func ensureStore() error {
	if store == nil {
		h, err := homedir.Dir()
		if err != nil {
			return err
		}

		store = diskv.New(diskv.Options{
			BasePath:     filepath.Join(h, ".config", "opr", "templates"),
			CacheSizeMax: 1024 * 1024,
		})
	}
	return nil
}

func WriteTemplate(name string, value string) error {
	err := ensureStore()
	if err != nil {
		return fmt.Errorf("initialise store: %w", err)
	}

	return store.Write(fmt.Sprintf("%s.md", name), []byte(value))
}

func ReadTemplate(name string) (string, error) {
	err := ensureStore()
	if err != nil {
		return "", fmt.Errorf("initialise store: %w", err)
	}

	d, err := store.Read(fmt.Sprintf("%s.md", name))
	return string(d), err
}
