//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderLiteralArg — FieldArg 리터럴 → quoted/unquoted TS 표현식 렌더 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderLiteralArg(t *testing.T) {
	if got := renderLiteralArg(ir.FieldArg{Literal: "active", IsQuoted: true}); got != "'active'" {
		t.Errorf("quoted: got %q, want 'active'", got)
	}
	if got := renderLiteralArg(ir.FieldArg{Literal: "42", IsQuoted: false}); got != "42" {
		t.Errorf("unquoted: got %q, want 42", got)
	}
}
