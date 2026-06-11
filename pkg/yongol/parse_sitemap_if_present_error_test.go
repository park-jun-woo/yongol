//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSitemapIfPresent_ParseErrorCollected — 파싱 에러 시 Sitemap nil + 진단 수집 검증

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSitemapIfPresent_ParseErrorCollected(t *testing.T) {
	frontend := t.TempDir()
	if err := os.WriteFile(filepath.Join(frontend, "sitemap.html"), []byte("<div>oops</div>"), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindSTML: {Kind: KindSTML, Path: frontend, Presence: SSOTPopulated},
	}
	parseSitemapIfPresent(fs, has)
	if fs.Sitemap != nil {
		t.Errorf("expected nil Sitemap on parse error, got %+v", fs.Sitemap)
	}
	if len(fs.ParseDiagnostics) == 0 {
		t.Errorf("expected parse diagnostics to be collected")
	}
}
