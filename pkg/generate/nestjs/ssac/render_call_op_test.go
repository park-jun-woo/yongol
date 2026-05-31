//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderCallOp(t *testing.T) {
	var b strings.Builder
	renderCallOp(&b, nil, "  ")
	if b.String() != "" {
		t.Errorf("nil op should be empty")
	}
	b.Reset()
	renderCallOp(&b, &ir.CallOp{Function: "DoIt", ResultVar: "r", Args: []ir.FieldArg{{Literal: "1"}}}, "  ")
	if !strings.Contains(b.String(), "const r = await doIt(1);") {
		t.Errorf("result-bound call = %q", b.String())
	}
	b.Reset()
	renderCallOp(&b, &ir.CallOp{Package: "billing", Function: "DoIt"}, "  ")
	if !strings.Contains(b.String(), "await this.billingService.doIt();") {
		t.Errorf("void di call = %q", b.String())
	}
}
