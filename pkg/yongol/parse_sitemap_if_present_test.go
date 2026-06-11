//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSitemapIfPresent_Present — sitemap.html 존재 시 Sitemap 적재 검증

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSitemapIfPresent_Present(t *testing.T) {
	frontend := t.TempDir()
	html := `<nav data-sitemap data-layout="app"><ul><li data-page="login" data-index>로그인</li></ul></nav>`
	if err := os.WriteFile(filepath.Join(frontend, "sitemap.html"), []byte(html), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindSTML: {Kind: KindSTML, Path: frontend, Presence: SSOTPopulated},
	}
	parseSitemapIfPresent(fs, has)
	if fs.Sitemap == nil {
		t.Fatalf("expected Sitemap to be set (diags=%+v)", fs.ParseDiagnostics)
	}
	if len(fs.Sitemap.Navs) != 1 || fs.Sitemap.Navs[0].Layout != "app" {
		t.Errorf("Sitemap = %+v", fs.Sitemap)
	}
}
