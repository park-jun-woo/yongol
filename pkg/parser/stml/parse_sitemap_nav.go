//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what parseSitemapNav — <nav data-sitemap> 블록에서 layout/entry 속성과 <ul> 트리 수집

package stml

import "golang.org/x/net/html"

// parseSitemapNav builds a SitemapNav from one <nav data-sitemap> element:
// data-layout / data-entry on the nav itself and the nested <ul>/<li> tree
// as Items. Every attribute of the sitemap vocabulary — the dynamic-group
// set included (plans/stml/sitemap Phase007) — is first-class now, so no
// reserved-attribute pass remains.
func parseSitemapNav(nav *html.Node, spec *SitemapSpec) SitemapNav {
	out := SitemapNav{
		Layout: getAttr(nav, "data-layout"),
		Entry:  hasAttr(nav, "data-entry"),
	}
	for c := nav.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "ul" {
			out.Items = append(out.Items, parseSitemapList(c, spec)...)
		}
	}
	return out
}
