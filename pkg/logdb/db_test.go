package logdb

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

var testDBService *LogDBService

const (
	testDBNamePrefix = "TEST_"
)

var (
	testInstanceID = strconv.FormatInt(time.Now().Unix(), 10)
)

func setupTestDBService() {
	dbConfig := GetDBConfig()
	testDBService = NewLogDBService(dbConfig)
}

func dropTestDB() {
	log.Println("Drop test database: studydb package")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := testDBService.DBClient.Database(testDBNamePrefix + testInstanceID + "_users").Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// Pre-Test Setup
func TestMain(m *testing.M) {
	setupTestDBService()
	result := m.Run()
	dropTestDB()
	os.Exit(result)
}
