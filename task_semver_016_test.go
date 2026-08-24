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

func TestTaskSemver016(t *testing.T) {
	var v Version
	if err := json.Unmarshal([]byte(`"1.2.3+linux"`), &v); err != nil || v.String() != "1.2.3+linux" {
		t.Fatalf("v=%s err=%v", v.String(), err)
	}
}

func TestTaskSemver016Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver016(t)
	}
}
