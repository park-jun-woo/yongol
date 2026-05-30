//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindEnum — FamilyEnum → SQLAlchemy String / Python str 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindEnum(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindEnum(ir.BindOpts{NotNull: true, DefaultLiteral: "'active'"})
		if b.Family != typemap.FamilyEnum || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "String" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "str" {
			t.Errorf("APIType = %q, want str", b.APIType)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
		if b.DefaultLiteral != "'active'" {
			t.Errorf("DefaultLiteral = %q", b.DefaultLiteral)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindEnum(ir.BindOpts{NotNull: false})
		if b.APIType != "str | None" {
			t.Errorf("APIType = %q, want str | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
