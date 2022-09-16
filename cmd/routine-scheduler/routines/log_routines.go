package routines

import (
	"encoding/json"
	"log"
	"time"

	"acme.spa/routine-scheduler/cmd/routine-scheduler/utils"
	"acme.spa/routine-scheduler/pkg/logs"
	"acme.spa/routine-scheduler/pkg/models"
)

func CleanLogs() {
	for {
		log.Println("[ROUTINES] Starting CleanLogs Routine...")

		byteValue, err := utils.OpenAndReadBytes("json_reporter.json")

		if err != nil {
			log.Println("[ROUTINES] Error in CleanLogs on OpenAndReadBytes: " + err.Error())
		}

		var jsonReporter models.JsonReporter

		err = json.Unmarshal(byteValue, &jsonReporter)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanLogs on Unmarshalling: " + err.Error())
		}

		err = logs.Clean()

		if err != nil {
			log.Println("[ROUTINES] Error In CleanLogs Routine")
			jsonReporter.LogsCleanRoutine.RunErrors += 1
		}
		jsonReporter.LogsCleanRoutine.RunNo += 1

		jsonRaw, err := json.Marshal(jsonReporter)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine on Marshalling: " + err.Error())
		}

		err = utils.WriteFile("json_reporter.json", jsonRaw, 0644)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine on WriteFile: " + err.Error())
		}

		log.Println("[ROUTINES] Finished CleanLogs Routine")
		time.Sleep(5 * time.Second)
	}
}
