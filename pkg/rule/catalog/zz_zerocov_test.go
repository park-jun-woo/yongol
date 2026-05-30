//ff:func feature=rule type=test topic=catalog
//ff:what zz_zerocov_test — catalog.Index / catalog.MustLoad 0% 커버리지 단위 테스트
package catalog

import "testing"

func TestIndex_ZeroCov(t *testing.T) {
	c := NewCatalog([]RuleMeta{{ID: "X-1"}, {ID: "X-2"}})
	if got := c.Index("X-1"); got != 0 {
		t.Errorf("Index(X-1)=%d want 0", got)
	}
	if got := c.Index("X-2"); got != 1 {
		t.Errorf("Index(X-2)=%d want 1", got)
	}
	if got := c.Index("missing"); got != -1 {
		t.Errorf("Index(missing)=%d want -1", got)
	}
	// nil receiver → -1.
	var nilC *Catalog
	if got := nilC.Index("X-1"); got != -1 {
		t.Errorf("nil.Index=%d want -1", got)
	}
}

func TestMustLoad_ZeroCov(t *testing.T) {
	// MustLoad on the embedded catalog must succeed (no log.Fatal path).
	c := MustLoad()
	if c == nil {
		t.Fatal("MustLoad returned nil")
	}
}
