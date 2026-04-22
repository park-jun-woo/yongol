//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateDDLCheck — CHECK IN enum 값이 Lookup+Schemas에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLCheck_EnumValues verifies CHECK IN enums are registered to
// both Lookup (set for O(1) containment) and Schemas (ordered slice).
func TestPopulateDDLCheck_EnumValues(t *testing.T) {
	tab := ddl.Table{
		Name: "orders",
		CheckEnums: map[string][]string{
			"status": {"pending", "paid", "cancelled"},
		},
	}
	g := newGround()

	populateDDLCheck(g, tab)

	set := g.Lookup["DDL.check.orders.status"]
	if !set["pending"] || !set["paid"] || !set["cancelled"] {
		t.Fatalf("DDL.check.orders.status = %v", set)
	}

	vals := g.Schemas["DDL.check.orders.status"]
	if len(vals) != 3 {
		t.Fatalf("Schemas[DDL.check.orders.status] = %v, want 3 values", vals)
	}
}

// TestPopulateDDLCheck_Empty verifies no panic/extra keys when no CHECK enum.
func TestPopulateDDLCheck_Empty(t *testing.T) {
	tab := ddl.Table{Name: "empty"}
	g := newGround()
	populateDDLCheck(g, tab)

	if _, ok := g.Lookup["DDL.check.empty.status"]; ok {
		t.Errorf("unexpected DDL.check key for empty CheckEnums")
	}
}
