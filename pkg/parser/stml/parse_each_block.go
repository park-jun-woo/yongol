//ff:func feature=stml-parse type=parser control=sequence
//ff:what data-each 요소에서 EachBlock 구성
package stml

import "golang.org/x/net/html"

// parseEachBlock builds an EachBlock from a data-each element.
func parseEachBlock(n *html.Node, field string) EachBlock {
	eb := EachBlock{
		Tag:       n.Data,
		ClassName: getAttr(n, "class"),
		Field:     field,
	}
	parseEachItemTemplate(n, &eb)
	return eb
}
