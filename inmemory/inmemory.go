// Package inmemory provides 'persistent' storage in memory.
package inmemory

import (
	"github.com/dawsonalex/rep"
	"strings"
	"sync"
)

type Store struct {
	sync.RWMutex
	sessions map[string][]*rep.Session
}

func NewStore() *Store {
	return &Store{
		sessions: make(map[string][]*rep.Session),
	}
}

func (s *Store) AddSession(session *rep.Session) {
	s.Lock()
	defer s.Unlock()

	s.sessions[nameToKey(session.Name)] = append(s.sessions[nameToKey(session.Name)], session)
}

func (s *Store) GetSessions(name string) ([]*rep.Session, bool) {
	s.RLock()
	defer s.RUnlock()

	sessions, found := s.sessions[nameToKey(name)]
	return sessions, found
}

func (s *Store) LastSession(name string) *rep.Session {
	sessions, ok := s.GetSessions(name)
	if !ok {
		return nil
	}
	return sessions[len(sessions)-1]
}

func nameToKey(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
