//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderArgValueVar — LocVar FieldArg → source.col / source / col 표현식 렌더 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderArgValueVar(t *testing.T) {
	t.Run("source and SourceColumn -> source.col", func(t *testing.T) {
		a := ir.FieldArg{Source: "order", SourceColumn: "owner_id"}
		if got := renderArgValueVar(a, "owner_id"); got != "order.owner_id" {
			t.Errorf("got %q, want order.owner_id", got)
		}
	})

	t.Run("SourceColumn from Field fallback", func(t *testing.T) {
		a := ir.FieldArg{Source: "order", Field: ".OwnerId"}
		if got := renderArgValueVar(a, "x"); got != "order.owner_id" {
			t.Errorf("got %q, want order.owner_id", got)
		}
	})

	t.Run("source only, no column -> source", func(t *testing.T) {
		a := ir.FieldArg{Source: "order"}
		if got := renderArgValueVar(a, ""); got != "order" {
			t.Errorf("got %q, want order", got)
		}
	})

	t.Run("column only, no source -> col", func(t *testing.T) {
		a := ir.FieldArg{SourceColumn: "id"}
		if got := renderArgValueVar(a, "id"); got != "id" {
			t.Errorf("got %q, want id", got)
		}
	})
}
