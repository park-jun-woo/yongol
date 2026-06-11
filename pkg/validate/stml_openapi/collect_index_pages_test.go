//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what collectIndexPages — manifest index / sitemap data-index / data-route="/" 수집 + 중복·미실존 제거 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectIndexPages(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "home", FileName: "home.html", Route: "/"},
	}

	t.Run("all three vehicles, deduplicated", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Manifest.Frontend.Index = "dashboard"
		fs.Sitemap = sitemapWithIndexFixture("dashboard")
		got := collectIndexPages(fs)
		want := []string{"dashboard", "home"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("collectIndexPages = %v, want %v", got, want)
		}
	})

	t.Run("nonexistent names are dropped", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Manifest.Frontend.Index = "ghost"
		got := collectIndexPages(fs)
		want := []string{"home"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("collectIndexPages = %v, want %v", got, want)
		}
	})

	t.Run("nothing declared, no / mount → empty", func(t *testing.T) {
		fs := makeFS([]stml.PageSpec{{Name: "dashboard", FileName: "dashboard.html"}}, nil)
		if got := collectIndexPages(fs); len(got) != 0 {
			t.Errorf("collectIndexPages = %v, want empty", got)
		}
	})
}
