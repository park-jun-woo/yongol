//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what tableHasColumn — 컬럼 존재/미존재/테이블 미존재 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestTableHasColumn(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{Name: "users", Columns: map[string]ddl.Column{"id": {}, "email": {}}},
		},
	}

	tests := []struct {
		name  string
		table string
		col   string
		want  bool
	}{
		{"column exists", "users", "email", true},
		{"column not exists", "users", "phone", false},
		{"table not found", "orders", "id", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tableHasColumn(fs, tt.table, tt.col)
			if got != tt.want {
				t.Errorf("tableHasColumn(%q, %q) = %v, want %v", tt.table, tt.col, got, tt.want)
			}
		})
	}
}
