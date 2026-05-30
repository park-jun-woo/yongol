//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindString — FamilyString → SQLAlchemy String / Python str 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindString(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindString(ir.BindOpts{NotNull: true, DefaultLiteral: "''"})
		if b.Family != typemap.FamilyString || !b.Supported {
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
		if b.DefaultLiteral != "''" {
			t.Errorf("DefaultLiteral = %q", b.DefaultLiteral)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindString(ir.BindOpts{NotNull: false})
		if b.APIType != "str | None" {
			t.Errorf("APIType = %q, want str | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
