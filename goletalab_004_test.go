package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver004(t *testing.T){v:=New(1,2,3,"rc.1","build");if v.Prerelease()!="rc.1"{t.Fatalf("pre=%q",v.Prerelease())}}
