//ff:func feature=validate type=test control=sequence topic=statemachine-structural
//ff:what TestRun — StateMachine 검증 전체 실행 집계 호출 검증

package statemachine

import (
	"testing"

	smparser "github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	// Diagrams pre-loaded -> st01Parse short-circuits, Run aggregates without error.
	fs := &yongol.Fullstack{StateDiagrams: []*smparser.StateDiagram{{}}}
	if diags := Run(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
