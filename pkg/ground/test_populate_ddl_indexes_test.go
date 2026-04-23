//ff:func feature=rule type=test control=sequence
//ff:what populateDDLIndexes — Index.Columns가 DDL.index.<table> 셋에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLIndexes_Columns verifies all columns of each index are
// flattened into a single per-table set.
func TestPopulateDDLIndexes_Columns(t *testing.T) {
	tab := ddl.Table{
		Name: "users",
		Indexes: []ddl.Index{
			{Name: "idx_email", Columns: []string{"email"}},
			{Name: "idx_composite", Columns: []string{"org_id", "role"}},
		},
	}
	g := newGround()
	populateDDLIndexes(g, tab)

	set := g.Lookup["DDL.index.users"]
	if !set["email"] || !set["org_id"] || !set["role"] {
		t.Errorf("DDL.index.users = %v, want email+org_id+role", set)
	}
	if len(set) != 3 {
		t.Errorf("len = %d, want 3", len(set))
	}
}
