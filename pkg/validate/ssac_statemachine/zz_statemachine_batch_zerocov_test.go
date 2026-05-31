//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버

package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestBuildDiagramByID_ZeroCov(t *testing.T) {
	m := buildDiagramByID([]*statemachine.StateDiagram{{ID: "a"}, {ID: "b"}})
	if len(m) != 2 || m["a"] == nil || m["b"] == nil {
		t.Fatalf("buildDiagramByID = %v", m)
	}
}

func TestBuildFuncByName_ZeroCov(t *testing.T) {
	m := buildFuncByName([]ssac.ServiceFunc{{Name: "X"}, {Name: "Y"}})
	if len(m) != 2 || m["X"].Name != "X" {
		t.Fatalf("buildFuncByName = %v", m)
	}
}

func TestBuildFuncNameSet_ZeroCov(t *testing.T) {
	m := buildFuncNameSet([]ssac.ServiceFunc{{Name: "X"}})
	if !m["X"] || m["Z"] {
		t.Fatalf("buildFuncNameSet = %v", m)
	}
}

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

func TestCheckStateInputType_ZeroCov(t *testing.T) {
	g := &rule.Ground{Types: map[string]string{
		"SSaC.var.Fn.bad": "int64",
	}}
	fn := ssac.ServiceFunc{Name: "Fn", FileName: "f.ssac"}
	// "bad" resolves to int64 (non-string) → diag; "lit" is a quoted literal → string-compatible, no diag
	seq := ssac.Sequence{Type: "state", Line: 3, Inputs: map[string]string{"k": "bad"}}
	diags := checkStateInputType(g, fn, seq)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	// string-compatible input → no diag
	seq2 := ssac.Sequence{Type: "state", Line: 3, Inputs: map[string]string{"k": `"draft"`}}
	if d := checkStateInputType(g, fn, seq2); len(d) != 0 {
		t.Errorf("expected 0 diags for string literal, got %v", d)
	}
}

func TestCheckTransitionValidity_ZeroCov(t *testing.T) {
	diagram := &statemachine.StateDiagram{ID: "wf", Transitions: []statemachine.Transition{
		{From: "draft", To: "active", Event: "activate", Line: 2},
	}}
	byID := map[string]*statemachine.StateDiagram{"wf": diagram}
	fn := ssac.ServiceFunc{Name: "Fn", FileName: "f.ssac"}

	// non-state seq
	if _, ok := checkTransitionValidity(fn, ssac.Sequence{Type: "get"}, byID); ok {
		t.Error("non-state should be skipped")
	}
	// unknown diagram
	if _, ok := checkTransitionValidity(fn, ssac.Sequence{Type: "state", DiagramID: "nope"}, byID); ok {
		t.Error("unknown diagram should be skipped")
	}
	// valid transition
	if _, ok := checkTransitionValidity(fn, ssac.Sequence{Type: "state", DiagramID: "wf", Transition: "activate"}, byID); ok {
		t.Error("valid transition should not raise")
	}
	// invalid transition → diag
	d, ok := checkTransitionValidity(fn, ssac.Sequence{Type: "state", DiagramID: "wf", Transition: "bogus", Line: 9}, byID)
	if !ok {
		t.Fatal("invalid transition should raise")
	}
	if d.Line != 9 || d.OperationID != "Fn" {
		t.Errorf("diag = %+v", d)
	}
	// file synthesized when empty
	d2, ok := checkTransitionValidity(ssac.ServiceFunc{Name: "Fn"}, ssac.Sequence{Type: "state", DiagramID: "wf", Transition: "bogus"}, byID)
	if !ok || d2.File != "ssac/Fn.ssac" {
		t.Errorf("synth file = %q", d2.File)
	}
}

func TestCollectGuardStateFuncs_ZeroCov(t *testing.T) {
	funcs := []ssac.ServiceFunc{
		{Name: "WithState", Sequences: []ssac.Sequence{{Type: "state", DiagramID: "d"}}},
		{Name: "NoState", Sequences: []ssac.Sequence{{Type: "get"}}},
	}
	m := collectGuardStateFuncs(funcs)
	if !m["WithState"] || m["NoState"] {
		t.Fatalf("collectGuardStateFuncs = %v", m)
	}
}

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

func TestCollectMissingGuards_ZeroCov(t *testing.T) {
	funcByName := map[string]ssac.ServiceFunc{
		"activate": {Name: "activate", FileName: "a.ssac", Line: 2},
		"close":    {Name: "close"},
	}
	guards := map[string]bool{"close": true}
	// "activate": has func, not a guard → diag; "close": guard → skip; "missing": no func → skip
	diags := collectMissingGuards("wf", []string{"activate", "close", "missing"}, funcByName, guards)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if diags[0].File != "a.ssac" {
		t.Errorf("file = %q", diags[0].File)
	}
	// func with empty FileName → synthesized
	fb := map[string]ssac.ServiceFunc{"go": {Name: "go"}}
	d := collectMissingGuards("wf", []string{"go"}, fb, map[string]bool{})
	if d[0].File != "ssac/go.ssac" {
		t.Errorf("synth file = %q", d[0].File)
	}
}

func TestCollectUnmatchedEvents_ZeroCov(t *testing.T) {
	// nil diagram
	if d := collectUnmatchedEvents(nil, map[string]bool{}); d != nil {
		t.Error("nil diagram should return nil")
	}
	diagram := &statemachine.StateDiagram{ID: "wf", File: "states/wf.md", Transitions: []statemachine.Transition{
		{Event: "activate", Line: 2},
		{Event: "activate", Line: 3}, // duplicate event → first-occurrence only
		{Event: "close", Line: 4},
	}}
	funcNames := map[string]bool{"activate": true}
	diags := collectUnmatchedEvents(diagram, funcNames)
	if len(diags) != 1 || diags[0].Line != 4 {
		t.Fatalf("expected 1 diag at line 4, got %v", diags)
	}
	// empty File → synthesized
	d2 := &statemachine.StateDiagram{ID: "wf", Transitions: []statemachine.Transition{{Event: "x", Line: 1}}}
	if d := collectUnmatchedEvents(d2, map[string]bool{}); d[0].File != "states/wf.md" {
		t.Errorf("synth file = %q", d[0].File)
	}
}
