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

func TestTaskSemver011(t *testing.T) {
	v := MustParse("1.2.9").IncMinor()
	if v.String() != "1.3.0" {
		t.Fatalf("got=%s", v.String())
	}
}

func TestTaskSemver011Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver011(t)
	}
}
