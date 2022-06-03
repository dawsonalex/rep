package parse

import (
	"github.com/dawsonalex/rep"
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	type test struct {
		name       string
		text       string
		shouldPass bool
		reps       int
		weight     int
		rpe        int
	}

	tests := []test{
		{
			name:       "1-2-3",
			text:       "1@2kg (3)",
			shouldPass: true,
			reps:       1,
			weight:     2,
			rpe:        3,
		},
		{
			name:       "fail due to missing number",
			text:       "reps@2kg (3)",
			shouldPass: false,
			reps:       0,
			weight:     0,
			rpe:        0,
		},
		{
			name:       "fail due to bad syntax",
			text:       "1@@2kg (3)",
			shouldPass: false,
			reps:       0,
			weight:     0,
			rpe:        0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if set, err := Set(test.text); !test.shouldPass {
				if err == nil {
					tt.Error("expected test to fail but got nil error")
					return
				}
			} else {
				if err != nil {
					tt.Errorf("unexpected error: %v", err)
					return
				}
				if set.Rpe != test.rpe || set.Weight != test.weight || set.RepCount != test.reps {
					tt.Errorf("set stats don't match expected: rpe: %d wanted: %d, repCount: %d wanted: %d, weight: %d wanted: %d", set.Rpe, test.rpe, set.RepCount, test.reps, set.Weight, test.weight)
					return
				}
			}
		})
	}
}

func TestSession(t *testing.T) {
	type test struct {
		name       string
		text       string
		shouldPass bool
		sets       []rep.Set
	}

	tests := []test{
		{
			name:       "Valid session",
			text:       "First Session:\n1@2kg (3)\n1@2kg(3)",
			shouldPass: true,
			sets: []rep.Set{
				{
					RepCount: 1,
					Weight:   2,
					Rpe:      3,
				},
				{
					RepCount: 1,
					Weight:   2,
					Rpe:      3,
				},
			},
		},
		{
			name:       "fail due to bad set syntax",
			text:       "reps@2kg (3)",
			shouldPass: false,
			sets:       []rep.Set{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if session, err := Session(test.text); !test.shouldPass {
				if err == nil {
					tt.Error("expected test to fail but got nil error")
					return
				}
			} else {
				if err != nil {
					tt.Errorf("unexpected error: %v", err)
					return
				}
				if !reflect.DeepEqual(session.Sets, test.sets) {
					tt.Errorf("expected sets: %+v, got %+v", test.sets, session.Sets)
					return
				}
			}
		})
	}
}
