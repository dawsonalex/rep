package inmemory

import (
	"github.com/dawsonalex/rep"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestLastSession(t *testing.T) {
	store := NewStore()

	store.AddSession(&rep.Session{
		Name: "Session Name",
		Date: time.Now(),
		Sets: []rep.Set{
			{
				RepCount: 1,
				Weight:   1,
				Rpe:      1,
			},
		},
	})

	time.Sleep(time.Millisecond * 10)

	expectedLastSession := &rep.Session{
		RWMutex: sync.RWMutex{},
		Name:    "Session Name",
		Date:    time.Now(),
		Sets: []rep.Set{
			{
				RepCount: 2,
				Weight:   2,
				Rpe:      2,
			},
		},
	}
	store.AddSession(expectedLastSession)

	lastSession := store.LastSession("Session Name")
	if !reflect.DeepEqual(lastSession, expectedLastSession) {
		t.Errorf("expected %+v, got %+v", expectedLastSession, lastSession)
	}
}
