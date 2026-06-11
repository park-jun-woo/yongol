//ff:func feature=validate type=util control=selection topic=stml-openapi
//ff:what 단일 ChildNode 를 Kind 별로 분기하여 each 행 액션을 수집한다 (each 진입 시 Actions 채집 + 하위 재귀)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectEachRowActionsChild handles one ChildNode for
// collectEachRowActions: an each block contributes its row Actions and
// recurses, every other container kind only recurses into its children —
// the same dispatch shape as collectLinkRefsChild.
func collectEachRowActionsChild(ch stml.ChildNode, out *[]stml.ActionBlock) {
	switch ch.Kind {
	case "each":
		*out = append(*out, ch.Each.Actions...)
		collectEachRowActions(ch.Each.Children, out)
	case "fetch":
		collectEachRowActions(ch.Fetch.Children, out)
	case "static":
		if ch.Static != nil {
			collectEachRowActions(ch.Static.Children, out)
		}
	case "state":
		if ch.State != nil {
			collectEachRowActions(ch.State.Children, out)
		}
	case "link":
		collectEachRowActions(ch.Link.Children, out)
	}
}
