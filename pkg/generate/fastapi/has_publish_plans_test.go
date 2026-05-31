//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiHelpers — fastapi plan/package/route 헬퍼 검증 (Op 종류·외부 패키지 수집·라우트 해석)
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasPublishPlans(t *testing.T) {
	withPub := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{{Kind: ir.OpPublish}}}},
	}
	if !hasPublishPlans(withPub) {
		t.Error("expected publish plans true")
	}
	if hasPublishPlans(map[string][]*ir.ServicePlan{}) {
		t.Error("empty map should be false")
	}
}
