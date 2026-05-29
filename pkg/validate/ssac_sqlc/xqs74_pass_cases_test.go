//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS74_SkipCases — dot target/no match/nil/no DDL/table not found 스킵 검증

package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS74_SkipCases(t *testing.T) {
	t.Run("DotTargetSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "CheckInstructor", FileName: "checkinstructor.ssac",
				Sequences: []ssacparser.Sequence{
					{Type: "get", Model: "RefreshToken.FindByHash", Result: &ssacparser.Result{Type: "RefreshToken", Var: "rt"}, Line: 3},
					{Type: "empty", Target: "rt.TokenHash", Line: 6},
				},
			}},
			DDLTables: []ddl.Table{{
				Name: "refresh_tokens", Columns: map[string]ddl.Column{"token_hash": {Name: "token_hash", RawType: "TEXT", NotNull: true}},
				ColumnOrder: []string{"token_hash"}, PrimaryKey: []string{"token_hash"},
			}},
		}
		diags := xqs74EmptyNonIntegerPK(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("NoTargetMatch", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "CheckSomething", FileName: "check.ssac",
				Sequences: []ssacparser.Sequence{{Type: "empty", Target: "unknown", Line: 5}},
			}},
			DDLTables: []ddl.Table{{
				Name: "refresh_tokens", Columns: map[string]ddl.Column{"token_hash": {Name: "token_hash", RawType: "TEXT", NotNull: true}},
				ColumnOrder: []string{"token_hash"}, PrimaryKey: []string{"token_hash"},
			}},
		}
		diags := xqs74EmptyNonIntegerPK(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("NilFullstack", func(t *testing.T) {
		diags := xqs74EmptyNonIntegerPK(nil)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("NoDDLTables", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetItem", FileName: "getitem.ssac",
				Sequences: []ssacparser.Sequence{
					{Type: "get", Model: "RefreshToken.FindByHash", Result: &ssacparser.Result{Type: "RefreshToken", Var: "rt"}, Line: 3},
					{Type: "empty", Target: "rt", Line: 6},
				},
			}},
		}
		diags := xqs74EmptyNonIntegerPK(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("TableNotFoundSkips", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetItem", FileName: "getitem.ssac",
				Sequences: []ssacparser.Sequence{
					{Type: "get", Model: "Widget.FindByID", Result: &ssacparser.Result{Type: "Widget", Var: "w"}, Line: 3},
					{Type: "empty", Target: "w", Line: 6},
				},
			}},
			DDLTables: []ddl.Table{{
				Name: "items", Columns: map[string]ddl.Column{"id": {Name: "id", RawType: "BIGINT", NotNull: true}},
				ColumnOrder: []string{"id"}, PrimaryKey: []string{"id"},
			}},
		}
		diags := xqs74EmptyNonIntegerPK(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})
}
