//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderEmptyExistsOp(t *testing.T) {
	var b strings.Builder
	renderEmptyOp(&b, nil, "  ")
	renderExistsOp(&b, nil, "  ")
	if b.String() != "" {
		t.Errorf("nil ops should be empty")
	}
	b.Reset()
	renderEmptyOp(&b, &ir.EmptyOp{VarName: "course", Message: "not found", StatusCode: 404}, "  ")
	if !strings.Contains(b.String(), "if (!course) {") || !strings.Contains(b.String(), "not found") {
		t.Errorf("empty op = %q", b.String())
	}
	b.Reset()
	renderExistsOp(&b, &ir.ExistsOp{VarName: "dup", Message: "conflict", StatusCode: 409}, "  ")
	if !strings.Contains(b.String(), "if (dup) {") || !strings.Contains(b.String(), "conflict") {
		t.Errorf("exists op = %q", b.String())
	}
}
