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

func TestTaskSemver005(t *testing.T) {
	if _, err := StrictNewVersion("01.2.3"); err == nil {
		t.Fatal("leading zero accepted")
	}
}

func TestTaskSemver005Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver005(t)
	}
}
