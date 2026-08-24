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

func TestTaskSemver008(t *testing.T) {
	v := New(2, 3, 4, "rc.1", "linux")
	if v.Original() != "2.3.4-rc.1+linux" {
		t.Fatalf("original=%q", v.Original())
	}
}

func TestTaskSemver008Boundary(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTaskSemver008(t)
	}
}
