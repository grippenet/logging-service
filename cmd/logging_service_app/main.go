package main

import (
	"context"
	"log"
	"os"

	"github.com/influenzanet/logging-service/pkg/logdb"
	"github.com/influenzanet/logging-service/pkg/server"
)

func main() {
	port := os.Getenv("LOGGING_SERVICE_LISTEN_PORT")
	logDBService := logdb.NewLogDBService(logdb.GetDBConfig())

	ctx := context.Background()

	if err := server.RunServer(
		ctx,
		port,
		logDBService,
	); err != nil {
		log.Fatal(err)
	}
}
