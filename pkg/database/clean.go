package database

import (
	"fmt"
	"time"
)

// Clean - Clean old entries from database
// DO NOT EDIT
func Clean() error {
	fmt.Println("[DATABASE] Cleaning old entries from database...")
	time.Sleep(10 * time.Second)
	fmt.Println("[DATABASE] Cleanup completed!")

	return nil
}
