package routines

import (
	"encoding/json"
	"log"
	"time"

	"acme.spa/routine-scheduler/cmd/routine-scheduler/utils"
	"acme.spa/routine-scheduler/pkg/database"
	"acme.spa/routine-scheduler/pkg/models"
)

func CleanDatabase() {
	for {
		log.Println("[ROUTINES] Starting CleanDatabase Routine")

		byteValue, err := utils.OpenAndReadBytes("json_reporter.go")
		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine on OpenAndReadBytes: " + err.Error())
		}

		var jsonReporter models.JsonReporter

		err = json.Unmarshal(byteValue, &jsonReporter)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine on Unmarshalling: " + err.Error())
		}

		err = database.Clean()

		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine")
			jsonReporter.DatabaseCleanRoutine.RunErrors += 1
		}

		jsonReporter.DatabaseCleanRoutine.RunNo += 1

		jsonRaw, err := json.Marshal(jsonReporter)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine on Marshalling: " + err.Error())
		}

		err = utils.WriteFile("json_reporter.json", jsonRaw, 0644)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine on WriteFile file: " + err.Error())
		}

		log.Println("[ROUTINES] Finished CleanDatabase Routine")
		time.Sleep(5 * time.Second)
	}
}
