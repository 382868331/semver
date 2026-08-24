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

func TestTaskSemver004(t *testing.T) {
	if _, err := StrictNewVersion("1.2.3+bad?"); err == nil {
		t.Fatal("invalid metadata accepted")
	}
}

func TestTaskSemver004Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver004(t)
	}
}
