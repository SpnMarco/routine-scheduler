package main

import (
	"log"
	"sync"

	"acme.spa/routine-scheduler/cmd/routine-scheduler/routines"
)

func main() {
	log.Println("Starting Routine Sheduler...")

	var wg sync.WaitGroup
	wg.Add(1)
	go routines.CleanDatabase()
	go routines.CleanLogs()

	wg.Wait()
	log.Println("Exiting Routine Sheduler...")

}
