//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what TestLoadCached — Load 는 sync.Once 로 같은 Catalog 인스턴스 반환
package catalog

import "testing"

// TestLoadCached verifies Load returns the same Catalog instance across calls.
func TestLoadCached(t *testing.T) {
	c1, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	c2, err := Load()
	if err != nil {
		t.Fatalf("Load second call: %v", err)
	}
	if c1 != c2 {
		t.Errorf("Load should return cached Catalog, got distinct pointers")
	}
	if c1.Len() == 0 {
		t.Errorf("Load returned empty catalog")
	}
}
