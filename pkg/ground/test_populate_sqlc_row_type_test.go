//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateSQLc — RowType 을 SQLc.rowType Lookup 에 등록하는 회귀 테스트

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

// TestPopulateSQLc_RowTypeLookup verifies that each non-empty QuerySpec.RowType
// is exposed as "SQLc.rowType.<RowType>" in Ground.Lookup, and that ":exec" /
// ":execresult" queries (empty RowType) are not registered.
func TestPopulateSQLc_RowTypeLookup(t *testing.T) {
	fs := newMinimalFullstack()
	fs.SQLcQueries = []sqlcparser.QuerySpec{
		{Name: "UserFindByID", Cardinality: "one", RowType: "UserFindByIDRow"},
		{Name: "ListUsers", Cardinality: "many", RowType: "ListUsersRow"},
		{Name: "DeleteUser", Cardinality: "exec", RowType: ""},
		{Name: "UpdateUser", Cardinality: "execresult", RowType: ""},
	}

	g := newGround()
	populateSQLc(g, fs)

	for _, want := range []string{"UserFindByIDRow", "ListUsersRow"} {
		set, ok := g.Lookup["SQLc.rowType."+want]
		if !ok {
			t.Errorf("SQLc.rowType.%s missing; Lookup=%v", want, g.Lookup)
			continue
		}
		if !set[want] {
			t.Errorf("SQLc.rowType.%s set does not contain %q: %v", want, want, set)
		}
	}

	// Empty RowType must not produce an entry.
	for _, noRow := range []string{"DeleteUser", "UpdateUser"} {
		if _, ok := g.Lookup["SQLc.rowType."+noRow]; ok {
			t.Errorf("SQLc.rowType.%s unexpectedly registered for exec/execresult", noRow)
		}
		if _, ok := g.Lookup["SQLc.rowType."+noRow+"Row"]; ok {
			t.Errorf("SQLc.rowType.%sRow unexpectedly registered for exec/execresult", noRow)
		}
	}
}

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
