//ff:func feature=orchestrator type=loader control=sequence
//ff:what 경로의 sitemap.html 을 stml.ParseSitemap 으로 파싱하고 *SitemapSpec 포인터로 래핑 — 단일/도메인 공용
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// parseSitemap parses the sitemap.html at path via stml.ParseSitemap (which
// returns a value) and wraps the result in a pointer so callers can store it
// directly into the *SitemapSpec fields (Fullstack.Sitemap / DomainSitemaps).
func parseSitemap(path string) (*stml.SitemapSpec, []diagnostic.Diagnostic) {
	spec, diags := stml.ParseSitemap(path)
	return &spec, diags
}
