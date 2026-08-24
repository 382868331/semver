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

func TestTaskSemver019(t *testing.T) {
	c := Collection{MustParse("2.0.0"), MustParse("1.0.0"), MustParse("1.5.0")}
	sort.Sort(c)
	if c[0].String() != "1.0.0" || c[2].String() != "2.0.0" {
		t.Fatalf("got=%v", c)
	}
}
