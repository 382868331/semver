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

func TestTaskSemver002(t *testing.T) {
	s := "1.2.3+" + strings.Repeat("a", MaxVersionLen-6)
	if len(s) != MaxVersionLen {
		t.Fatal(len(s))
	}
	if _, err := StrictNewVersion(s); err != nil {
		t.Fatalf("err=%v", err)
	}
}
