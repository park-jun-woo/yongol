//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindBytea — BYTEA → SQLAlchemy LargeBinary / Python bytes 바인딩

package types

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindBytea(t *testing.T) {
	t.Run("NotNull", func(t *testing.T) {
		b := bindBytea(ir.BindOpts{NotNull: true, DefaultLiteral: "''"})
		if b.Family != typemap.FamilyBytea || !b.Supported {
			t.Errorf("unexpected binding: %+v", b)
		}
		if b.DBType != "LargeBinary" {
			t.Errorf("DBType = %q", b.DBType)
		}
		if b.APIType != "bytes" {
			t.Errorf("APIType = %q, want bytes", b.APIType)
		}
		if b.DefaultLiteral != "''" {
			t.Errorf("DefaultLiteral = %q", b.DefaultLiteral)
		}
	})
	t.Run("Nullable", func(t *testing.T) {
		b := bindBytea(ir.BindOpts{NotNull: false})
		if !strings.Contains(b.APIType, "bytes") || b.APIType == "bytes" {
			t.Errorf("expected nullable bytes APIType, got %q", b.APIType)
		}
	})
}
