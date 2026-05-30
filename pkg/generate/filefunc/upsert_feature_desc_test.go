//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestUpsertFeatureDesc — 빈 이름/기존 비어있음/기존 채워짐 분기 검증

package filefunc

import "testing"

func TestUpsertFeatureDesc(t *testing.T) {
	dst := map[string]string{
		"hasdesc": "keep",
		"empty":   "",
	}

	upsertFeatureDesc(dst, "", "ignored")     // empty name skipped
	upsertFeatureDesc(dst, "hasdesc", "new")  // existing non-empty preserved
	upsertFeatureDesc(dst, "empty", "filled") // existing empty overwritten
	upsertFeatureDesc(dst, "fresh", "desc")   // new key inserted

	if _, ok := dst[""]; ok {
		t.Errorf("empty name should not be inserted")
	}
	if got := dst["hasdesc"]; got != "keep" {
		t.Errorf("hasdesc: expected preserved, got %q", got)
	}
	if got := dst["empty"]; got != "filled" {
		t.Errorf("empty: expected overwritten, got %q", got)
	}
	if got := dst["fresh"]; got != "desc" {
		t.Errorf("fresh: expected inserted, got %q", got)
	}
}
