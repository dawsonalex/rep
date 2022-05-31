package rep

import (
	"errors"
	"sync"
	"time"
)

var ErrSessionInProgress = errors.New("already a session in progress")

// Session is a named group of sets, with specific date and time of completion.
type Session struct {
	sync.RWMutex
	Name string
	Date time.Time
	Sets []Set
}

// Set defines a single set of an exercise.
type Set struct {
	RepCount int
	Weight   int // The weight in Kg for the set.
	Rpe      int // Rate of perceived exertion for the set.
}

// Tracker manages a single session.
type Tracker struct {
	sync.Mutex
	session           *Session
	sessionInProgress bool
	CompletedSessions chan<- *Session //
}

// BeginSession starts a session with the given name
func (t *Tracker) BeginSession(name string) error {
	if t.sessionInProgress {
		return ErrSessionInProgress
	}
	t.Lock()
	defer t.Unlock()

	t.sessionInProgress = true

	t.session = &Session{
		Name: name,
		Date: time.Now(),
	}
	return nil
}

// EndSession ends the current session.
func (t *Tracker) EndSession() {
	t.Lock()
	defer t.Unlock()

	t.sessionInProgress = false
	// TODO: will this cause a deadlock if there isn't a receiver for the session?
	t.CompletedSessions <- t.session
}

// LogSet logs a single set on the current session.
func (t *Tracker) LogSet(repCount, weight, rpe int) {
	t.Lock()
	defer t.Unlock()

	t.session.logSet(Set{
		RepCount: repCount,
		Weight:   weight,
		Rpe:      rpe,
	})
}

func (s *Session) logSet(set Set) {
	s.Lock()
	defer s.Unlock()

	s.Sets = append(s.Sets, set)
}
