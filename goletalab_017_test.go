package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver017(t *testing.T){a,_:=NewVersion("2.0.0");b,_:=NewVersion("1.9.9");if !a.GreaterThan(b)||b.GreaterThan(a){t.Fatal("ordering reversed")}}
