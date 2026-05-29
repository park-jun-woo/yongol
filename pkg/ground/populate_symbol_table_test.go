//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateSymbolTable — DDL 테이블명 → PascalCase singular Model 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateSymbolTable_Models verifies plural table names are converted to
// singular PascalCase model names in SymbolTable.model.
func TestPopulateSymbolTable_Models(t *testing.T) {
	fs := newMinimalFullstack(withDDLTables(
		ddl.Table{Name: "users"},
		ddl.Table{Name: "courses"},
		ddl.Table{Name: "order_items"},
	))
	g := newGround()

	populateSymbolTable(g, fs)

	models := g.Lookup["SymbolTable.model"]
	if !models["User"] {
		t.Errorf("SymbolTable.model missing User: %v", models)
	}
	if !models["Course"] {
		t.Errorf("SymbolTable.model missing Course: %v", models)
	}
	if !models["OrderItem"] {
		t.Errorf("SymbolTable.model missing OrderItem: %v", models)
	}
}
