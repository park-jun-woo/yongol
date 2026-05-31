//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiHelpers — fastapi plan/package/route 헬퍼 검증 (Op 종류·외부 패키지 수집·라우트 해석)
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasAuthPlans(t *testing.T) {
	withAuth := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{{Kind: ir.OpAuth}}}},
	}
	if !hasAuthPlans(withAuth) {
		t.Error("expected auth plans true")
	}
	without := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{{Kind: ir.OpCall}}}},
	}
	if hasAuthPlans(without) {
		t.Error("expected auth plans false")
	}
}
