//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what parseSitemapList — 사이트맵 <ul> 의 직접 <li> 자식들을 SitemapNode 목록으로 변환

package stml

import "golang.org/x/net/html"

// parseSitemapList converts the direct <li> children of a sitemap <ul>
// into SitemapNodes, preserving document order (document order = menu order).
func parseSitemapList(ul *html.Node, spec *SitemapSpec) []SitemapNode {
	var items []SitemapNode
	for c := ul.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "li" {
			items = append(items, parseSitemapItem(c, spec))
		}
	}
	return items
}
