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
	Weight   int    // The weight in Kg for the set.
	Rpe      int    // Rate of perceived exertion for the set.
	Note     string // An arbitrary note for the set.
}

// Tracker manages a single session.
type Tracker struct {
	sync.Mutex
	session           *Session
	sessionInProgress bool

	subMux sync.Mutex      // mutex for reading from the subscriber slice
	subs   []chan *Session // A slice of subscriber channels
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

// EndSession marks the current session as ended in the tracker, and sends it to CompletedSessions.
func (t *Tracker) EndSession() *Session {
	t.Lock()
	defer t.Unlock()

	t.sessionInProgress = false

	return t.session
}

// LogSet logs a single set on the current session.
func (t *Tracker) LogSet(repCount, weight, rpe int, note string) {
	t.Lock()
	defer t.Unlock()

	t.session.logSet(Set{
		RepCount: repCount,
		Weight:   weight,
		Rpe:      rpe,
		Note:     note,
	})
}

func (s *Session) logSet(set Set) {
	s.Lock()
	defer s.Unlock()

	s.Sets = append(s.Sets, set)
}
