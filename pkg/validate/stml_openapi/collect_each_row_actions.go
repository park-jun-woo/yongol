//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectEachRowActions — ChildNode 트리에서 data-each 행 단위 액션 블록을 재귀 수집 (page.Actions 미포함분)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectEachRowActions gathers the row-level data-action blocks living
// only on EachBlock.Actions — the parser appends top-level and
// static-nested actions to page.Actions, but each-row actions stay on
// their block. collectPageEdges adds them so a row action's data-redirect
// (e.g. delete → list) still counts as a reachability edge.
func collectEachRowActions(children []stml.ChildNode, out *[]stml.ActionBlock) {
	for _, ch := range children {
		collectEachRowActionsChild(ch, out)
	}
}
