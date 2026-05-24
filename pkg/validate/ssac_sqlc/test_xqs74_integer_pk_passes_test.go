//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS74_IntegerPKPasses — integer PK (BIGINT/SERIAL) 통과 케이스 검증

package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS74_IntegerPKPasses(t *testing.T) {
	t.Run("BIGINT", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetCourse", FileName: "getcourse.ssac",
				Sequences: []ssacparser.Sequence{
					{Type: "get", Model: "Course.FindByID", Result: &ssacparser.Result{Type: "Course", Var: "course"}, Line: 5},
					{Type: "empty", Target: "course", Message: "Course not found", Line: 8},
				},
			}},
			DDLTables: []ddl.Table{{
				Name: "courses", Columns: map[string]ddl.Column{"id": {Name: "id", RawType: "BIGINT", NotNull: true}, "title": {Name: "title", RawType: "TEXT"}},
				ColumnOrder: []string{"id", "title"}, PrimaryKey: []string{"id"},
			}},
		}
		diags := xqs74EmptyNonIntegerPK(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("SERIAL", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetItem", FileName: "getitem.ssac",
				Sequences: []ssacparser.Sequence{
					{Type: "get", Model: "Item.FindByID", Result: &ssacparser.Result{Type: "Item", Var: "item"}, Line: 3},
					{Type: "empty", Target: "item", Line: 5},
				},
			}},
			DDLTables: []ddl.Table{{
				Name: "items", Columns: map[string]ddl.Column{"id": {Name: "id", RawType: "SERIAL", NotNull: true}},
				ColumnOrder: []string{"id"}, PrimaryKey: []string{"id"},
			}},
		}
		diags := xqs74EmptyNonIntegerPK(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})
}
