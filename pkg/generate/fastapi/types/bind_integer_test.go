//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindInteger — FamilyInteger → SQLAlchemy Integer / Python int 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindInteger(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindInteger(ir.BindOpts{NotNull: true, DefaultLiteral: "0"})
		if b.Family != typemap.FamilyInteger || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "Integer" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "int" {
			t.Errorf("APIType = %q, want int", b.APIType)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
		if b.DefaultLiteral != "0" {
			t.Errorf("DefaultLiteral = %q", b.DefaultLiteral)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindInteger(ir.BindOpts{NotNull: false})
		if b.APIType != "int | None" {
			t.Errorf("APIType = %q, want int | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
