//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindFloat — FamilyFloat → SQLAlchemy Float / Python float 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindFloat(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindFloat(ir.BindOpts{NotNull: true, DefaultLiteral: "0.0"})
		if b.Family != typemap.FamilyFloat || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "Float" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "float" {
			t.Errorf("APIType = %q, want float", b.APIType)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
		if b.DefaultLiteral != "0.0" {
			t.Errorf("DefaultLiteral = %q", b.DefaultLiteral)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindFloat(ir.BindOpts{NotNull: false})
		if b.APIType != "float | None" {
			t.Errorf("APIType = %q, want float | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
