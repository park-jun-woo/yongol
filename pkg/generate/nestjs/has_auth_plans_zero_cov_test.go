//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasAuthPlans_ZeroCov(t *testing.T) {
	with := map[string][]*ir.ServicePlan{"f": {{Ops: []ir.Op{{Kind: ir.OpAuth}}}}}
	if !hasAuthPlans(with) {
		t.Error("expected auth present")
	}
	without := map[string][]*ir.ServicePlan{"f": {{Ops: []ir.Op{{Kind: ir.OpGet}}}}}
	if hasAuthPlans(without) {
		t.Error("auth should be absent")
	}
}
