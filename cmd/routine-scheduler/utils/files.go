package utils

import (
	"io"
	"io/fs"
	"log"
	"os"
	"sync"
)

var fileMutex sync.Mutex

func OpenAndReadBytes(filePath string) ([]byte, error) {
	log.Println("Starting OpenAndReadBytes")
	fileMutex.Lock()
	log.Println("LOCKING")
	defer fileMutex.Unlock()

	jsonReporterRaw, err := os.Open("json_reporter.json")
	if err != nil {
		return nil, err
	}
	byteValue, _ := io.ReadAll(jsonReporterRaw)
	if err != nil {
		return nil, err
	}
	log.Println("UNLOCKING")

	return byteValue, nil
}

func WriteFile(filePath string, data []byte, fileMode fs.FileMode) error {
	log.Println("Starting WriteFile")
	fileMutex.Lock()
	log.Println("LOCKING")
	defer fileMutex.Unlock()

	err := os.WriteFile(filePath, data, 0644)

	if err != nil {
		return err
	}
	log.Println("UNLOCKING")

	return nil
}
