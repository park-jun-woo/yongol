//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ChildNode 트리를 순회하여 모든 ActionBlock을 수집한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectAllActions walks the ChildNode tree and collects all ActionBlocks.
func collectAllActions(nodes []stmlparser.ChildNode) []stmlparser.ActionBlock {
	var actions []stmlparser.ActionBlock
	for _, ch := range nodes {
		switch ch.Kind {
		case "action":
			actions = append(actions, *ch.Action)
		case "fetch":
			actions = append(actions, collectAllActions(ch.Fetch.Children)...)
		case "state":
			actions = append(actions, collectAllActions(ch.State.Children)...)
		case "static":
			actions = append(actions, collectAllActions(ch.Static.Children)...)
		case "each":
			actions = append(actions, collectAllActions(ch.Each.Children)...)
		}
	}
	return actions
}
