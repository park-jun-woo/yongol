//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiPackages_ZeroCov — addOpPackageRef/collectOpsPackages/buildSortedPackages 커버
package fastapi

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
		t.Fatalf("packages not sorted: %+v", got)
	}
	if len(got[1].Methods) != 2 || got[1].Methods[0] != "A" || got[1].Methods[1] != "B" {
		t.Errorf("methods not sorted: %+v", got[1].Methods)
	}
}
