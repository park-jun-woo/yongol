//ff:func feature=stml-parse type=test control=sequence
//ff:what TestCollectSitemapNavs — nav data-sitemap 수집 / nav 아닌 최상위 요소 에러 / nav without data-sitemap 거부 검증

package stml

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCollectSitemapNavs(t *testing.T) {
	t.Run("collects every nav data-sitemap in order", func(t *testing.T) {
		body := firstElementNode(t, `
<nav data-sitemap data-layout="app"><ul><li data-page="home">홈</li></ul></nav>
<nav data-sitemap data-layout="bare"><ul><li data-page="login">로그인</li></ul></nav>`, "body")
		var spec SitemapSpec
		diags := collectSitemapNavs(body, "sitemap.html", &spec)
		if len(diags) != 0 {
			t.Fatalf("expected no diags, got %+v", diags)
		}
		if len(spec.Navs) != 2 || spec.Navs[0].Layout != "app" || spec.Navs[1].Layout != "bare" {
			t.Errorf("Navs = %+v, want app then bare", spec.Navs)
		}
	})

	t.Run("non-nav top-level element is a parse error", func(t *testing.T) {
		body := firstElementNode(t, `<div><ul><li data-page="home">홈</li></ul></div>`, "body")
		var spec SitemapSpec
		diags := collectSitemapNavs(body, "sitemap.html", &spec)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %+v", diags)
		}
		d := diags[0]
		if d.File != "sitemap.html" || d.Phase != diagnostic.PhaseParse || d.Level != diagnostic.LevelError {
			t.Errorf("diag meta = %+v", d)
		}
		if !strings.Contains(d.Message, "<div>") {
			t.Errorf("Message should name the offending tag, got %q", d.Message)
		}
		if len(spec.Navs) != 0 {
			t.Errorf("Navs = %+v, want none", spec.Navs)
		}
	})

	t.Run("nav without data-sitemap is rejected too", func(t *testing.T) {
		body := firstElementNode(t, `<nav><ul><li data-page="home">홈</li></ul></nav>`, "body")
		var spec SitemapSpec
		diags := collectSitemapNavs(body, "sitemap.html", &spec)
		if len(diags) != 1 || !strings.Contains(diags[0].Message, "<nav>") {
			t.Fatalf("expected one <nav> error, got %+v", diags)
		}
	})
}
