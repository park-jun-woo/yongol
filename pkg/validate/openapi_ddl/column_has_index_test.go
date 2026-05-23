//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what columnHasIndex — PK/인덱스 선행 컬럼/미발견 테이블/미인덱스 컬럼 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestColumnHasIndex(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name:       "users",
				PrimaryKey: []string{"id"},
				Indexes: []ddl.Index{
					{Name: "idx_email", Columns: []string{"email"}},
					{Name: "idx_multi", Columns: []string{"org_id", "name"}},
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
		{"primary key column", "users", "id", true},
		{"index leading column", "users", "email", true},
		{"multi-column index leading", "users", "org_id", true},
		{"multi-column index non-leading", "users", "name", false},
		{"unindexed column", "users", "phone", false},
		{"table not found", "orders", "id", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := columnHasIndex(fs, tt.table, tt.col)
			if got != tt.want {
				t.Errorf("columnHasIndex(%q, %q) = %v, want %v", tt.table, tt.col, got, tt.want)
			}
		})
	}
}
