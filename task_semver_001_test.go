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

func TestTaskSemver001(t *testing.T) {
	_, err := StrictNewVersion("")
	if !errors.Is(err, ErrEmptyString) {
		t.Fatalf("err=%v", err)
	}
}
