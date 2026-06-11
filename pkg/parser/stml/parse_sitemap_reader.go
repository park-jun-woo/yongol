//ff:func feature=stml-parse type=parser control=sequence
//ff:what ParseSitemapReader — io.Reader 의 sitemap HTML 을 SitemapSpec 으로 파싱, 빈 파일은 파싱 에러

package stml

import (
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"golang.org/x/net/html"
)

// ParseSitemapReader parses sitemap HTML from a reader and returns a
// SitemapSpec. Every top-level element must be a <nav data-sitemap> block
// (collectSitemapNavs); anything else — including a file with no block at
// all — is a parse error, so an empty or mistyped sitemap is visible
// immediately instead of silently declaring nothing.
func ParseSitemapReader(filename string, r io.Reader) (SitemapSpec, []diagnostic.Diagnostic) {
	doc, err := html.Parse(r)
	if err != nil {
		return SitemapSpec{}, []diagnostic.Diagnostic{{
			File:    filename,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "html parse: " + err.Error(),
		}}
	}

	spec := SitemapSpec{FileName: filename}
	var diags []diagnostic.Diagnostic
	if body := findBodyNode(doc); body != nil {
		diags = collectSitemapNavs(body, filename, &spec)
	}
	if len(diags) == 0 && len(spec.Navs) == 0 {
		diags = append(diags, diagnostic.Diagnostic{
			File:    filename,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "sitemap declares no <nav data-sitemap> block",
			Advice:  "Declare at least one <nav data-sitemap> with a nested <ul>/<li data-page> tree, or delete the empty sitemap.html",
		})
	}
	if len(diags) > 0 {
		return SitemapSpec{}, diags
	}
	return spec, nil
}
