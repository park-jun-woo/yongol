//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 정적 action 래퍼 내 컴포넌트의 자식을 순회
package stml

import "golang.org/x/net/html"

// walkStaticActionNestedChildren walks children of a component inside a static action wrapper.
func walkStaticActionNestedChildren(c *html.Node, ab *ActionBlock) {
	for gc := c.FirstChild; gc != nil; gc = gc.NextSibling {
		if gc.Type == html.ElementNode {
			walkStaticActionChild(gc, ab, nil)
		}
	}
}
