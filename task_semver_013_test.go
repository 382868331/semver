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

func TestTaskSemver013(t *testing.T) {
	a := MustParse("1.2.3+linux")
	b := MustParse("1.2.3+darwin")
	if a.Compare(b) != 0 {
		t.Fatalf("compare=%d", a.Compare(b))
	}
}

func TestTaskSemver013Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver013(t)
	}
}
