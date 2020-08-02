package types

import "github.com/influenzanet/logging-service/pkg/api"

type LogQuery struct {
	EventType string
	Origin    string
	Start     int64
	End       int64
	EventName string
	UserID    string
}

func LogQueryFromAPI(q *api.LogQuery) LogQuery {
	if q == nil {
		return LogQuery{}
	}
	return LogQuery{
		EventType: q.EventName,
		Origin:    q.Origin,
		Start:     q.Start,
		End:       q.End,
		EventName: q.EventName,
		UserID:    q.UserId,
	}
}
