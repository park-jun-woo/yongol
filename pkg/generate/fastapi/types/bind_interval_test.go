//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindInterval — FamilyInterval → SQLAlchemy Interval / Python timedelta 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindInterval(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindInterval(ir.BindOpts{NotNull: true})
		if b.Family != typemap.FamilyInterval || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "Interval" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "timedelta" {
			t.Errorf("APIType = %q, want timedelta", b.APIType)
		}
		if len(b.APIImports) != 1 {
			t.Errorf("APIImports = %v", b.APIImports)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindInterval(ir.BindOpts{NotNull: false})
		if b.APIType != "timedelta | None" {
			t.Errorf("APIType = %q, want timedelta | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
