//ff:func feature=rule type=test control=iteration dimension=1
//ff:what populateSQLc — 빈 SQLcQueries 입력 시 SQLc.rowType.* 키 미등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestPopulateSQLc_RowTypeLookup_Empty ensures an empty SQLcQueries slice does
// not register any SQLc.rowType.* keys and does not panic.
func TestPopulateSQLc_RowTypeLookup_Empty(t *testing.T) {
	fs := newMinimalFullstack()
	g := newGround()
	populateSQLc(g, fs)

	for k := range g.Lookup {
		if len(k) >= len("SQLc.rowType.") && k[:len("SQLc.rowType.")] == "SQLc.rowType." {
			t.Errorf("unexpected SQLc.rowType entry on empty fs: %q", k)
		}
	}

	// populate_sqlc must be callable against a nil-free but empty fullstack.
	_ = yongol.Fullstack{} // compile-time sanity
}
