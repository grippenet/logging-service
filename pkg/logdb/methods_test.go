package logdb

import (
	"testing"
	"time"

	"github.com/influenzanet/logging-service/pkg/types"
)

func TestSaveLogEvent(t *testing.T) {
	t.Run("Add log event ", func(t *testing.T) {
		_, err := testDBService.SaveLogEvent(testInstanceID, types.LogEvent{
			Time: time.Now().Unix(),
			Msg:  "test event",
		})
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}
	})
}

func TestFindLogEvents(t *testing.T) {
	testLogEvents := []types.LogEvent{
		{Time: 15, EventType: "Type1", UserID: "testuser1", EventName: "action1", Msg: "content", InstanceID: testInstanceID},
		{Time: 20, EventType: "Type1", UserID: "testuser1", EventName: "action1", Msg: "content", InstanceID: testInstanceID},
		{Time: 25, EventType: "Type1", UserID: "testuser1", EventName: "action1", Msg: "content", InstanceID: testInstanceID},
		{Time: 15, EventType: "Type2", UserID: "testuser1", EventName: "action1", Msg: "content", InstanceID: testInstanceID},
		{Time: 25, EventType: "Type2", UserID: "testuser1", EventName: "action1", Msg: "content", InstanceID: testInstanceID},
		{Time: 15, EventType: "Type1", UserID: "testuser2", EventName: "action1", Msg: "content", InstanceID: testInstanceID},
		{Time: 25, EventType: "Type1", UserID: "testuser2", EventName: "action1", Msg: "content", InstanceID: testInstanceID},
		{Time: 15, EventType: "Type1", UserID: "testuser2", EventName: "action2", Msg: "content", InstanceID: testInstanceID},
		{Time: 25, EventType: "Type1", UserID: "testuser2", EventName: "action2", Msg: "content", InstanceID: testInstanceID},
	}

	for _, e := range testLogEvents {
		_, err := testDBService.SaveLogEvent(testInstanceID, e)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	}

	t.Run("Finding all events ", func(t *testing.T) {
		err := testDBService.FindLogEvents(testInstanceID, types.LogQuery{
			EventType: "Type1",
		},
			func(instanceID string, event types.LogEvent, args ...interface{}) error {
				if event.EventType != "Type1" {
					t.Errorf("unexpected event returned: %v", event)
				}
				return nil
			},
		)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})
}
