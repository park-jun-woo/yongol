//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what collectSitemapNavs — body 직속 요소를 순회하며 <nav data-sitemap> 블록 수집, 비-nav 최상위 요소는 파싱 에러

package stml

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"golang.org/x/net/html"
)

// collectSitemapNavs walks the direct element children of <body>, parsing
// each <nav data-sitemap> into spec.Navs and rejecting any other top-level
// element — a typo'd or unwrapped sitemap tree must not silently vanish.
func collectSitemapNavs(body *html.Node, filename string, spec *SitemapSpec) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for c := body.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		if c.Data == "nav" && hasAttr(c, "data-sitemap") {
			spec.Navs = append(spec.Navs, parseSitemapNav(c, spec))
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    filename,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("sitemap top-level <%s> is not a <nav data-sitemap> block", c.Data),
			Advice:  "Wrap every sitemap tree in <nav data-sitemap> (optionally with data-layout / data-entry)",
		})
	}
	return diags
}
