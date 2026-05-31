//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestBuildXsm27DiagAndAdvice_ZeroCov(t *testing.T) {
	diagram := &statemachine.StateDiagram{ID: "workflow", InitialState: "draft"}
	// with diagram + resultVar
	target := &statefulTarget{Resource: "workflow", Diagram: diagram, Model: "Workflow"}
	d := buildXsm27Diag(ssac.ServiceFunc{Name: "ArchiveWorkflow", FileName: "x.ssac", Line: 5}, target, "wf")
	if d.Level != "WARNING" && d.Message == "" {
		t.Fatalf("unexpected diag: %+v", d)
	}
	if d.File != "x.ssac" || d.Line != 5 {
		t.Errorf("loc = %q:%d", d.File, d.Line)
	}
	if d.Advice == "" {
		t.Error("expected non-empty advice")
	}
	// no file → synthesized; no diagram; empty resultVar → fallback to lowercase resource
	target2 := &statefulTarget{Resource: "Order"}
	d2 := buildXsm27Diag(ssac.ServiceFunc{Name: "CancelOrder"}, target2, "")
	if d2.File != "ssac/CancelOrder.ssac" {
		t.Errorf("synthesized file = %q", d2.File)
	}
	// advice direct, no diagram (initial empty path)
	adv := buildXsm27Advice("Fn", target2, "", "", "order")
	if adv == "" {
		t.Error("expected advice body")
	}
}
