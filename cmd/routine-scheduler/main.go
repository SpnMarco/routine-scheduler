package main

import (
	"encoding/json"
	"log"
	"sync"

	"acme.spa/routine-scheduler/cmd/routine-scheduler/reporter"
	"acme.spa/routine-scheduler/cmd/routine-scheduler/routines"
	"acme.spa/routine-scheduler/cmd/routine-scheduler/utils"
)

func main() {
	log.Println("Starting Routine Sheduler...")

	var jsonReporter reporter.JsonReporter

	bytes, _ := utils.OpenAndReadBytes("json_reporter.json")

	err := json.Unmarshal(bytes, &jsonReporter)

	if err != nil {
		log.Println("Cannot acquire JSON file for reporting routines, creating a new one")

		err := reporter.CreateDefault()

		if err != nil {
			log.Fatal("Cannot create reporting file, aborting. Err: " + err.Error())
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go routines.CleanDatabase(&jsonReporter)
	go routines.CleanLogs(&jsonReporter)

	wg.Wait()
}
