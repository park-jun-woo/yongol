//ff:func feature=orchestrator type=loader control=sequence
//ff:what 한 도메인의 frontend/sitemap.html 을 os.Stat 가드 후 적재 — 단일 사이트 parseSitemapIfPresent 와 동일 시맨틱
package yongol

import (
	"os"
	"path/filepath"
)

// loadDomainSitemap loads one domain's optional frontend/sitemap.html into
// fs.DomainSitemaps. It is os.Stat-guarded and stores only on a clean parse,
// exactly like the single-site parseSitemapIfPresent loader.
func loadDomainSitemap(fs *Fullstack, name, frontDir string) {
	sitemapPath := filepath.Join(frontDir, "sitemap.html")
	info, err := os.Stat(sitemapPath)
	if err != nil || info.IsDir() {
		return
	}
	spec, diags := parseSitemap(sitemapPath)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.DomainSitemaps[name] = spec
	}
}
