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

func TestTaskSemver009(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("MustParse did not panic")
		}
	}()
	_ = MustParse("not-semver")
}

func TestTaskSemver009Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver009(t)
	}
}
