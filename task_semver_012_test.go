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

func TestTaskSemver012(t *testing.T) {
	v := MustParse("1.2.3")
	got, err := v.SetPrerelease("rc..1")
	if err == nil {
		t.Fatalf("got=%s", got.String())
	}
}

func TestTaskSemver012Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver012(t)
	}
}
