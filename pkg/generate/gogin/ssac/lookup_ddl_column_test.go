//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what lookupDDLColumn 단위 테스트 ((table, column) → *ddl.Column)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestLookupDDLColumn(t *testing.T) {
	tables := []ddl.Table{
		{
			Name: "workflows",
			Columns: map[string]ddl.Column{
				"id":         {Name: "id", RawType: "UUID", NotNull: true},
				"created_at": {Name: "created_at", RawType: "TIMESTAMPTZ", NotNull: true},
			},
		},
	}
	cases := []struct {
		name    string
		model   string
		column  string
		wantNil bool
		wantRaw string
	}{
		{"existing id", "Workflow", "Id", false, "UUID"},
		{"PascalCase to snake", "Workflow", "CreatedAt", false, "TIMESTAMPTZ"},
		{"missing column", "Workflow", "Missing", true, ""},
		{"missing table", "Nonexistent", "Id", true, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := lookupDDLColumn(tables, tc.model, tc.column)
			if tc.wantNil {
				if got != nil {
					t.Errorf("expected nil, got %+v", got)
				}
				return
			}
			if got == nil {
				t.Fatalf("expected column, got nil")
			}
			if got.RawType != tc.wantRaw {
				t.Errorf("RawType = %q, want %q", got.RawType, tc.wantRaw)
			}
		})
	}
}
