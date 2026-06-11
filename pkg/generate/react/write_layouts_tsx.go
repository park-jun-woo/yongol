//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what 모든 LayoutSpec에 대해 레이아웃 TSX 파일을 생성한다 (sitemap 존재 시 레이아웃별 메뉴 모델 구성)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// writeLayoutsTSX emits layout TSX files for all provided LayoutSpecs.
// routePatterns resolves data-nav page-name references and authMode wires
// the data-logout emission (page-flow Phase010). With a sitemap
// (plans/stml/sitemap Phase003) each layout gets its menu model from its
// own nav blocks (data-layout match, "" delegates to defaultLayout); a
// nil sitemap keeps the legacy data-nav path byte-identical. roleField is
// frontend.auth.role_field — the claim the data-roles menu filter reads
// (Phase005); "" disables role conditions (TM-47 blocks generate before an
// unwired data-roles ever reaches here). crumbLayouts is the
// crumbFieldLayouts set (Phase006) — the layouts hosting a
// data-crumb-field page get the dynamic crumb-label wiring; nil keeps
// every layout byte-identical.
func writeLayoutsTSX(srcDir string, layouts []stml.LayoutSpec, routePatterns map[string]string, authMode string, sitemap *stml.SitemapSpec, defaultLayout string, roleField string, crumbLayouts map[string]bool) error {
	for _, l := range layouts {
		var menu *sitemapMenu
		if sitemap != nil {
			menu = &sitemapMenu{Items: buildSitemapMenu(sitemapNavsForLayout(sitemap, l.Name, defaultLayout), routePatterns), RoleField: roleField, DynamicCrumb: crumbLayouts[l.Name]}
		}
		if err := writeLayoutTSX(srcDir, l, routePatterns, authMode, menu); err != nil {
			return err
		}
	}
	return nil
}
