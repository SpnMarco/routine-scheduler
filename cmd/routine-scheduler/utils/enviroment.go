package utils

import (
	"log"
	"os"
	"strconv"
)

func EnvToUint32(env string, fallback uint32) uint32 {
	s := os.Getenv(env)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Unable to read env %s, using fallback", env)
		return fallback
	}
	return uint32(i)
}
