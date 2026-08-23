package semver
import("sort";"strings";"testing")
var _=sort.Sort;var _=strings.Contains
func TestGoletaSemver019(t *testing.T){c:=Collection{MustParse("2.0.0"),MustParse("1.0.0")};sort.Sort(c);if c[0].String()!="1.0.0"{t.Fatalf("first=%s",c[0])}}
