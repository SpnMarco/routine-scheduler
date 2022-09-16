package models

// {
// 	"database_clean_routine": {
// 	  "run_no": 13,
// 	  "run_errors": 0
// 	},
// 	"logs_clean_routine": {
// 	  "run_no": 4,
// 	  "run_errors": 1
// 	}
// 	...
//   }

type databaseCleanRoutine struct {
	RunNo     int `json:"run_no"`
	RunErrors int `json:"run_errors"`
}

type logsCleanroutine struct {
	RunNo     int `json:"run_no"`
	RunErrors int `json:"run_errors"`
}

type JsonReporter struct {
	DatabaseCleanRoutine databaseCleanRoutine `json:"database_clean_routine"`
	LogsCleanRoutine     logsCleanroutine     `json:"logs_clean_routine"`
}
