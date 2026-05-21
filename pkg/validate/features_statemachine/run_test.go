//ff:func feature=validate type=test control=sequence topic=features-statemachine
//ff:what Run — Features↔StateMachine 교차 검증 발동 테스트
package features_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestRun_Fires(t *testing.T) {
	ft := map[string]features.TableDef{
		"workflows": {States: []string{"draft", "active", "missing"}},
	}
	diagrams := []*statemachine.StateDiagram{
		{ID: "workflows", States: []string{"draft", "active"}},
	}
	fs := buildFSForXFS(ft, diagrams)
	diags := Run(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
}
