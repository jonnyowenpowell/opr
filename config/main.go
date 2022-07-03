package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name string
}

func Exists() (bool, error) {
	_, err := read()
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func get() (Config, error) {
	data, err := read()
	if os.IsNotExist(err) {
		return Config{}, nil
	} else if err != nil {
		return Config{}, err
	}

	var c Config
	err = yaml.Unmarshal(data, &c)
	return c, err
}

func set() error {
	data, err := yaml.Marshal(&Config{
		Name: "config_val",
	})
	if err != nil {
		return err
	}

	return write(data)
}
