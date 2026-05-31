//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasPublishPlans_ZeroCov(t *testing.T) {
	with := map[string][]*ir.ServicePlan{"f": {{Ops: []ir.Op{{Kind: ir.OpPublish}}}}}
	if !hasPublishPlans(with) {
		t.Error("expected publish present")
	}
	without := map[string][]*ir.ServicePlan{"f": {{Ops: []ir.Op{{Kind: ir.OpGet}}}}}
	if hasPublishPlans(without) {
		t.Error("publish should be absent")
	}
}
