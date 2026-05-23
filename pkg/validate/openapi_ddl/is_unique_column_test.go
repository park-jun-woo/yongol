//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what isUniqueColumn — PK/UNIQUE 인덱스/비UNIQUE/테이블 미존재 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIsUniqueColumn(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name:       "users",
				PrimaryKey: []string{"id"},
				Indexes: []ddl.Index{
					{Name: "idx_email", Columns: []string{"email"}, IsUnique: true},
					{Name: "idx_name", Columns: []string{"name"}, IsUnique: false},
					{Name: "idx_multi", Columns: []string{"org_id", "name"}, IsUnique: true},
				},
			},
		},
	}

	tests := []struct {
		name  string
		table string
		col   string
		want  bool
	}{
		{"primary key is unique", "users", "id", true},
		{"single-column unique index", "users", "email", true},
		{"non-unique index", "users", "name", false},
		{"multi-column unique index (not single)", "users", "org_id", false},
		{"column not in any index", "users", "phone", false},
		{"table not found", "orders", "id", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isUniqueColumn(fs, tt.table, tt.col)
			if got != tt.want {
				t.Errorf("isUniqueColumn(%q, %q) = %v, want %v", tt.table, tt.col, got, tt.want)
			}
		})
	}
}
