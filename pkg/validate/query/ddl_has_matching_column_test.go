//ff:func feature=validate type=test control=iteration dimension=2 topic=query-structural
//ff:what ddlHasMatchingColumn — nil filter/empty tables/no match/match 검증

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestDdlHasMatchingColumn(t *testing.T) {
	t.Run("nil filter returns false", func(t *testing.T) {
		tables := []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"id": {Name: "id", RawType: "BIGINT"},
				},
			},
		}
		if ddlHasMatchingColumn(tables, nil) {
			t.Fatal("expected false for nil filter")
		}
	})

	t.Run("empty tables returns false", func(t *testing.T) {
		filter := func(col ddl.Column) bool { return true }
		if ddlHasMatchingColumn(nil, filter) {
			t.Fatal("expected false for nil tables")
		}
	})

	t.Run("no matching column returns false", func(t *testing.T) {
		tables := []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"id":   {Name: "id", RawType: "BIGINT"},
					"name": {Name: "name", RawType: "TEXT"},
				},
			},
		}
		filter := func(col ddl.Column) bool {
			return headTokenEquals(col.RawType, "UUID")
		}
		if ddlHasMatchingColumn(tables, filter) {
			t.Fatal("expected false when no column matches")
		}
	})

	t.Run("matching column in first table returns true", func(t *testing.T) {
		tables := []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"id": {Name: "id", RawType: "UUID"},
				},
			},
		}
		filter := func(col ddl.Column) bool {
			return headTokenEquals(col.RawType, "UUID")
		}
		if !ddlHasMatchingColumn(tables, filter) {
			t.Fatal("expected true when column matches")
		}
	})

	t.Run("matching column in second table returns true", func(t *testing.T) {
		tables := []ddl.Table{
			{
				Name: "orders",
				Columns: map[string]ddl.Column{
					"id": {Name: "id", RawType: "BIGINT"},
				},
			},
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"id": {Name: "id", RawType: "UUID"},
				},
			},
		}
		filter := func(col ddl.Column) bool {
			return headTokenEquals(col.RawType, "UUID")
		}
		if !ddlHasMatchingColumn(tables, filter) {
			t.Fatal("expected true when column matches in second table")
		}
	})
}
