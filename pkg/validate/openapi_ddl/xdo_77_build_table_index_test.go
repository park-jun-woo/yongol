//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what TestXdo77BuildTableIndex — empty/복수 테이블 인덱스 빌드 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo77BuildTableIndex(t *testing.T) {
	t.Run("empty DDLTables", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		idx := xdo77BuildTableIndex(fs)
		if len(idx) != 0 {
			t.Errorf("expected empty, got %v", idx)
		}
	})

	t.Run("builds index from tables", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id":    {RawType: "BIGINT"},
						"email": {RawType: "TEXT"},
					},
				},
			},
		}
		idx := xdo77BuildTableIndex(fs)
		if len(idx) != 1 {
			t.Fatalf("expected 1 table, got %d", len(idx))
		}
		if idx["users"]["id"] != "int64" {
			t.Errorf("users.id = %q, want int64", idx["users"]["id"])
		}
		if idx["users"]["email"] != "string" {
			t.Errorf("users.email = %q, want string", idx["users"]["email"])
		}
	})

	t.Run("UUID column mapped to uuid marker", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{
					Name: "orders",
					Columns: map[string]ddl.Column{
						"id":      {RawType: "UUID"},
						"user_id": {RawType: "UUID"},
						"name":    {RawType: "TEXT"},
					},
				},
			},
		}
		idx := xdo77BuildTableIndex(fs)
		if idx["orders"]["id"] != "uuid" {
			t.Errorf("orders.id = %q, want uuid", idx["orders"]["id"])
		}
		if idx["orders"]["user_id"] != "uuid" {
			t.Errorf("orders.user_id = %q, want uuid", idx["orders"]["user_id"])
		}
		if idx["orders"]["name"] != "string" {
			t.Errorf("orders.name = %q, want string", idx["orders"]["name"])
		}
	})
}
