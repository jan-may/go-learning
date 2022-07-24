package network

import (
	"io/fs"
	"os"
)

func ReadHosts(filesystem fs.FS) (string, error) {
	content, err := fs.ReadFile(filesystem, "hosts")
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ReadHostsFromFile(hostfile string) (string, error) {
	content, err := os.ReadFile(hostfile)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
