//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderLiteralArg — renderLiteralArg quoted/unquoted 리터럴 Python 표현식 렌더 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderLiteralArg(t *testing.T) {
	if got := renderLiteralArg(ir.FieldArg{Literal: "active", IsQuoted: true}); got != `"active"` {
		t.Errorf("quoted: got %q, want %q", got, `"active"`)
	}
	if got := renderLiteralArg(ir.FieldArg{Literal: "42", IsQuoted: false}); got != "42" {
		t.Errorf("unquoted: got %q, want %q", got, "42")
	}
}
