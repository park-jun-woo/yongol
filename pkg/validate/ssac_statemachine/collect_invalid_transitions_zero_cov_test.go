//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestCollectInvalidTransitions_ZeroCov(t *testing.T) {
	diagram := &statemachine.StateDiagram{ID: "wf"}
	byID := map[string]*statemachine.StateDiagram{"wf": diagram}
	fn := ssac.ServiceFunc{Name: "Fn", FileName: "f.ssac", Sequences: []ssac.Sequence{
		{Type: "state", DiagramID: "wf", Transition: "bogus", Line: 4},
		{Type: "get"},
	}}
	diags := collectInvalidTransitions(fn, byID)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d", len(diags))
	}
}
