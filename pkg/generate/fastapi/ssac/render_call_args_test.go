//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderCallArgs — FieldArg 배열 → 쉼표 구분 인자 문자열

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderCallArgs(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		if got := renderCallArgs(nil); got != "" {
			t.Errorf("got %q, want empty", got)
		}
	})
	t.Run("Multiple", func(t *testing.T) {
		args := []ir.FieldArg{
			{Literal: "42"},
			{Literal: "active", IsQuoted: true},
			{Location: ir.LocBody, ColumnName: "title"},
		}
		want := `42, "active", body.title`
		if got := renderCallArgs(args); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
