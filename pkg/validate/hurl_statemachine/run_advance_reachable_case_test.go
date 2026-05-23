//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-statemachine
//ff:what runAdvanceReachableCase — advanceReachable 개별 케이스 검증

package hurl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func runAdvanceReachableCase(t *testing.T, d *statemachine.StateDiagram, op string, wantKeys []string) {
	t.Helper()
	reachable := map[string]bool{}
	advanceReachable(d, op, reachable)
	if len(reachable) != len(wantKeys) {
		t.Fatalf("got %d reachable, want %d; got=%v", len(reachable), len(wantKeys), reachable)
	}
	for _, k := range wantKeys {
		if !reachable[k] {
			t.Errorf("missing %q in reachable", k)
		}
	}
}
