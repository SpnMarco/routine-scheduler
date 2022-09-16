package logs

import (
	"errors"
	"math/rand"
	"time"
)

// Clean - Clean log file from old entries
// DO NOT EDIT
func Clean() error {
	time.Sleep(2*time.Second)
	rand.Seed(time.Now().Unix())

	random := rand.Intn(2)
	if random >= 1 {
		return errors.New("Unable to open log file")
	}

	return nil
}
