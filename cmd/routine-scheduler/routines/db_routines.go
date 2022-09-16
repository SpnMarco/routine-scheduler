package routines

import (
	"log"
	"sync"
	"time"

	"acme.spa/routine-scheduler/pkg/database"
)

func CleanDatabase(wg *sync.WaitGroup) {
	for {
		log.Println("Starting CleanDatabase Routine")
		err := database.Clean()

		if err != nil {
			log.Println("Error in CleanDatabaseRoutine")
			//handle error
		}

		log.Println("Finished CleanDatabase Routine")
		time.Sleep(5 * time.Second)
	}
}
