//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what hasFKColumn — FK 발견/미발견/테이블 미존재 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestHasFKColumn(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name:    "orders",
				Columns: map[string]ddl.Column{"id": {}, "user_id": {}},
				ForeignKeys: []ddl.ForeignKey{
					{Column: "user_id", RefTable: "users", RefColumn: "id"},
				},
			},
		},
	}

	tests := []struct {
		name string
		src  string
		col  string
		dst  string
		want bool
	}{
		{"FK found", "orders", "user_id", "users", true},
		{"wrong ref table", "orders", "user_id", "products", false},
		{"wrong column name", "orders", "id", "users", false},
		{"table not found", "products", "user_id", "users", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasFKColumn(fs, tt.src, tt.col, tt.dst)
			if got != tt.want {
				t.Errorf("hasFKColumn(%q, %q, %q) = %v, want %v", tt.src, tt.col, tt.dst, got, tt.want)
			}
		})
	}
}
