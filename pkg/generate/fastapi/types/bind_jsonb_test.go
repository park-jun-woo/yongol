//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindJSONB — FamilyJSONB → SQLAlchemy JSONB / Python dict[str, Any] 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindJSONB(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindJSONB(ir.BindOpts{NotNull: true})
		if b.Family != typemap.FamilyJSONB || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "JSONB" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "dict[str, Any]" {
			t.Errorf("APIType = %q, want dict[str, Any]", b.APIType)
		}
		if len(b.DBImports) != 1 || len(b.APIImports) != 1 {
			t.Errorf("imports = %v / %v", b.DBImports, b.APIImports)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindJSONB(ir.BindOpts{NotNull: false})
		if b.APIType != "dict[str, Any] | None" {
			t.Errorf("APIType = %q, want dict[str, Any] | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
