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

func TestTaskSemver015(t *testing.T) {
	a := MustParse("1.0.0-10")
	b := MustParse("1.0.0-2")
	if a.Compare(b) <= 0 {
		t.Fatalf("compare=%d", a.Compare(b))
	}
}

func TestTaskSemver015Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver015(t)
	}
}
