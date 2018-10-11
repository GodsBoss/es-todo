package blob

import (
	"io/ioutil"
)

type FileBlobStore string

func NewFileBlobStore(filename string) FileBlobStore {
	return FileBlobStore(filename)
}

func (store FileBlobStore) Load() ([]byte, error) {
	return ioutil.ReadFile(string(store))
}

func (store FileBlobStore) Save(data []byte) error {
	return ioutil.WriteFile(string(store), data, 0644)
}
