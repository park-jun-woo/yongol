//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what findDDLTableWithColumn — 컬럼 미발견/발견 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFindDDLTableWithColumn(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name:    "users",
				Columns: map[string]ddl.Column{"id": {}, "email": {}},
			},
			{
				Name:    "orders",
				Columns: map[string]ddl.Column{"id": {}, "total": {}},
			},
		},
	}

	t.Run("column found returns table", func(t *testing.T) {
		tbl := findDDLTableWithColumn(fs, "email")
		if tbl == nil {
			t.Fatal("expected non-nil table")
		}
		if tbl.Name != "users" {
			t.Errorf("Name = %q, want users", tbl.Name)
		}
	})

	t.Run("column in second table", func(t *testing.T) {
		tbl := findDDLTableWithColumn(fs, "total")
		if tbl == nil {
			t.Fatal("expected non-nil table")
		}
		if tbl.Name != "orders" {
			t.Errorf("Name = %q, want orders", tbl.Name)
		}
	})

	t.Run("column not found returns nil", func(t *testing.T) {
		tbl := findDDLTableWithColumn(fs, "nonexistent")
		if tbl != nil {
			t.Errorf("expected nil, got %+v", tbl)
		}
	})

	t.Run("empty DDLTables returns nil", func(t *testing.T) {
		tbl := findDDLTableWithColumn(&yongol.Fullstack{}, "id")
		if tbl != nil {
			t.Errorf("expected nil, got %+v", tbl)
		}
	})
}
