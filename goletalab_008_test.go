package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver008(t *testing.T){v:=New(1,2,3,"rc.1","");if v.String()!="1.2.3-rc.1"{t.Fatalf("got=%q",v.String())}}
