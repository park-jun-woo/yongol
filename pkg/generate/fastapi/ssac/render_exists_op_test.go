//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderExistsOp — ExistsOp → 존재 가드 HTTPException 렌더링

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderExistsOp(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var b strings.Builder
		renderExistsOp(&b, nil, "    ")
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("Basic", func(t *testing.T) {
		var b strings.Builder
		op := &ir.ExistsOp{VarName: "existing", StatusCode: 409, Message: "conflict"}
		renderExistsOp(&b, op, "    ")
		want := "    if existing:\n" +
			`        raise HTTPException(status_code=409, detail="conflict")` + "\n"
		if b.String() != want {
			t.Errorf("got %q, want %q", b.String(), want)
		}
	})
}
