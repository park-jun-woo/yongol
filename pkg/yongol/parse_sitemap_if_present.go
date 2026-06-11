//ff:func feature=orchestrator type=loader control=sequence
//ff:what STML 존재 시 frontend 디렉토리 직속 sitemap.html 파싱 — 진단은 수집, 성공 시 Sitemap 설정

package yongol

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// parseSitemapIfPresent parses the fixed-name sitemap.html directly under
// the STML frontend directory when it exists (plans/stml/sitemap Phase001).
// Called after STML page parsing — ParseDir excludes sitemap.html from the
// page list, this loader is its only consumer. An absent file leaves
// fs.Sitemap nil (the sitemap is an optional SSOT file).
func parseSitemapIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindSTML]
	if !ok {
		return
	}
	path := filepath.Join(d.Path, "sitemap.html")

	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		return
	}

	spec, diags := stml.ParseSitemap(path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.Sitemap = &spec
	}
}
