//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindUUID — FamilyUUID → SQLAlchemy Uuid / Python uuid.UUID 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindUUID(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindUUID(ir.BindOpts{NotNull: true})
		if b.Family != typemap.FamilyUUID || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "Uuid" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "uuid.UUID" {
			t.Errorf("APIType = %q, want uuid.UUID", b.APIType)
		}
		if len(b.DBImports) != 1 || len(b.APIImports) != 1 {
			t.Errorf("imports = %v / %v", b.DBImports, b.APIImports)
		}
		if b.NilCheckExpr != "" {
			t.Errorf("NilCheckExpr = %q, want empty", b.NilCheckExpr)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindUUID(ir.BindOpts{NotNull: false})
		if b.APIType != "uuid.UUID | None" {
			t.Errorf("APIType = %q, want uuid.UUID | None", b.APIType)
		}
		if b.NilCheckExpr != "{var} is None" {
			t.Errorf("NilCheckExpr = %q", b.NilCheckExpr)
		}
	})
}
