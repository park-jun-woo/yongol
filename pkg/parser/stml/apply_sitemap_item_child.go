//ff:func feature=stml-parse type=parser control=sequence
//ff:what applySitemapItemChild — <li> 자식 요소 분기: 중첩 <ul> → Children, <a href> → 외부 링크(첫 번째만)

package stml

import "golang.org/x/net/html"

// applySitemapItemChild folds one child of a sitemap <li> into the node:
// a nested <ul> contributes Children plus the dynamic-group vocabulary it
// carries (applySitemapDynamicGroup — plans/stml/sitemap Phase007), the
// first <a href> becomes the external link (its text fills the label when
// the li has no direct text).
func applySitemapItemChild(c *html.Node, node *SitemapNode, spec *SitemapSpec) {
	if c.Type != html.ElementNode {
		return
	}
	if c.Data == "ul" {
		applySitemapDynamicGroup(c, node)
		node.Children = append(node.Children, parseSitemapList(c, spec)...)
		return
	}
	if c.Data != "a" || node.Href != "" {
		return
	}
	href := getAttr(c, "href")
	if href == "" {
		return
	}
	node.Href = href
	if node.Label == "" {
		node.Label = extractAllText(c)
	}
}
