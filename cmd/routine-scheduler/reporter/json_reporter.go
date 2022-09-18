package reporter

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

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
	Mu                   sync.Mutex           `json:"-"`
}

func (j *JsonReporter) IncrementDatabaseReporter(err error) {
	j.Mu.Lock()
	defer j.Mu.Unlock()

	if err != nil {
		j.DatabaseCleanRoutine.RunErrors += 1
	} else {
		j.DatabaseCleanRoutine.RunNo += 1
	}
}

func (j *JsonReporter) IncrementLogsReporter(err error) {
	j.Mu.Lock()
	defer j.Mu.Unlock()

	if err != nil {
		j.LogsCleanRoutine.RunErrors += 1
	} else {
		j.LogsCleanRoutine.RunNo += 1
	}
}

func OpenReadAndUnmarshall(filePath string, jsonReporter *JsonReporter) (*JsonReporter, error) {

	jsonReporterRaw, err := os.Open("json_reporter.json")
	if err != nil {
		return nil, err
	}
	byteValue, _ := io.ReadAll(jsonReporterRaw)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(byteValue, jsonReporter)

	return jsonReporter, nil
}
