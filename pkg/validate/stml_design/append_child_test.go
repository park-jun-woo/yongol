//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestWalkForOverrides — walkForOverrides DOM 순회 @override class 추출 분기 검증
package stml_design

import (
	"golang.org/x/net/html"
)

func appendChild(parent, child *html.Node) {
	if parent.FirstChild == nil {
		parent.FirstChild = child
		parent.LastChild = child
		return
	}
	parent.LastChild.NextSibling = child
	parent.LastChild = child
}
