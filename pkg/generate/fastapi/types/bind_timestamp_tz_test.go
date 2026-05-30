//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindTimestampTZ — FamilyTimestampTZ → SQLAlchemy DateTime(timezone=True) / Python datetime 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindTimestampTZ(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindTimestampTZ(ir.BindOpts{NotNull: true})
		if b.Family != typemap.FamilyTimestampTZ || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "DateTime(timezone=True)" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "datetime" {
			t.Errorf("APIType = %q, want datetime", b.APIType)
		}
		if b.ToResponseExpr != "{var}.isoformat()" {
			t.Errorf("ToResponseExpr = %q", b.ToResponseExpr)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindTimestampTZ(ir.BindOpts{NotNull: false})
		if b.APIType != "datetime | None" {
			t.Errorf("APIType = %q, want datetime | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
