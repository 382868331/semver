package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver012(t *testing.T){v:=New(1,6,7,"","").IncMajor();if v.String()!="2.0.0"{t.Fatalf("got=%s",v.String())}}

func TestGoletaSemver012AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	v:=New(0,9,0,"beta","").IncMajor();if v.String()!="1.0.0"{t.Fatalf("got=%s",v.String())}
}
