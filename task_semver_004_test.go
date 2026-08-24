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

func TestTaskSemver004(t *testing.T) {
	if _, err := StrictNewVersion("1.2.3+bad?"); err == nil {
		t.Fatal("invalid metadata accepted")
	}
}
