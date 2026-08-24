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

func TestTaskSemver014(t *testing.T) {
	release := MustParse("1.0.0")
	pre := MustParse("1.0.0-rc.1")
	if release.Compare(pre) <= 0 {
		t.Fatalf("compare=%d", release.Compare(pre))
	}
}
