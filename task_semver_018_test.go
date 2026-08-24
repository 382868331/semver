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

func TestTaskSemver018(t *testing.T) {
	c, err := NewConstraint("<1.0.0 || >=2.0.0")
	if err != nil || !c.Check(MustParse("2.1.0")) {
		t.Fatalf("err=%v", err)
	}
}

func TestTaskSemver018Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver018(t)
	}
}
