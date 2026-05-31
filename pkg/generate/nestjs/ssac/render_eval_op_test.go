//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderEvalOp(t *testing.T) {
	var b strings.Builder
	renderEvalOp(&b, nil, "  ")
	if b.String() != "" {
		t.Errorf("nil eval")
	}
	b.Reset()
	renderEvalOp(&b, &ir.EvalOp{Function: "IsExpired", Message: "expired", StatusCode: 400}, "  ")
	if !strings.Contains(b.String(), "if (await isExpired()) {") {
		t.Errorf("eval op = %q", b.String())
	}
}
