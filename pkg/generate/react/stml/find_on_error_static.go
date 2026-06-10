//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ChildNode 트리에서 data-on-error 마커 StaticElement를 찾는다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// findOnErrorStatic returns the first StaticElement carrying the
// data-on-error marker in the ChildNode tree, or nil when absent.
func findOnErrorStatic(nodes []stmlparser.ChildNode) *stmlparser.StaticElement {
	for _, ch := range nodes {
		if ch.Kind != "static" || ch.Static == nil {
			continue
		}
		if ch.Static.OnError {
			return ch.Static
		}
		if found := findOnErrorStatic(ch.Static.Children); found != nil {
			return found
		}
	}
	return nil
}
