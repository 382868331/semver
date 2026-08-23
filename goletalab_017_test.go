package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver017(t *testing.T){a,_:=NewVersion("2.0.0");b,_:=NewVersion("1.9.9");if !a.GreaterThan(b)||b.GreaterThan(a){t.Fatal("ordering reversed")}}

func TestGoletaSemver017AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	a,_:=NewVersion("1.0.0");b,_:=NewVersion("1.0.0-rc.1");if !a.GreaterThan(b){t.Fatal("release ordering")}
}
