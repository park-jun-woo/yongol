//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBindImports — Numeric/Bytea 의 DBImports 검증 + DefaultLiteral 전파 검증

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindImports(t *testing.T) {
	reg := NewRegistry()

	t.Run("Numeric_DBImports", func(t *testing.T) {
		b := reg.Bind(typemap.FamilyNumeric, ir.BindOpts{NotNull: true})
		if len(b.DBImports) != 1 {
			t.Fatalf("DBImports len = %d, want 1", len(b.DBImports))
		}
		if b.DBImports[0] != "Decimal from @prisma/client/runtime/library" {
			t.Errorf("DBImports[0] = %q, want %q", b.DBImports[0], "Decimal from @prisma/client/runtime/library")
		}
	})

	t.Run("Bytea_DBImports", func(t *testing.T) {
		b := reg.Bind(typemap.FamilyBytea, ir.BindOpts{NotNull: true})
		if len(b.DBImports) != 1 {
			t.Fatalf("DBImports len = %d, want 1", len(b.DBImports))
		}
		if b.DBImports[0] != "Buffer from buffer" {
			t.Errorf("DBImports[0] = %q, want %q", b.DBImports[0], "Buffer from buffer")
		}
	})

	t.Run("Integer_NoDBImports", func(t *testing.T) {
		b := reg.Bind(typemap.FamilyInteger, ir.BindOpts{NotNull: true})
		if len(b.DBImports) != 0 {
			t.Errorf("DBImports len = %d, want 0", len(b.DBImports))
		}
	})

	t.Run("DefaultLiteral_Propagation", func(t *testing.T) {
		b := reg.Bind(typemap.FamilyInteger, ir.BindOpts{
			NotNull:        true,
			DefaultLiteral: "42",
		})
		if b.DefaultLiteral != "42" {
			t.Errorf("DefaultLiteral = %q, want %q", b.DefaultLiteral, "42")
		}
	})

	t.Run("DefaultLiteral_Empty", func(t *testing.T) {
		b := reg.Bind(typemap.FamilyString, ir.BindOpts{NotNull: true})
		if b.DefaultLiteral != "" {
			t.Errorf("DefaultLiteral = %q, want %q", b.DefaultLiteral, "")
		}
	})
}
