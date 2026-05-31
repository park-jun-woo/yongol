//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestToIRBinding_Mapping(t *testing.T) {
	gb := GoTypeBinding{
		SqlcGoType:     "pgtype.Int4",
		ApiField:       "*int64",
		ConvertExpr:    "conv",
		InsertExpr:     "ins",
		ResponseExpr:   "resp",
		NilCheckExpr:   "nil",
		Imports:        []string{"pkg/a"},
		Supported:      true,
		DefaultLiteral: "0",
	}
	opts := ir.BindOpts{NotNull: false}
	got := toIRBinding(typemap.FamilyInteger, opts, gb)
	if got.DBType != "pgtype.Int4" || got.APIType != "*int64" {
		t.Errorf("type mapping mismatch: %+v", got)
	}
	if got.ToDBExpr != "ins" || got.ToAPIExpr != "conv" || got.ToResponseExpr != "resp" {
		t.Errorf("expr mapping mismatch: %+v", got)
	}
	if got.NotNull != false || got.Supported != true || got.DefaultLiteral != "0" {
		t.Errorf("scalar mapping mismatch: %+v", got)
	}
	if len(got.APIImports) != 1 || got.APIImports[0] != "pkg/a" {
		t.Errorf("imports mapping mismatch: %+v", got.APIImports)
	}
}
