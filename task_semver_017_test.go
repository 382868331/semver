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

func TestTaskSemver017(t *testing.T) {
	var v Version
	err := v.Scan([]byte("2.3.4"))
	if err != nil || v.String() != "2.3.4" {
		t.Fatalf("v=%s err=%v", v.String(), err)
	}
}
