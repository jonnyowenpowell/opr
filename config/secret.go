package config

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

const (
	keyRingServiceName = "com.jonnyowenpowell.opr"
)

func ReadSecret(namespace, name string) (string, error) {
	return keyring.Get(fmt.Sprintf("%s.%s", keyRingServiceName, namespace), name)
}

func WriteSecret(namespace, name, value string) error {
	return keyring.Set(fmt.Sprintf("%s.%s", keyRingServiceName, namespace), name, value)
}
