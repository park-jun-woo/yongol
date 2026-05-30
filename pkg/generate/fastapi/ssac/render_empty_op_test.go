//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderEmptyOp — EmptyOp → not 가드 HTTPException 렌더링

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderEmptyOp(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var b strings.Builder
		renderEmptyOp(&b, nil, "    ")
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("Basic", func(t *testing.T) {
		var b strings.Builder
		op := &ir.EmptyOp{VarName: "course", StatusCode: 404, Message: "not found"}
		renderEmptyOp(&b, op, "    ")
		want := "    if not course:\n" +
			`        raise HTTPException(status_code=404, detail="not found")` + "\n"
		if b.String() != want {
			t.Errorf("got %q, want %q", b.String(), want)
		}
	})
}
