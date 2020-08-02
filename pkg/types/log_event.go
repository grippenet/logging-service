package types

import (
	"github.com/influenzanet/logging-service/pkg/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogEvent struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	InstanceID string             `bson:"instanceID,omitempty"`
	Time       int64              `bson:"time,omitempty"`
	TimeStr    string             `bson:"timeStr,omitempty"`
	EventType  string             `bson:"eventType,omitempty"`
	Origin     string             `bson:"origin,omitempty"`
	EventName  string             `bson:"eventName,omitempty"`
	UserID     string             `bson:"userID,omitempty"`
	Msg        string             `bson:"msg,omitempty"`
}

func (e LogEvent) ToAPI() *api.LogEvent {
	eventType := api.LogEventType_NONE
	switch e.EventType {
	case "SECURITY":
		eventType = api.LogEventType_SECURITY
	case "ERROR":
		eventType = api.LogEventType_ERROR
	case "LOG":
		eventType = api.LogEventType_LOG
	}
	return &api.LogEvent{
		Id:         e.ID.Hex(),
		Time:       e.Time,
		EventType:  eventType,
		Origin:     e.Origin,
		InstanceId: e.InstanceID,
		EventName:  e.EventName,
		UserId:     e.UserID,
		Msg:        e.Msg,
	}
}
