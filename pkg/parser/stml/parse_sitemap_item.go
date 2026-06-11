//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what parseSitemapItem — 사이트맵 <li> 하나를 SitemapNode 로 변환 (속성 + 자식 요소 위임)

package stml

import "golang.org/x/net/html"

// parseSitemapItem builds one SitemapNode from a sitemap <li>: data-page /
// data-index / data-menu / data-icon / data-roles attributes, the direct
// text as Label, and the child elements (external <a href> link, nested
// <ul>) via applySitemapItemChild. An li with neither data-page nor href
// is a group label. The parser records data-page and href even when both
// are present — TM-39 rejects that contradiction with a precise diagnostic
// instead of the parser silently picking one. data-roles (Phase005) is the
// comma-separated role allowlist of the menu entry; TM-46/47 validate the
// values and their claim wiring. data-crumb-field (Phase006) names the
// fetch-response field that becomes the dynamic crumb label; the parser
// keeps it even on a group <li> — TM-39 rejects that misplacement and
// TM-50 validates the field against the page's first fetch.
func parseSitemapItem(li *html.Node, spec *SitemapSpec) SitemapNode {
	node := SitemapNode{
		Page:       getAttr(li, "data-page"),
		Label:      directText(li),
		Index:      hasAttr(li, "data-index"),
		Menu:       getAttr(li, "data-menu") != "false",
		Icon:       getAttr(li, "data-icon"),
		Roles:      splitRolesAttr(getAttr(li, "data-roles")),
		CrumbField: getAttr(li, "data-crumb-field"),
	}
	for c := li.FirstChild; c != nil; c = c.NextSibling {
		applySitemapItemChild(c, &node, spec)
	}
	return node
}
