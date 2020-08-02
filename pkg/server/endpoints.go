package server

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/influenzanet/go-utils/pkg/api_types"
	"github.com/influenzanet/go-utils/pkg/token_checks"
	"github.com/influenzanet/logging-service/pkg/api"
	"github.com/influenzanet/logging-service/pkg/types"
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
	if req == nil || len(req.InstanceId) < 1 || req.EventType == api.LogEventType_NONE {
		return nil, status.Error(codes.InvalidArgument, "missing arguments")
	}

	event := types.LogEvent{
		Time:       time.Now().Unix(),
		TimeStr:    time.Now().String(),
		InstanceID: req.InstanceId,
		EventType:  req.EventType.Enum().String(),
		EventName:  req.EventName,
		Origin:     req.Origin,
		UserID:     req.UserId,
		Msg:        req.Msg,
	}
	_, err := s.logDBservice.SaveLogEvent(req.InstanceId, event)
	if err != nil {
		log.Printf("ERROR: unexpected error when saving log event: %v", err)
		return nil, status.Error(codes.Internal, "unexpected error")
	}
	return &api_types.ServiceStatus{
		Status:  api_types.ServiceStatus_NORMAL,
		Msg:     "event saved",
		Version: apiVersion,
	}, nil
}

func (s *loggingServer) GetLogs(req *api.LogQuery, stream api.LoggingServiceApi_GetLogsServer) error {
	if req == nil || token_checks.IsTokenEmpty(req.Token) || stream == nil {
		return status.Error(codes.InvalidArgument, "missing arguments")
	}
	if !token_checks.CheckRoleInToken(req.Token, "ADMIN") {
		log.Println("SECURITY: no admin rolen in token")
		return status.Error(codes.PermissionDenied, "permission denied")
	}
	query := types.LogQueryFromAPI(req)
	err := s.logDBservice.FindLogEvents(
		req.Token.InstanceId,
		query,
		func(instanceID string, event types.LogEvent, args ...interface{}) error {
			if len(args) != 1 {
				return errors.New("StreamUsers callback: unexpected number of args")
			}
			stream, ok := args[0].(api.LoggingServiceApi_GetLogsServer)
			if !ok {
				return errors.New(("StreamUsers callback: can't parse stream"))
			}

			if err := stream.Send(event.ToAPI()); err != nil {
				return err
			}
			return nil

		},
		stream,
	)
	if err != nil {
		log.Printf("ERROR: unexpected error when sending log event: %v", err)
		return status.Error(codes.Internal, "unexpected error")
	}
	return nil
}
