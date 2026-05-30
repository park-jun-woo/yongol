//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindNumeric — FamilyNumeric → SQLAlchemy Numeric / Python Decimal 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindNumeric(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindNumeric(ir.BindOpts{NotNull: true})
		if b.Family != typemap.FamilyNumeric || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "Numeric" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "Decimal" {
			t.Errorf("APIType = %q, want Decimal", b.APIType)
		}
		if b.ToDBExpr != "Decimal({var})" {
			t.Errorf("ToDBExpr = %q", b.ToDBExpr)
		}
		if b.ToResponseExpr != "str({var})" {
			t.Errorf("ToResponseExpr = %q", b.ToResponseExpr)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindNumeric(ir.BindOpts{NotNull: false})
		if b.APIType != "Decimal | None" {
			t.Errorf("APIType = %q, want Decimal | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
