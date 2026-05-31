//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what TestBuildDriverRules — catalog 전체 / fired-only fallback / 빈 fired nil 분기 검증
package sarif

import (
	"testing"
)

func TestBuildDriverRules_FullCatalog(t *testing.T) {
	cat := testCatalog() // S-1, S-2, X-3
	got := buildDriverRules(cat, map[string]struct{}{"only-this": {}})
	if len(got) != cat.Len() {
		t.Fatalf("rules: got %d, want %d (full catalog)", len(got), cat.Len())
	}
	// Order follows catalog slice order.
	wantIDs := []string{"S-1", "S-2", "X-3"}
	for i, id := range wantIDs {
		if got[i].ID != id {
			t.Errorf("rules[%d].id: got %q, want %q", i, got[i].ID, id)
		}
	}
}
