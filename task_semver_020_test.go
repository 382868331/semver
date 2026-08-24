package semver

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"
	"testing"
)

var _ = json.Marshal
var _ = errors.Is
var _ sort.Interface
var _ = strings.Repeat

func TestTaskSemver020(t *testing.T) {
	parts := make([]string, MaxConstraintGroups)
	for i := range parts {
		parts[i] = ">=1.0.0"
	}
	if _, err := NewConstraint(strings.Join(parts, " || ")); err != nil {
		t.Fatalf("err=%v", err)
	}
}

func TestTaskSemver020Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver020(t)
	}
}
