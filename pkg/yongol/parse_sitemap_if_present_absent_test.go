//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSitemapIfPresent_AbsentFile — sitemap.html 부재 시 Sitemap nil + 무진단 검증

package yongol

import "testing"

func TestParseSitemapIfPresent_AbsentFile(t *testing.T) {
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindSTML: {Kind: KindSTML, Path: t.TempDir(), Presence: SSOTPopulated},
	}
	parseSitemapIfPresent(fs, has)
	if fs.Sitemap != nil || len(fs.ParseDiagnostics) != 0 {
		t.Errorf("expected nil Sitemap and no diags, got %+v / %+v", fs.Sitemap, fs.ParseDiagnostics)
	}
}
