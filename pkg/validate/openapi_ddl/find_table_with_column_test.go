//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what findTableWithColumn — 컬럼 발견/미발견 시 테이블명 반환 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFindTableWithColumn(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{Name: "users", Columns: map[string]ddl.Column{"id": {}, "email": {}}},
			{Name: "orders", Columns: map[string]ddl.Column{"id": {}, "total": {}}},
		},
	}

	tests := []struct {
		name string
		col  string
		want string
	}{
		{"found in first table", "email", "users"},
		{"found in second table", "total", "orders"},
		{"not found", "nonexistent", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTableWithColumn(fs, tt.col)
			if got != tt.want {
				t.Errorf("findTableWithColumn(%q) = %q, want %q", tt.col, got, tt.want)
			}
		})
	}
}
