//ff:func feature=ground type=test control=iteration dimension=1 topic=ddl
//ff:what populateDDL apifield — DDL.apifield.<M>.<f> 등록 + Struct.<M>.<f> 키 토큰 동일성 고정 (BUG-099)
package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestPopulateDDL_Apifield(t *testing.T) {
	tab := ddl.Table{
		Name: "workflows",
		Columns: map[string]ddl.Column{
			"id":     {Name: "id", RawType: "UUID", NotNull: true},
			"org_id": {Name: "org_id", RawType: "UUID", NotNull: true},
			"name":   {Name: "name", RawType: "TEXT", NotNull: true},
			"count":  {Name: "count", RawType: "INTEGER", NotNull: true},
		},
		ColumnOrder: []string{"id", "org_id", "name", "count"},
	}
	fs := newMinimalFullstack(withDDLTables(tab))
	g := newGround()

	populateDDL(g, fs)

	tests := []struct {
		key  string
		want string
	}{
		// UUID columns: ApiField = openapi_types.UUID (not string).
		{"DDL.apifield.Workflow.ID", "openapi_types.UUID"},
		{"DDL.apifield.Workflow.OrgID", "openapi_types.UUID"},
		// Non-UUID columns keep their api-surface type.
		{"DDL.apifield.Workflow.Name", "string"},
		{"DDL.apifield.Workflow.Count", "int64"},
	}
	for _, tt := range tests {
		got := g.Types[tt.key]
		if got != tt.want {
			t.Errorf("Types[%q] = %q, want %q", tt.key, got, tt.want)
		}
	}
}
