//ff:func feature=ground type=test control=sequence dimension=1 topic=ddl
//ff:what populateDDL — DDL 컬럼 Go 타입이 Types["DDL.field.<Model>.<Field>"]에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDL_FieldTypes verifies that populateDDL registers DDL column
// Go types under the DDL.field.<ModelName>.<FieldName> key pattern.
func TestPopulateDDL_FieldTypes(t *testing.T) {
	tab := ddl.Table{
		Name: "workflows",
		Columns: map[string]ddl.Column{
			"id":     {Name: "id", RawType: "UUID", NotNull: true},
			"org_id": {Name: "org_id", RawType: "UUID", NotNull: true},
			"name":   {Name: "name", RawType: "TEXT", NotNull: true},
			"status": {Name: "status", RawType: "TEXT", NotNull: true},
			"count":  {Name: "count", RawType: "INTEGER", NotNull: true},
		},
		ColumnOrder: []string{"id", "org_id", "name", "status", "count"},
	}
	fs := newMinimalFullstack(withDDLTables(tab))
	g := newGround()

	populateDDL(g, fs)

	tests := []struct {
		key  string
		want string
	}{
		{"DDL.field.Workflow.ID", "pgtype.UUID"},
		{"DDL.field.Workflow.OrgID", "pgtype.UUID"},
		{"DDL.field.Workflow.Name", "string"},
		{"DDL.field.Workflow.Status", "string"},
		{"DDL.field.Workflow.Count", "int64"},
	}
	for _, tt := range tests {
		got := g.Types[tt.key]
		if got != tt.want {
			t.Errorf("Types[%q] = %q, want %q", tt.key, got, tt.want)
		}
	}
}
