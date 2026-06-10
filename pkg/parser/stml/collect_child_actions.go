//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what collectChildActions — ChildNode 트리에서 모든 ActionBlock을 DOM 순서로 수집
package stml

// collectChildActions walks the ChildNode tree and collects all
// ActionBlocks in DOM order. It mirrors the react emitter's traversal
// (pkg/generate/react/stml collectAllActions) so route derivation sees
// the same actions the page emitter renders.
func collectChildActions(nodes []ChildNode) []ActionBlock {
	var actions []ActionBlock
	for _, ch := range nodes {
		switch ch.Kind {
		case "action":
			actions = append(actions, *ch.Action)
		case "fetch":
			actions = append(actions, collectChildActions(ch.Fetch.Children)...)
		case "state":
			actions = append(actions, collectChildActions(ch.State.Children)...)
		case "static":
			actions = append(actions, collectChildActions(ch.Static.Children)...)
		case "each":
			actions = append(actions, collectChildActions(ch.Each.Children)...)
		}
	}
	return actions
}
