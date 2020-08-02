package server

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/influenzanet/logging-service/pkg/logdb"
)

var testLogDBService *logdb.LogDBService

const (
	testDBNamePrefix = "TEST_SERVICE_"
)

var (
	testInstanceID = strconv.FormatInt(time.Now().Unix(), 10)
)

// Pre-Test Setup
func TestMain(m *testing.M) {
	setupTestLogDBService()
	result := m.Run()
	dropTestDB()
	os.Exit(result)
}

func setupTestLogDBService() {
	testLogDBService = logdb.NewLogDBService(logdb.GetDBConfig())
}

func dropTestDB() {
	log.Println("Drop test database: service package")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := testLogDBService.DBClient.Database(testDBNamePrefix + testInstanceID + "_users").Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
