//ff:func feature=rule type=test control=iteration dimension=1
//ff:what populateSQLc — RowType 을 SQLc.rowType Lookup 에 등록하는 회귀 테스트

package ground

import (
	"testing"

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
		assertSQLcRowTypeLookupEntry(t, g, want)
	}

	// Empty RowType must not produce an entry.
	for _, noRow := range []string{"DeleteUser", "UpdateUser"} {
		assertSQLcRowTypeLookupAbsent(t, g, noRow)
	}
}
