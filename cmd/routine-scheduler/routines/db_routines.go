package routines

import (
	"encoding/json"
	"log"
	"time"

	"acme.spa/routine-scheduler/cmd/routine-scheduler/reporter"
	"acme.spa/routine-scheduler/cmd/routine-scheduler/utils"
	"acme.spa/routine-scheduler/pkg/database"
)

func CleanDatabase(jsonReporter *reporter.JsonReporter) {
	c := time.NewTicker(time.Duration(utils.EnvToUint32("CleanDatabaseEvery", 1)) * time.Minute).C
	for now := range c {
		log.Println(now)
		log.Println("[ROUTINES] Starting CleanDatabase Routine at: " + now.String())

		err := database.Clean()

		if err != nil {
			log.Println("[ROUTINES] error in database.Clean: " + err.Error())
		}

		jsonReporter.IncrementDatabaseReporter(err)

		jsonRaw, err := json.Marshal(jsonReporter)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine on Marshalling: " + err.Error())
		}

		err = utils.WriteFile("json_reporter.json", jsonRaw, 0644)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanDatabaseRoutine on WriteFile file: " + err.Error())
		}

		log.Println("[ROUTINES] Finished CleanDatabase Routine")
	}
}
