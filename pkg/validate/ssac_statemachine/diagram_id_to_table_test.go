//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what TestDiagramIDToTable — diagramIDToTable 분기별 (table, column) 도출 검증
package ssac_statemachine

import (
	"testing"
)

func TestDiagramIDToTable(t *testing.T) {
	cases := []struct {
		name       string
		id         string
		wantTable  string
		wantColumn string
	}{
		{name: "no underscore defaults status", id: "order", wantTable: "orders", wantColumn: "status"},
		{name: "leading underscore (idx<=0) defaults status", id: "_status", wantTable: "_statuses", wantColumn: "status"},
		{name: "trailing underscore (idx==len-1) defaults status", id: "order_", wantTable: "order_", wantColumn: "status"},
		{name: "entity and column split", id: "order_state", wantTable: "orders", wantColumn: "state"},
		{name: "uppercase entity lowered", id: "Order_phase", wantTable: "orders", wantColumn: "phase"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotTable, gotColumn := diagramIDToTable(tc.id)
			if gotTable != tc.wantTable {
				t.Errorf("diagramIDToTable(%q) table = %q, want %q", tc.id, gotTable, tc.wantTable)
			}
			if gotColumn != tc.wantColumn {
				t.Errorf("diagramIDToTable(%q) column = %q, want %q", tc.id, gotColumn, tc.wantColumn)
			}
		})
	}
}
