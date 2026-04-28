//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-statemachine
//ff:what checkDiagramOrder — 한 state diagram 에 대해 entries 의 전이 순서 검증

package hurl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// checkDiagramOrder walks entries looking for transition labels from
// one diagram. For each observed transition, it verifies that some
// earlier entry supplied a predecessor transition — i.e. the current
// state just before this call is a state from which this transition
// can legally depart.
func checkDiagramOrder(entries []hurl.HurlEntry, opID map[string]string, d *statemachine.StateDiagram) []diagnostic.Diagnostic {
	if d == nil || d.InitialState == "" {
		return nil
	}
	reachable := map[string]bool{d.InitialState: true}
	var diags []diagnostic.Diagnostic
	for _, e := range entries {
		diags = append(diags, inspectDiagramEntry(e, opID, d, reachable)...)
	}
	return diags
}
