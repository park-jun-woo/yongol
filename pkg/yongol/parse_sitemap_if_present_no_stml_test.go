//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSitemapIfPresent_NoSTML — STML 미탐지 시 아무것도 안 함 검증

package yongol

import "testing"

func TestParseSitemapIfPresent_NoSTML(t *testing.T) {
	fs := &Fullstack{}
	parseSitemapIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.Sitemap != nil {
		t.Errorf("expected nil Sitemap, got %+v", fs.Sitemap)
	}
}
