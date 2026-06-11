//ff:func feature=gen-react type=util control=sequence
//ff:what breadcrumbCrumbFor — 사이트맵 노드 하나를 crumb 으로 변환 (라벨 + MenuRenderable 페이지만 href)

package react

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/validate/stml_openapi"
)

// breadcrumbCrumbFor converts one sitemap node into a breadcrumb crumb
// (plans/stml/sitemap Phase004). The label is the <li> text, falling back
// to the page name for labelless page nodes (same fallback as the
// document.title supply). The href exists only when the node names a page
// AND the shared stml_openapi.MenuRenderable judgment admits it at this
// depth — exactly the condition under which the validator counts the
// breadcrumb up-link as a reachability edge (DESIGN §4.10 edge (d));
// emission and validation must never disagree. Group labels, external
// links and non-renderable pages (required route param, depth > 2,
// data-menu="false") stay label-only. depth is 1-based.
func breadcrumbCrumbFor(node stml.SitemapNode, depth int, routePatterns map[string]string) breadcrumbCrumb {
	label := node.Label
	if label == "" && node.Page != "" {
		label = node.Page
	}
	c := breadcrumbCrumb{Label: label}
	if node.Page != "" && stml_openapi.MenuRenderable(node, depth, sitemapNodePatterns(node, routePatterns)) {
		c.Href = navLinkPath(node.Page, routePatterns)
	}
	return c
}
