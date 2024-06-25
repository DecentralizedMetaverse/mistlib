package content

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type LocalFileSystem struct {
	mu sync.RWMutex
}

var localFS = LocalFileSystem{}

func (fs *LocalFileSystem) Stat(path string) (bool, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("file not found: %s", path)
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (fs *LocalFileSystem) ReadFile(path string) ([]byte, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("file not found: %s", path)
	}
	return data, nil
}

func (fs *LocalFileSystem) WriteFile(path string, data []byte) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fs *LocalFileSystem) MkdirAll(path string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (fs *LocalFileSystem) Rename(oldPath, newPath string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	err := os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("file not found: %s", oldPath)
	}
	return nil
}

func (fs *LocalFileSystem) Remove(path string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("file not found: %s", path)
	}
	return nil
}

func (fs *LocalFileSystem) CopyFile(src, dst string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	data, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("file not found: %s", src)
	}
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
