//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderResponseOpSourceCasing -- tsSourceExpr 으로 PascalCase → camelCase 변환 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderResponseOpSingleVar(t *testing.T) {
	op := &ir.ResponseOp{
		SingleVar: "course",
	}

	var b strings.Builder
	renderResponseOp(&b, op, "    ")
	got := b.String()

	want := "    return course;\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
