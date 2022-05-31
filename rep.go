package rep

import (
	"sync"
	"time"
)

// session is a named group of sets, with specific date and time of completion.
type session struct {
	sync.RWMutex
	name string
	date time.Time
	sets []set
}

// set defines a single set of an exercise.
type set struct {
	repCount int
	weight   int // The weight in Kg for the set.
	rpe      int // Rate of perceived exertion for the set.
}

// Tracker manages a single session.
type Tracker struct {
	sync.Mutex
	session *session
}

// BeginSession starts a session with the given name
func (t *Tracker) BeginSession(name string) {
	t.Lock()
	defer t.Unlock()

	t.session = &session{
		name: name,
		date: time.Now(),
	}
}

// EndSession ends the current session.
func (t *Tracker) EndSession() {

}

// LogSet logs a single set on the current session.
func (t *Tracker) LogSet(repCount, weight, rpe int) {
	t.Lock()
	defer t.Unlock()

	t.session.logSet(set{
		repCount: repCount,
		weight:   weight,
		rpe:      rpe,
	})
}

func (s *session) logSet(set set) {
	s.Lock()
	defer s.Unlock()

	s.sets = append(s.sets, set)
}
