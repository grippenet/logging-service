package logdb

import (
	"testing"
)

func TestSaveLogEvent(t *testing.T) {
	t.Run("Add log event ", func(t *testing.T) {
		t.Error("test unimplemented")
	})
}

func TestFindLogEvents(t *testing.T) {
	/*
		pStates := []types.ParticipantState{
			{
				ParticipantID: "1",
				StudyStatus:   "active",
				Flags: map[string]string{
					"test1": "1",
				},
			},
			{
				ParticipantID: "2",
				StudyStatus:   "active",
			},
		}

		for _, ps := range pStates {
			_, err := testDBService.SaveParticipantState(testInstanceID, testStudyKey, ps)
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
				return
			}
		}

		t.Run("Finding inactive status ", func(t *testing.T) {
			err := testDBService.FindAndExecuteOnParticipantsStates(
				testInstanceID,
				testStudyKey,
				func(dbService *StudyDBService, p types.ParticipantState, instanceID, studyKey string) error {
					_, ok := p.Flags["test1"]
					if !ok {
						p.Flags = map[string]string{
							"test1": "1",
						}
					} else {
						p.Flags["test1"] = "newvalue"
					}
					_, err := dbService.SaveParticipantState(instanceID, studyKey, p)
					return err
				})
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
				return
			}

			p, err := testDBService.FindParticipantState(testInstanceID, testStudyKey, pStates[0].ParticipantID)
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
				return
			}
			testval, ok := p.Flags["test1"]
			if !ok || testval != "newvalue" {
				t.Errorf("unexpected flags for p1: %s", p.Flags)
			}

			p, err = testDBService.FindParticipantState(testInstanceID, testStudyKey, pStates[1].ParticipantID)
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
				return
			}
			testval, ok = p.Flags["test1"]
			if !ok || testval != "1" {
				t.Errorf("unexpected flags for p2: %s", p.Flags)
			}
		})*/
	t.Error("test unimplemented")
}
