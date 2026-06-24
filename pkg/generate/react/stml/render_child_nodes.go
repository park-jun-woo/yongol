//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what ChildNode 슬라이스를 Kind별로 분기하여 DOM 순서대로 JSX를 렌더링한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// renderChildNodes renders ChildNode slice in DOM order for fetch context.
func renderChildNodes(nodes []stmlparser.ChildNode, dataVar, itemVar string, indent int, noBodyOps map[string]bool, ctx bindCtx) []string {
	var lines []string
	for _, ch := range nodes {
		switch ch.Kind {
		case "bind":
			lines = append(lines, renderBindJSX(*ch.Bind, orDefault(itemVar, dataVar), indent, ctx))
		case "each":
			lines = append(lines, renderEachJSX(*ch.Each, dataVar, indent, noBodyOps, ctx))
		case "state":
			lines = append(lines, renderStateJSX(*ch.State, dataVar, itemVar, indent, noBodyOps, ctx))
		case "component":
			lines = append(lines, renderComponentJSX(*ch.Component, dataVar, indent))
		case "static":
			lines = append(lines, renderStaticJSX(*ch.Static, dataVar, itemVar, indent, noBodyOps, ctx))
		case "action":
			lines = append(lines, renderActionJSX(*ch.Action, indent, noBodyOps))
		case "fetch":
			lines = append(lines, renderFetchJSX(*ch.Fetch, indent, noBodyOps, ctx))
		case "link":
			lines = append(lines, renderLinkJSX(*ch.Link, dataVar, itemVar, indent, noBodyOps, ctx))
		}
	}
	return lines
}
