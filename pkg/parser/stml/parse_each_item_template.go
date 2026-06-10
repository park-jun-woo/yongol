//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what data-each의 첫 번째 자식 요소를 반복 항목 템플릿으로 파싱
package stml

import "golang.org/x/net/html"

// parseEachItemTemplate finds the first element child as the item template.
// A data-link on the template element itself declares the whole row as a
// link to a page (page-flow Phase007) — recorded as RowLink while the
// template's children keep parsing as row cells.
func parseEachItemTemplate(n *html.Node, eb *EachBlock) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		eb.ItemTag = c.Data
		eb.ItemClassName = getAttr(c, "class")
		if getAttr(c, "data-link") != "" {
			lr := parseLinkRef(c)
			eb.RowLink = &lr
		}
		walkEachItemChildren(c, eb)
		break
	}
}
