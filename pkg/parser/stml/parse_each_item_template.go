//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what data-each의 첫 번째 자식 요소를 반복 항목 템플릿으로 파싱
package stml

import "golang.org/x/net/html"

// parseEachItemTemplate finds the first element child as the item template.
func parseEachItemTemplate(n *html.Node, eb *EachBlock) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			eb.ItemTag = c.Data
			eb.ItemClassName = getAttr(c, "class")
			walkEachItemChildren(c, eb)
			break
		}
	}
}
