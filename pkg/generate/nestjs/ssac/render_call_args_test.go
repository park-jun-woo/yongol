//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderCallArgs(t *testing.T) {
	got := renderCallArgs([]ir.FieldArg{
		{Literal: "1"},
		{Literal: "x", IsQuoted: true},
	})
	if got != "1, 'x'" {
		t.Errorf("renderCallArgs = %q", got)
	}
}
