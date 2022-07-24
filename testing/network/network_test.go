package network_test

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/jan-may/go-tests/network"
	"github.com/stretchr/testify/assert"
)

func TestReadHosts(t *testing.T) {

	filesystem := fstest.MapFS{
		"hosts": {
			Data: []byte("127.0.0.1 localhost\n"),
		},
	}

	content, err := network.ReadHosts(filesystem)
	assert.Equal(t, "127.0.0.1 localhost\n", content)
	assert.NoError(t, err)
}

func TestReadHostsFromFile(t *testing.T){
	tempDirectory, err := os.MkdirTemp("", "test-")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDirectory)

	hostfile := filepath.Join(tempDirectory, "hosts")
	err = os.WriteFile(hostfile, []byte("127.0.0.1 localhost\n"), 0644)
	assert.NoError(t, err)

	content, err := network.ReadHostsFromFile(hostfile)
	assert.Equal(t, "127.0.0.1 localhost\n", content)
	assert.NoError(t, err)
}