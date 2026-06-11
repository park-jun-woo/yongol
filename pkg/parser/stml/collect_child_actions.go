//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what CollectChildActions — ChildNode 트리에서 모든 ActionBlock을 DOM 순서로 수집
package stml

// CollectChildActions walks the ChildNode tree and collects all
// ActionBlocks in DOM order. It mirrors the react emitter's traversal
// (pkg/generate/react/stml collectAllActions) so route derivation and
// cross-validation (TM-52) see the same actions the page emitter renders.
func CollectChildActions(nodes []ChildNode) []ActionBlock {
	var actions []ActionBlock
	for _, ch := range nodes {
		switch ch.Kind {
		case "action":
			actions = append(actions, *ch.Action)
		case "fetch":
			actions = append(actions, CollectChildActions(ch.Fetch.Children)...)
		case "state":
			actions = append(actions, CollectChildActions(ch.State.Children)...)
		case "static":
			actions = append(actions, CollectChildActions(ch.Static.Children)...)
		case "each":
			actions = append(actions, CollectChildActions(ch.Each.Children)...)
		}
	}
	return actions
}
