package server

import (
	"context"
	"testing"

	"github.com/influenzanet/go-utils/pkg/api_types"
	"github.com/influenzanet/go-utils/pkg/testutils"
	"github.com/influenzanet/logging-service/pkg/api"
	"google.golang.org/grpc"
)

func TestSaveLogEvent(t *testing.T) {
	s := loggingServer{
		logDBservice: testLogDBService,
	}

	t.Run("with nil ", func(t *testing.T) {
		_, err := s.SaveLogEvent(context.TODO(), nil)
		ok, msg := testutils.ShouldHaveGrpcErrorStatus(err, "missing arguments")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty argument", func(t *testing.T) {
		_, err := s.SaveLogEvent(context.TODO(), &api.NewLogEvent{})
		ok, msg := testutils.ShouldHaveGrpcErrorStatus(err, "missing arguments")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with normal argument", func(t *testing.T) {
		_, err := s.SaveLogEvent(context.TODO(), &api.NewLogEvent{
			EventType:  api.LogEventType_SECURITY,
			EventName:  "test",
			InstanceId: testInstanceID,
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

type LoggingServiceAPI_GetLogs struct {
	grpc.ServerStream
	Results []*api.LogEvent
}

func (_m *LoggingServiceAPI_GetLogs) Send(event *api.LogEvent) error {
	_m.Results = append(_m.Results, event)
	return nil
}

func TestGetLogs(t *testing.T) {
	s := loggingServer{
		logDBservice: testLogDBService,
	}

	t.Run("with nil ", func(t *testing.T) {
		err := s.GetLogs(nil, nil)
		ok, msg := testutils.ShouldHaveGrpcErrorStatus(err, "missing arguments")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty", func(t *testing.T) {
		mock := &LoggingServiceAPI_GetLogs{}
		err := s.GetLogs(&api.LogQuery{}, mock)
		ok, msg := testutils.ShouldHaveGrpcErrorStatus(err, "missing arguments")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with missing role", func(t *testing.T) {
		mock := &LoggingServiceAPI_GetLogs{}
		err := s.GetLogs(&api.LogQuery{
			Token: &api_types.TokenInfos{
				InstanceId:       testInstanceID,
				Id:               "testuser",
				AccountConfirmed: true,
				Payload: map[string]string{
					"roles": "PARTICIPANT",
				},
			},
			EventType: api.LogEventType_SECURITY,
		}, mock)
		ok, msg := testutils.ShouldHaveGrpcErrorStatus(err, "permission denied")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with valid arguments", func(t *testing.T) {
		mock := &LoggingServiceAPI_GetLogs{}
		err := s.GetLogs(&api.LogQuery{
			Token: &api_types.TokenInfos{
				InstanceId:       testInstanceID,
				Id:               "testuser",
				AccountConfirmed: true,
				Payload: map[string]string{
					"roles": "PARTICIPANT,ADMIN",
				},
			},
			EventType: api.LogEventType_SECURITY,
		}, mock)
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
			return
		}
		if len(mock.Results) < 1 {
			t.Errorf("unexpected number of log events: %d", len(mock.Results))
			return
		}
	})
}
