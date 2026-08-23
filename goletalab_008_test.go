package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver008(t *testing.T){v:=New(1,2,3,"rc.1","");if v.String()!="1.2.3-rc.1"{t.Fatalf("got=%q",v.String())}}

func TestGoletaSemver008AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	v:=New(0,1,0,"alpha","");if v.String()!="0.1.0-alpha"{t.Fatalf("got=%q",v.String())}
}
