//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what each 행 내부의 data-link 요소를 파싱 (자식 bind는 eb에 등록, ChildNode는 링크에 보존)
package stml

import "golang.org/x/net/html"

// parseLinkInEach parses a data-link element inside a data-each row. Child
// binds register with the enclosing EachBlock for validation while their
// ChildNodes stay under the link for codegen (the <Link> wraps them).
func parseLinkInEach(n *html.Node, eb *EachBlock) LinkRef {
	lr := parseLinkRef(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		lr.Children = append(lr.Children, dispatchStaticEachChild(c, eb))
	}
	return lr
}
