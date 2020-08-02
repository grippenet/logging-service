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

var (
	testInstanceID = strconv.FormatInt(time.Now().Unix(), 10)
)

func setupTestDBService() {
	dbConfig := GetDBConfig()
	testDBService = NewLogDBService(dbConfig)
}

func dropTestDB() {
	log.Println("Drop test database: logdb package")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := testDBService.DBClient.Database(testDBService.DBNamePrefix + testInstanceID + "_users").Drop(ctx)
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
