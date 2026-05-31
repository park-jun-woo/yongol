//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what XDM-28/Run/helper test — 초기 전이 vs DDL DEFAULT 일치 + 대상 매핑 검증
package ddl_statemachine

import (
	"testing"
)

func TestDiagramDDLTarget(t *testing.T) {
	cases := []struct {
		id         string
		wantTable  string
		wantColumn string
	}{
		{"Course", "courses", "status"},      // PascalEntity convention
		{"order", "orders", "status"},        // no underscore
		{"order_phase", "orders", "phase"},   // entity_column convention
		{"_leading", "_leadings", "status"},  // idx<=0 -> status fallback
		{"trailing_", "trailing_", "status"}, // trailing underscore -> fallback (plural of whole lowercased id)
	}
	for _, c := range cases {
		gotT, gotC := diagramDDLTarget(c.id)
		if gotT != c.wantTable || gotC != c.wantColumn {
			t.Errorf("diagramDDLTarget(%q) = (%q,%q), want (%q,%q)", c.id, gotT, gotC, c.wantTable, c.wantColumn)
		}
	}
}
