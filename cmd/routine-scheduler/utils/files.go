package utils

import (
	"io"
	"io/fs"
	"os"
	"sync"
)

var fileMutex sync.Mutex

func OpenAndReadBytes(filePath string) ([]byte, error) {

	jsonReporterRaw, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	byteValue, _ := io.ReadAll(jsonReporterRaw)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}

func WriteFile(filePath string, data []byte, fileMode fs.FileMode) error {

	fileMutex.Lock()
	defer fileMutex.Unlock()

	err := os.WriteFile(filePath, data, 0644)

	if err != nil {
		return err
	}
	return nil
}
