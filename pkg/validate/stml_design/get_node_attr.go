//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what getNodeAttr — HTML 노드에서 지정 속성 값 반환
package stml_design

import (
	"golang.org/x/net/html"
)

// getNodeAttr returns the value of the named attribute on an html.Node.
func getNodeAttr(n *html.Node, name string) string {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
}
