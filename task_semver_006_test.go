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

func TestTaskSemver006(t *testing.T) {
	v, err := NewVersion("7")
	if err != nil || v.String() != "7.0.0" {
		t.Fatalf("v=%v err=%v", v, err)
	}
}
