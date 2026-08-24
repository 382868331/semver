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

func TestTaskSemver003(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic=%v", r)
		}
	}()
	if _, err := StrictNewVersion("1.2"); err == nil {
		t.Fatal("missing patch accepted")
	}
}
