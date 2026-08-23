package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver012(t *testing.T){v:=New(1,6,7,"","").IncMajor();if v.String()!="2.0.0"{t.Fatalf("got=%s",v.String())}}
