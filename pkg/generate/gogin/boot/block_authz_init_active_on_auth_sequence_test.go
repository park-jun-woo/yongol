//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockAuthzInit — OPA authz.Init(policyPath, ownerships) — DB 의존 없음
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockAuthzInit_ActiveOnAuthSequence(t *testing.T) {
	// Active predicate must be hasAuthSequence (gated, not always active).
	block := blockAuthzInit(&yongol.Fullstack{})
	if block.Active == nil {
		t.Fatalf("authz-init must carry an Active predicate")
	}
	if block.Active(&yongol.Fullstack{}) {
		t.Errorf("authz-init must be inactive without an @auth sequence")
	}
}
