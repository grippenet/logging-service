package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/influenzanet/go-utils/pkg/api_types"
	"github.com/influenzanet/logging-service/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *loggingServer) Status(ctx context.Context, _ *empty.Empty) (*api_types.ServiceStatus, error) {
	return &api_types.ServiceStatus{
		Status:  api_types.ServiceStatus_NORMAL,
		Msg:     "service running",
		Version: apiVersion,
	}, nil
}

func (s *loggingServer) SaveLogEvent(ctx context.Context, req *api.NewLogEvent) (*api_types.ServiceStatus, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *loggingServer) GetLogs(req *api.LogQuery, stream api.LoggingServiceApi_GetLogsServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}
