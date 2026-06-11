//ff:func feature=generate type=test control=sequence
//ff:what collectDocumentTitles — sitemap 부재 nil / 라벨·앱명 결합 / 중첩 노드 / 앱명 없는 라벨 단독 검증

package generate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectDocumentTitles(t *testing.T) {
	sitemap := &stml.SitemapSpec{Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
		{Page: "dashboard", Label: "대시보드", Menu: true},
		{Label: "건물 관리", Menu: true, Children: []stml.SitemapNode{ // group — no title of its own
			{Page: "building-detail", Label: "건물 상세", Menu: true},
		}},
	}}}}

	t.Run("nil without a sitemap — no title effects emitted", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Metadata: manifest.Metadata{Name: "zenflow"}}}
		if got := collectDocumentTitles(fs); got != nil {
			t.Errorf("titles = %v, want nil", got)
		}
		if got := collectDocumentTitles(nil); got != nil {
			t.Errorf("titles(nil) = %v, want nil", got)
		}
	})

	t.Run("labels join the manifest app name, nested nodes included", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &manifest.ProjectConfig{Metadata: manifest.Metadata{Name: "zenflow"}},
			Sitemap:  sitemap,
		}
		got := collectDocumentTitles(fs)
		if got["dashboard"] != "대시보드 · zenflow" || got["building-detail"] != "건물 상세 · zenflow" {
			t.Errorf("titles = %v", got)
		}
		if len(got) != 2 {
			t.Errorf("group labels must not produce entries, got %v", got)
		}
	})

	t.Run("missing app name keeps the bare label", func(t *testing.T) {
		fs := &yongol.Fullstack{Sitemap: sitemap}
		if got := collectDocumentTitles(fs); got["dashboard"] != "대시보드" {
			t.Errorf("titles = %v, want bare labels", got)
		}
	})
}
