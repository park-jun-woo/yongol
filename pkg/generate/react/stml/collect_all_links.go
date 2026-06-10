//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what each 행 자식 트리에서 링크 노드를 수집한다 (static/state 래퍼 재귀, 행 링크 셀 렌더용)
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectAllLinks walks an each row's ChildNode tree and collects every
// link node in DOM order (direct and static/state-nested) so renderEachJSX
// can emit one trailing cell per link, mirroring the row-action cells.
func collectAllLinks(nodes []stmlparser.ChildNode) []stmlparser.LinkRef {
	var links []stmlparser.LinkRef
	for _, ch := range nodes {
		switch ch.Kind {
		case "link":
			links = append(links, *ch.Link)
		case "state":
			links = append(links, collectAllLinks(ch.State.Children)...)
		case "static":
			links = append(links, collectAllLinks(ch.Static.Children)...)
		}
	}
	return links
}
