package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver018(t *testing.T){a,_:=NewVersion("1.2.3");b,_:=NewVersion("1.2.3+other");if !a.GreaterThanEqual(b){t.Fatal("equal version rejected")}}
