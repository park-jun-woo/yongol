//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what ChildNode 슬라이스를 Kind별로 분기하여 DOM 순서대로 JSX를 렌더링한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// renderChildNodes renders ChildNode slice in DOM order for fetch context.
func renderChildNodes(nodes []stmlparser.ChildNode, dataVar, itemVar string, indent int, noBodyOps map[string]bool) []string {
	var lines []string
	for _, ch := range nodes {
		switch ch.Kind {
		case "bind":
			lines = append(lines, renderBindJSX(*ch.Bind, dataVar, indent))
		case "each":
			lines = append(lines, renderEachJSX(*ch.Each, dataVar, indent))
		case "state":
			lines = append(lines, renderStateJSX(*ch.State, dataVar, indent, noBodyOps))
		case "component":
			lines = append(lines, renderComponentJSX(*ch.Component, dataVar, indent))
		case "static":
			lines = append(lines, renderStaticJSX(*ch.Static, dataVar, itemVar, indent, noBodyOps))
		case "action":
			lines = append(lines, renderActionJSX(*ch.Action, indent, noBodyOps))
		case "fetch":
			lines = append(lines, renderFetchJSX(*ch.Fetch, indent, noBodyOps))
		}
	}
	return lines
}
