//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestIsPrecededByOverride — isPrecededByOverride @override 주석 선행 여부 분기 검증
package stml_design

import (
	"golang.org/x/net/html"
)

func link(prev, n *html.Node) {
	n.PrevSibling = prev
}
