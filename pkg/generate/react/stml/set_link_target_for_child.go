//ff:func feature=stml-gen type=util control=selection
//ff:what 단일 ChildNode를 Kind별로 분기하여 LinkRef.TargetPattern을 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// setLinkTargetForChild handles one ChildNode for setLinkTargetsInChildren:
// link nodes (and data-each RowLinks) get their target page's resolved
// route pattern, container kinds recurse.
func setLinkTargetForChild(ch stmlparser.ChildNode, routePatterns map[string]string) {
	switch ch.Kind {
	case "link":
		ch.Link.TargetPattern = routePatterns[ch.Link.TargetPage]
		setLinkTargetsInChildren(ch.Link.Children, routePatterns)
	case "fetch":
		setLinkTargetsInChildren(ch.Fetch.Children, routePatterns)
	case "each":
		if ch.Each.RowLink != nil {
			ch.Each.RowLink.TargetPattern = routePatterns[ch.Each.RowLink.TargetPage]
		}
		setLinkTargetsInChildren(ch.Each.Children, routePatterns)
	case "static":
		setLinkTargetsInChildren(ch.Static.Children, routePatterns)
	case "state":
		setLinkTargetsInChildren(ch.State.Children, routePatterns)
	}
}
