package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func path() (path string, err error) {
	h, err := homedir.Dir()
	if err != nil {
		return
	}

	path = filepath.Join(h, ".config", "opr", "config.yml")
	return
}

func read() ([]byte, error) {
	path, err := path()
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path)
}

func write(data []byte) error {
	err := ensureConfig()
	if err != nil {
		return err
	}

	p, err := path()
	if err != nil {
		return err
	}

	return os.WriteFile(p, data, 0600)
}

func ensureConfig() error {
	p, err := path()
	if err != nil {
		return err
	}

	d := filepath.Dir(p)
	return ensureDir(d)
}

func ensureDir(path string) (err error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0700)
		}
		return
	}

	if !fi.IsDir() {
		err = fmt.Errorf("not a directory: %s", path)
	}
	return
}
