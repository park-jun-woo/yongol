//ff:func feature=gen-react type=generator control=sequence
//ff:what writeBreadcrumbs — sitemap 존재 시 src/lib/breadcrumbs.ts + components/ui/Breadcrumb.tsx 방출 (부재 시 무방출)

package react

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// writeBreadcrumbs emits the breadcrumb artifacts of plans/stml/sitemap
// Phase004: src/lib/breadcrumbs.ts (the generate-time trail constants)
// and src/components/ui/Breadcrumb.tsx (the shared component the layouts
// place above their Outlet). A nil sitemap writes nothing — the
// sitemap-absent output stays byte-identical. With no layout to host it
// (layoutCount == 0, plans/stml/sitemap Phase008, BUG-129) the component
// would have no mount point — render_layout_tsx.go emits no layout, so
// imports are zero — and the pair becomes dead code; the gate skips it
// (TM-12/TM-41 ERROR already block the defaultLayout/data-layout +
// empty-layouts/ combinations before generate). Both files are emitted
// even when every trail is empty (all pages at depth 1): the layouts
// import the component unconditionally once a sitemap and a host exist.
// Any data-crumb-field declaration (Phase006) switches both artifacts to
// their dynamic form — keyed on the declaration, not on trail depth, so
// the component always accepts the label prop the wired layouts pass.
func writeBreadcrumbs(srcDir string, sitemap *stml.SitemapSpec, layoutCount int, routePatterns map[string]string) error {
	if sitemap == nil || layoutCount == 0 {
		return nil
	}
	libDir := filepath.Join(srcDir, "lib")
	if err := os.MkdirAll(libDir, 0o755); err != nil {
		return err
	}
	dynamic := len(stml.SitemapCrumbFields(sitemap)) > 0
	trails := buildBreadcrumbTrails(sitemap, routePatterns)
	if err := os.WriteFile(filepath.Join(libDir, "breadcrumbs.ts"), []byte(renderBreadcrumbsTS(trails, dynamic)), 0o644); err != nil {
		return err
	}
	uiDir := filepath.Join(srcDir, "components", "ui")
	if err := os.MkdirAll(uiDir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(uiDir, "Breadcrumb.tsx"), []byte(renderBreadcrumbComponent(dynamic)), 0o644)
}
