//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ChildNode 트리를 재귀 순회하며 LinkRef.TargetPattern을 설정한다 (RowLink 포함)
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// setLinkTargetsInChildren recursively sets LinkRef.TargetPattern on every
// link node (including data-each RowLink) reachable from the ChildNode tree.
func setLinkTargetsInChildren(children []stmlparser.ChildNode, routePatterns map[string]string) {
	for _, ch := range children {
		setLinkTargetForChild(ch, routePatterns)
	}
}
