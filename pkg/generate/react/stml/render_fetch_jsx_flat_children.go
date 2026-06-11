//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what FetchBlock의 flat 슬라이스(Binds, Eaches, States, Components) JSX를 생성한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// renderFetchJSXFlatChildren renders flat slices for backward compatibility.
func renderFetchJSXFlatChildren(f stmlparser.FetchBlock, alias string, indent int, noBodyOps map[string]bool, ctx bindCtx) []string {
	var lines []string
	for _, b := range f.Binds {
		lines = append(lines, renderBindJSX(b, alias, indent, ctx))
	}
	for _, e := range f.Eaches {
		lines = append(lines, renderEachJSX(e, alias, indent, noBodyOps, ctx))
	}
	for _, s := range f.States {
		lines = append(lines, renderStateJSX(s, alias, indent, noBodyOps, ctx))
	}
	for _, c := range f.Components {
		lines = append(lines, renderComponentJSX(c, alias, indent))
	}
	return lines
}
