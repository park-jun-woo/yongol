//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindBoolean — BOOLEAN → SQLAlchemy Boolean / Python bool 바인딩

package types

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindBoolean(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindBoolean(ir.BindOpts{NotNull: true, DefaultLiteral: "false"})
		if b.Family != typemap.FamilyBoolean || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "Boolean" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "bool" {
			t.Errorf("APIType = %q, want bool", b.APIType)
		}
		if b.DefaultLiteral != "false" {
			t.Errorf("DefaultLiteral = %q", b.DefaultLiteral)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindBoolean(ir.BindOpts{NotNull: false})
		if !strings.Contains(b.APIType, "bool") || b.APIType == "bool" {
			t.Errorf("expected nullable bool APIType, got %q", b.APIType)
		}
	})
}
