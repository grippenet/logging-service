package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/coneno/logger"
	"github.com/influenzanet/logging-service/pkg/logdb"
	"github.com/influenzanet/logging-service/pkg/types"
)

func collectLogEntries(instanceID string, logEvent types.LogEvent, args ...interface{}) error {
	collector := args[0].(*[]types.LogEvent)
	*collector = append(*collector, logEvent)
	return nil
}

type Interval struct {
	Start int64
	End   int64
}

func main() {
	// read config
	instances := strings.Split(os.Getenv("INSTANCES"), ",")
	outputRootDir := os.Getenv("OUTPUT_DIR")
	nMonth, err := strconv.Atoi(os.Getenv("FOR_LAST_N_MONTH"))
	if err != nil {
		logger.Error.Fatal("FOR_LAST_N_MONTH must be an integer: " + err.Error())
	}

	logDBService := logdb.NewLogDBService(logdb.GetDBConfig())

	intervals := []Interval{}
	now := time.Now()

	for i := 0; i < nMonth; i++ {
		t := now.AddDate(0, -i, 0)
		firstday := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
		lastday := firstday.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
		intervals = append(intervals, Interval{
			Start: firstday.Unix(),
			End:   lastday.Unix(),
		})
	}

	for _, instance := range instances {
		for _, i := range intervals {

			logEntries := []types.LogEvent{}

			query := types.LogQuery{
				Start: i.Start,
				End:   i.End,
			}

			err := logDBService.FindLogEvents(instance, query, collectLogEntries, &logEntries)
			if err != nil {
				logger.Error.Fatalf("%s: %v", instance, err)
			}

			// write files
			outputDir := outputRootDir + "/" + instance
			err = os.MkdirAll(outputDir, os.ModePerm)
			if err != nil {
				logger.Error.Fatalf("%s: %v", instance, err)
			}
			file, _ := json.MarshalIndent(logEntries, "", "  ")

			startDate := time.Unix(i.Start, 0)
			fileName := outputDir + "/" + fmt.Sprintf("%d-%02d.json", startDate.Year(), startDate.Month())
			_ = ioutil.WriteFile(fileName, file, 0644)
			logger.Info.Printf("file created: %s with %d entries", fileName, len(logEntries))
		}
	}

}
