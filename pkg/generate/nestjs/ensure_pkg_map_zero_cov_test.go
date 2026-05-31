//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"
)

func TestEnsurePkgMap_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{}
	ensurePkgMap(pm, "billing")
	if pm["billing"] == nil {
		t.Error("expected submap created")
	}
	// idempotent
	pm["billing"]["x"] = true
	ensurePkgMap(pm, "billing")
	if !pm["billing"]["x"] {
		t.Error("ensurePkgMap should not reset existing submap")
	}
}
