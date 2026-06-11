//ff:func feature=gen-react type=test control=sequence
//ff:what writeBreadcrumbs — sitemap 존재 시 breadcrumbs.ts + Breadcrumb.tsx 방출 / nil sitemap 무방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteBreadcrumbs(t *testing.T) {
	routePatterns := map[string]string{
		"building-list":   "/buildings",
		"building-detail": "/buildings/:BuildingID",
	}
	sitemap := &stml.SitemapSpec{Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
		{Page: "building-list", Label: "건물 목록", Menu: true, Children: []stml.SitemapNode{
			{Page: "building-detail", Label: "건물 상세", Menu: true},
		}},
	}}}}

	t.Run("sitemap present writes both artifacts", func(t *testing.T) {
		srcDir := t.TempDir()
		if err := writeBreadcrumbs(srcDir, sitemap, 1, routePatterns); err != nil {
			t.Fatalf("writeBreadcrumbs: %v", err)
		}
		lib, err := os.ReadFile(filepath.Join(srcDir, "lib", "breadcrumbs.ts"))
		if err != nil {
			t.Fatalf("breadcrumbs.ts missing: %v", err)
		}
		assertContains(t, string(lib), "'building-detail': [")
		assertContains(t, string(lib), "{ label: '건물 목록', href: '/buildings' },")
		comp, err := os.ReadFile(filepath.Join(srcDir, "components", "ui", "Breadcrumb.tsx"))
		if err != nil {
			t.Fatalf("Breadcrumb.tsx missing: %v", err)
		}
		assertContains(t, string(comp), "export function Breadcrumb()")
	})

	t.Run("data-crumb-field switches both artifacts to the dynamic form", func(t *testing.T) {
		srcDir := t.TempDir()
		dyn := &stml.SitemapSpec{Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "building-list", Label: "건물 목록", Menu: true, Children: []stml.SitemapNode{
				{Page: "building-detail", Label: "건물 상세", Menu: true, CrumbField: "building_name"},
			}},
		}}}}
		if err := writeBreadcrumbs(srcDir, dyn, 1, routePatterns); err != nil {
			t.Fatalf("writeBreadcrumbs: %v", err)
		}
		lib, err := os.ReadFile(filepath.Join(srcDir, "lib", "breadcrumbs.ts"))
		if err != nil {
			t.Fatalf("breadcrumbs.ts missing: %v", err)
		}
		assertContains(t, string(lib), "dynamic?: boolean")
		assertContains(t, string(lib), "{ label: '건물 상세', dynamic: true },")
		comp, err := os.ReadFile(filepath.Join(srcDir, "components", "ui", "Breadcrumb.tsx"))
		if err != nil {
			t.Fatalf("Breadcrumb.tsx missing: %v", err)
		}
		assertContains(t, string(comp), "export function Breadcrumb({ label }: { label?: string | null })")
	})

	t.Run("nil sitemap writes nothing — byte-identical backward compat", func(t *testing.T) {
		srcDir := t.TempDir()
		if err := writeBreadcrumbs(srcDir, nil, 1, routePatterns); err != nil {
			t.Fatalf("writeBreadcrumbs: %v", err)
		}
		if _, err := os.Stat(filepath.Join(srcDir, "lib", "breadcrumbs.ts")); !os.IsNotExist(err) {
			t.Error("breadcrumbs.ts must not exist without a sitemap")
		}
		if _, err := os.Stat(filepath.Join(srcDir, "components", "ui", "Breadcrumb.tsx")); !os.IsNotExist(err) {
			t.Error("Breadcrumb.tsx must not exist without a sitemap")
		}
	})

	t.Run("zero layouts writes nothing — no host for the breadcrumb (Phase008)", func(t *testing.T) {
		srcDir := t.TempDir()
		if err := writeBreadcrumbs(srcDir, sitemap, 0, routePatterns); err != nil {
			t.Fatalf("writeBreadcrumbs: %v", err)
		}
		if _, err := os.Stat(filepath.Join(srcDir, "lib", "breadcrumbs.ts")); !os.IsNotExist(err) {
			t.Error("breadcrumbs.ts must not exist without a layout host")
		}
		if _, err := os.Stat(filepath.Join(srcDir, "components", "ui", "Breadcrumb.tsx")); !os.IsNotExist(err) {
			t.Error("Breadcrumb.tsx must not exist without a layout host")
		}
	})
}
