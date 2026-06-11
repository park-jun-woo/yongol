//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-49 — frontend ON + 페이지 존재 + sitemap 부재 WARNING / sitemap 존재·frontend OFF·페이지 0 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM49SitemapAbsent(t *testing.T) {
	pages := []stml.PageSpec{{Name: "login", FileName: "login.html"}}

	t.Run("absent sitemap with frontend ON and pages fires", func(t *testing.T) {
		fs := makeFS(pages, nil)
		diags := tm49SitemapAbsent(fs)
		if got := countDiag(diags, "[TM-49]"); got != 1 {
			t.Fatalf("expected 1 TM-49, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelWarning {
			t.Errorf("Level = %v, want LevelWarning", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "sitemap.html") {
			t.Errorf("Message should name the missing file, got %q", diags[0].Message)
		}
	})

	t.Run("present sitemap is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html"}
		if diags := tm49SitemapAbsent(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("frontend off is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		off := false
		fs.Manifest.Frontend.Enabled = &off
		if diags := tm49SitemapAbsent(fs); len(diags) != 0 {
			t.Errorf("expected silence with frontend OFF, got %+v", diags)
		}
	})

	t.Run("zero pages is silent", func(t *testing.T) {
		fs := makeFS(nil, nil)
		if diags := tm49SitemapAbsent(fs); len(diags) != 0 {
			t.Errorf("expected silence with zero pages, got %+v", diags)
		}
	})
}
