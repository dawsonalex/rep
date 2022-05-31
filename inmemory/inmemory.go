// Package inmemory provides 'persistent' storage in memory.
package inmemory

import (
	rep "github.com/dawsonalex/golang-cli"
	"strings"
	"sync"
)

type Store struct {
	sync.RWMutex
	sessions map[string][]*rep.Session
}

func (s *Store) AddSession(session *rep.Session) {
	s.Lock()
	defer s.Unlock()

	s.sessions[nameToKey(session.Name())] = append(s.sessions[nameToKey(session.Name())], session)
}

func (s *Store) GetSessions(name string) []*rep.Session {
	s.RLock()
	defer s.RUnlock()

	return s.sessions[nameToKey(name)]
}

func (s *Store) LastSession(name string) *rep.Session {
	sessions := s.GetSessions(name)
	if len(sessions) > 0 {
		return sessions[0]
	}
	return nil
}

func nameToKey(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
