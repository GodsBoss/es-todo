package blob

import (
	"io/ioutil"
)

type FileBlobStore string

func (store FileBlobStore) Load() ([]byte, error) {
	return ioutil.ReadFile(string(store))
}

func (store FileBlobStore) Save(data []byte) error {
	return ioutil.WriteFile(string(store), data, 0644)
}
