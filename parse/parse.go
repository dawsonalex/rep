// Package parse provides functions for creating rep.Sessions and rep.Sets from string input.
package parse

import (
	"fmt"
	"github.com/dawsonalex/rep"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var setPattern = regexp.MustCompile(`(\d+)@(\d+)kg\s*\((\d+)\)`)

var sessionHeaderPattern = regexp.MustCompile(`(\w+\s*):`)

// PatternError describes a problem parsing a given input string.
// Calling Error() will return a message that describes what
// input was expected, and what was received.
type PatternError struct {
	wanted string
	got    string
}

func (e PatternError) Error() string {
	return fmt.Sprintf("invalid pattern, wanted: %s, got: %s", e.wanted, e.got)
}

// Set returns a rep.Set from a string, and an error if one occurred.
// set should be in the pattern `numberOfReps@weightkg (rpe)`.
func Set(set string) (rep.Set, error) {
	matches := setPattern.FindStringSubmatch(set)
	if len(matches) < 3 {
		return rep.Set{}, setPatternError(set)
	}
	reps, err := strconv.Atoi(matches[1])
	if err != nil {
		return rep.Set{}, err
	}

	weight, err := strconv.Atoi(matches[2])
	if err != nil {
		return rep.Set{}, err
	}

	rpe, err := strconv.Atoi(matches[3])
	if err != nil {
		return rep.Set{}, err
	}

	return rep.Set{
		RepCount: reps,
		Weight:   weight,
		Rpe:      rpe,
	}, nil
}

// Session returns a *rep.Session from a string. session can be either a session header (sessionName:) or
// a session header followed by one or more sets.
func Session(session string) (*rep.Session, error) {
	sessionLines := strings.Split(session, "\n")
	if !sessionHeaderPattern.MatchString(sessionLines[0]) {
		return nil, sessionPatternError(sessionLines[0])
	}

	sessionName := sessionLines[0][:len(sessionLines[0])-1]
	sets := make([]rep.Set, 0, len(sessionLines)-1)
	for _, setLine := range sessionLines[1:] {
		set, err := Set(setLine)
		if err != nil {
			return nil, err
		}

		sets = append(sets, set)
	}

	return &rep.Session{
		Name: sessionName,
		Date: time.Now(),
		Sets: sets,
	}, nil
}

func setPatternError(got string) error {
	return PatternError{
		wanted: "rep_count@weightkg (rpe)",
		got:    got,
	}
}

func sessionPatternError(got string) error {
	return PatternError{
		wanted: "session_name:",
		got:    got,
	}
}
