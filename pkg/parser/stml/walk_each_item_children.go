//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what each 항목 템플릿의 자식을 순회
package stml

import "golang.org/x/net/html"

// walkEachItemChildren walks grandchildren of the item template element.
func walkEachItemChildren(c *html.Node, eb *EachBlock) {
	for gc := c.FirstChild; gc != nil; gc = gc.NextSibling {
		walkEachChildren(gc, eb)
	}
}
