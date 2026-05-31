//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"
)

func TestBuildSortedPackages_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{
		"zeta":  {"B": true, "A": true},
		"alpha": {"X": true},
	}
	got := buildSortedPackages(pm)
	if len(got) != 2 || got[0].Name != "alpha" || got[1].Name != "zeta" {
		t.Fatalf("unexpected order: %+v", got)
	}
	if len(got[1].Methods) != 2 || got[1].Methods[0] != "A" || got[1].Methods[1] != "B" {
		t.Errorf("methods not sorted: %+v", got[1].Methods)
	}
}
