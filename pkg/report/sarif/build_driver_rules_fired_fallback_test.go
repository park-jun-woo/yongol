//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what TestBuildDriverRules — catalog 전체 / fired-only fallback / 빈 fired nil 분기 검증
package sarif

import (
	"testing"
)

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
