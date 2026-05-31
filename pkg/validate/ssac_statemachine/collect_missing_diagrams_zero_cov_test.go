//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestCollectMissingDiagrams_ZeroCov(t *testing.T) {
	byID := map[string]*statemachine.StateDiagram{"present": {ID: "present"}}
	fn := ssac.ServiceFunc{Name: "Fn", Sequences: []ssac.Sequence{
		{Type: "state", DiagramID: "missing", Line: 3},
		{Type: "state", DiagramID: "present"},
		{Type: "get"},
	}}
	diags := collectMissingDiagrams(fn, byID)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if diags[0].File != "ssac/Fn.ssac" {
		t.Errorf("synth file = %q", diags[0].File)
	}
	// with FileName set
	fn2 := ssac.ServiceFunc{Name: "Fn", FileName: "f.ssac", Sequences: []ssac.Sequence{{Type: "state", DiagramID: "missing"}}}
	if d := collectMissingDiagrams(fn2, byID); d[0].File != "f.ssac" {
		t.Errorf("file = %q", d[0].File)
	}
}
