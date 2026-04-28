//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-statemachine
//ff:what XOH-05 — 같은 파일 내 operation 호출 순서가 state machine 전이 규칙을 위반하지 않음

package hurl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh05StateTransitionOrder enforces XOH-05 (WARNING): within a single
// hurl file, the operations belonging to a state diagram must appear in
// an order that the diagram's transitions allow.
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
