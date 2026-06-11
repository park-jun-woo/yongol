//ff:func feature=stml-gen type=generator control=sequence
//ff:what FetchBlock의 자식 노드 JSX를 생성한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// renderFetchJSXBody generates the inner content of a fetch JSX block.
func renderFetchJSXBody(f stmlparser.FetchBlock, alias string, indent int, noBodyOps map[string]bool, ctx bindCtx) []string {
	var lines []string

	if len(f.Children) > 0 {
		lines = append(lines, renderChildNodes(f.Children, alias, "item", indent, noBodyOps, ctx)...)
	} else {
		lines = append(lines, renderFetchJSXFlatChildren(f, alias, indent, noBodyOps, ctx)...)
	}

	return lines
}
