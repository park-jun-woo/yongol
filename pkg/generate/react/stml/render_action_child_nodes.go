//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what Action 컨텍스트에서 ChildNode를 Kind별로 분기하여 렌더링한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// renderActionChildNodes renders ChildNode slice in DOM order for action
// context. errVar names the action's error-message state (data-on-error);
// empty when the action declares no error slot. idPrefix is threaded to
// renderFieldJSX to form-scope each field's DOM id (BUG-127).
func renderActionChildNodes(nodes []stmlparser.ChildNode, formName, idPrefix, errVar string, indent int) []string {
	var lines []string
	for _, ch := range nodes {
		switch ch.Kind {
		case "bind":
			lines = append(lines, renderFieldJSX(*ch.Bind, formName, idPrefix, indent))
		case "static":
			lines = append(lines, renderStaticActionJSX(*ch.Static, formName, idPrefix, errVar, indent))
		}
	}
	return lines
}
