//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindInet — FamilyInet → SQLAlchemy INET / Python str 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindInet(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindInet(ir.BindOpts{NotNull: true})
		if b.Family != typemap.FamilyInet || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "INET" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "str" {
			t.Errorf("APIType = %q, want str", b.APIType)
		}
		if len(b.DBImports) != 1 {
			t.Errorf("DBImports = %v", b.DBImports)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindInet(ir.BindOpts{NotNull: false})
		if b.APIType != "str | None" {
			t.Errorf("APIType = %q, want str | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
