//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what TestBuildDriverRules — catalog 전체 / fired-only fallback / 빈 fired nil 분기 검증
package sarif

import (
	"testing"

	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

// TestBuildDriverRules_FullCatalog covers the catalog branch: every
// catalogued rule is emitted regardless of fired set.
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

// TestBuildDriverRules_FiredFallback covers the no-catalog branch: only fired
// rules are emitted, sorted deterministically.
func TestBuildDriverRules_FiredFallback(t *testing.T) {
	fired := map[string]struct{}{"X-3": {}, "S-1": {}, "M-2": {}}
	got := buildDriverRules(nil, fired)
	if len(got) != 3 {
		t.Fatalf("rules: got %d, want 3", len(got))
	}
	want := []string{"M-2", "S-1", "X-3"} // sortStrings order
	for i, id := range want {
		if got[i].ID != id {
			t.Errorf("rules[%d].id: got %q, want %q", i, got[i].ID, id)
		}
	}
}

// TestBuildDriverRules_EmptyFired covers nil-catalog + empty-fired → nil.
func TestBuildDriverRules_EmptyFired(t *testing.T) {
	if got := buildDriverRules(nil, map[string]struct{}{}); got != nil {
		t.Errorf("empty fired: got %+v, want nil", got)
	}
}

// TestBuildDriverRules_EmptyCatalogFallsThrough covers a non-nil but empty
// catalog (Len()==0): the catalog branch is skipped and fired rules win.
func TestBuildDriverRules_EmptyCatalogFallsThrough(t *testing.T) {
	emptyCat := rulecatalog.NewCatalog(nil)
	got := buildDriverRules(emptyCat, map[string]struct{}{"S-9": {}})
	if len(got) != 1 || got[0].ID != "S-9" {
		t.Errorf("empty catalog should fall back to fired: got %+v", got)
	}
}
