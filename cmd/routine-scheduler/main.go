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
		log.Fatal("Cannot acquire JSON file for reporting routines, error: " + err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go routines.CleanDatabase(&jsonReporter)
	go routines.CleanLogs(&jsonReporter)

	wg.Wait()
}
