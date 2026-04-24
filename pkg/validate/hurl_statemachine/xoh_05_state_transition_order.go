//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-statemachine
//ff:what XOH-05 — 같은 파일 내 operation 호출 순서가 state machine 전이 규칙을 위반하지 않음

package hurl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh05StateTransitionOrder enforces XOH-05 (WARNING): within a single
// hurl file, the operations belonging to a state diagram must appear in
// an order that the diagram's transitions allow. The classic failure is
// calling `ExecuteWorkflow` before `ActivateWorkflow` (workflow still
// in `draft`) — runtime returns 409 and the test fails, but the SSOTs
// were already telling the user the order was wrong.
//
// Severity is WARNING because a hurl file can legitimately probe
// invalid transitions (4xx response assertions). Users suppress noise
// by arranging the happy-path order, or by splitting probe-scenarios
// into dedicated files.
func xoh05StateTransitionOrder(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.StateDiagrams) == 0 || len(fs.HurlEntries) == 0 {
		return nil
	}
	opIdByMethodPath := buildOpIDLookup(fs)
	if len(opIdByMethodPath) == 0 {
		return nil
	}
	byFile := groupByFile(fs.HurlEntries)
	var diags []diagnostic.Diagnostic
	for _, entries := range byFile {
		diags = append(diags, checkFileOrder(entries, opIdByMethodPath, fs.StateDiagrams)...)
	}
	return diags
}

// checkFileOrder walks one file's entries in order and reports any
// transition that arrives earlier than its required predecessor.
func checkFileOrder(entries []hurl.HurlEntry, opID map[string]string, diagrams []*statemachine.StateDiagram) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, d := range diagrams {
		for _, diag := range checkDiagramOrder(entries, opID, d) {
			diags = append(diags, diag)
		}
	}
	return diags
}

// checkDiagramOrder walks entries looking for transition labels from
// one diagram. For each observed transition, it verifies that some
// earlier entry supplied a predecessor transition — i.e. the current
// state just before this call is a state from which this transition
// can legally depart. The tolerance is loose: as long as *any* path
// through the diagram from the initial state to the current transition
// is possible given the observed order, we stay silent.
func checkDiagramOrder(entries []hurl.HurlEntry, opID map[string]string, d *statemachine.StateDiagram) []diagnostic.Diagnostic {
	if d == nil || d.InitialState == "" {
		return nil
	}
	reachable := map[string]bool{d.InitialState: true}
	var diags []diagnostic.Diagnostic
	for _, e := range entries {
		key := e.Method + " " + normPath(e.Path, nil, nil)
		op := opID[key]
		if op == "" {
			continue
		}
		from := transitionFrom(d, op)
		if from == "" {
			continue
		}
		if !reachable[from] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    e.File,
				Line:    e.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: "[XOH-05] call " + op + " requires state " + from + " which was not reached by prior hurl steps",
				Advice:  "Add the preceding transition(s) declared in states/" + d.ID + ".md before " + op + ", or accept this as a negative-path scenario",
			})
			continue
		}
		for _, t := range d.Transitions {
			if t.Event == op {
				reachable[t.To] = true
			}
		}
	}
	return diags
}

// transitionFrom returns the From state of the first transition whose
// label matches op. Empty string means the op is not part of this
// diagram.
func transitionFrom(d *statemachine.StateDiagram, op string) string {
	if d == nil {
		return ""
	}
	for _, t := range d.Transitions {
		if t.Event == op {
			return t.From
		}
	}
	return ""
}

// groupByFile bins entries by source file so each file's sequence is
// reasoned about independently.
func groupByFile(entries []hurl.HurlEntry) map[string][]hurl.HurlEntry {
	out := map[string][]hurl.HurlEntry{}
	for _, e := range entries {
		out[e.File] = append(out[e.File], e)
	}
	return out
}
