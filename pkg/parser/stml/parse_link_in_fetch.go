//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what fetch 블록 내부의 data-link 요소를 파싱 (자식 bind는 fb에 등록, ChildNode는 링크에 보존)
package stml

import "golang.org/x/net/html"

// parseLinkInFetch parses a data-link element inside a data-fetch block.
// Child binds register with the enclosing FetchBlock for validation while
// their ChildNodes stay under the link for codegen.
func parseLinkInFetch(n *html.Node, fb *FetchBlock) LinkRef {
	lr := parseLinkRef(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		if child := dispatchStaticFetchChild(c, fb); child != nil {
			lr.Children = append(lr.Children, *child)
		}
	}
	return lr
}
