package routines

import (
	"encoding/json"
	"log"
	"time"

	"acme.spa/routine-scheduler/cmd/routine-scheduler/reporter"
	"acme.spa/routine-scheduler/cmd/routine-scheduler/utils"
	"acme.spa/routine-scheduler/pkg/logs"
)

func CleanLogs(jsonReporter *reporter.JsonReporter) {
	c := time.NewTicker(time.Duration(utils.EnvToUint32("CleanLogsEvery", 1)) * time.Minute).C

	for now := range c {
		log.Println("[ROUTINES] Starting CleanLogs Routine at: " + now.String())

		err := logs.Clean()

		if err != nil {
			log.Println("[ROUTINES] error in logs.Clean: " + err.Error())
		}

		jsonReporter.IncrementLogsReporter(err)

		jsonRaw, err := json.Marshal(jsonReporter)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanLogsRoutine on Marshalling: " + err.Error())
		}

		err = utils.WriteFile("json_reporter.json", jsonRaw, 0644)

		if err != nil {
			log.Println("[ROUTINES] Error in CleanLogsRoutine on WriteFile: " + err.Error())
		}

		log.Println("[ROUTINES] Finished CleanLogs Routine")
	}
}
