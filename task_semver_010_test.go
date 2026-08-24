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

func TestTaskSemver010(t *testing.T) {
	v := MustParse("1.2.3").IncPatch()
	if v.String() != "1.2.4" {
		t.Fatalf("got=%s", v.String())
	}
}
