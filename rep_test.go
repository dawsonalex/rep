package rep

import (
	"reflect"
	"testing"
)

func TestTrackerSession(t *testing.T) {
	tracker := &Tracker{}

	err := tracker.BeginSession("session 1")
	if err != nil {
		t.Error(err)
	}

	tracker.LogSet(1, 2, 3, "")

	//  test that getting a finished set after closing a recv channel works.
	tracker.LogSet(1, 2, 3, "second set")

	session := tracker.EndSession()

	expectedSession := &Session{
		Name: "session 1",
		Sets: []Set{
			{
				RepCount: 1,
				Weight:   2,
				Rpe:      3,
				Note:     "",
			},
			{
				RepCount: 1,
				Weight:   2,
				Rpe:      3,
				Note:     "second set",
			},
		},
	}
	if !reflect.DeepEqual(session.Sets, expectedSession.Sets) || session.Name != expectedSession.Name {
		t.Errorf("session does not match. Expected: %+v but got %+v", expectedSession, session)
	}
}
