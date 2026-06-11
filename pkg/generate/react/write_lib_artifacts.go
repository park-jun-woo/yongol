//ff:func feature=gen-react type=generator control=sequence
//ff:what writeLibArtifacts — src/lib 산출물 일괄 방출 (utils.ts 상시 + sitemap 존재 시 breadcrumbs.ts/Breadcrumb.tsx)

package react

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// writeLibArtifacts emits the src/lib artifacts in one step of the
// frontend setup sequence: the unconditional shadcn cn() helper
// (writeLibUtils) and, with a sitemap present, the Phase004 breadcrumb
// pair (writeBreadcrumbs — src/lib/breadcrumbs.ts plus the shared
// components/ui/Breadcrumb.tsx; a nil sitemap writes nothing and the
// output stays byte-identical).
func writeLibArtifacts(srcDir string, fs *yongol.Fullstack, stmlPages []stml.PageSpec) error {
	if err := writeLibUtils(srcDir); err != nil {
		return err
	}
	return writeBreadcrumbs(srcDir, fsSitemap(fs), navRoutePatterns(stmlPages))
}
