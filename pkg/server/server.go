package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/influenzanet/logging-service/pkg/api"
	"github.com/influenzanet/logging-service/pkg/logdb"
	"google.golang.org/grpc"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type loggingServer struct {
	logDBservice *logdb.LogDBService
}

// NewUserManagementServer creates a new service instance
func NewLoggingServer(
	logDBservice *logdb.LogDBService,
) api.LoggingServiceApiServer {
	return &loggingServer{
		logDBservice: logDBservice,
	}
}

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context,
	port string,
	logDBservice *logdb.LogDBService,
) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register service
	server := grpc.NewServer()
	api.RegisterLoggingServiceApiServer(server, NewLoggingServer(
		logDBservice,
	))

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	log.Println("wait connections on port " + port)
	return server.Serve(lis)
}
