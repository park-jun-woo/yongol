//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-statemachine
//ff:what checkFileOrder — 한 파일의 entries 를 각 diagram 마다 checkDiagramOrder 로 검증

package hurl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// checkFileOrder walks one file's entries in order and reports any
// transition that arrives earlier than its required predecessor.
func checkFileOrder(entries []hurl.HurlEntry, opID map[string]string, diagrams []*statemachine.StateDiagram) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, d := range diagrams {
		diags = append(diags, checkDiagramOrder(entries, opID, d)...)
	}
	return diags
}
