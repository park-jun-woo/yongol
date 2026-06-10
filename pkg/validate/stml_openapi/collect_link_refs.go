//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what ChildNode 트리에서 data-link 참조를 each 컨텍스트(item 스키마)와 함께 재귀 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectLinkRefs walks a ChildNode slice carrying the TM-30 enclosing
// context (nearest data-fetch operationId, innermost data-each item
// schema) and collects every data-link reference — child links and
// data-each RowLinks alike — for TM-31/TM-32.
func collectLinkRefs(children []stml.ChildNode, opID string, itemFields map[string]bool, inEach bool, raif map[string]map[string]map[string]bool, out *[]linkRefCtx) {
	for _, ch := range children {
		collectLinkRefsChild(ch, opID, itemFields, inEach, raif, out)
	}
}
