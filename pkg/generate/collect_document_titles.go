//ff:func feature=generate type=util control=iteration dimension=1
//ff:what collectDocumentTitles — sitemap 라벨 + manifest 앱명으로 페이지명 → document.title 맵 구성 (sitemap 부재 시 nil)

package generate

import "github.com/park-jun-woo/yongol/pkg/yongol"

// collectDocumentTitles builds the page-name → document.title table the
// page emitter consumes (plans/stml/sitemap Phase004): the sitemap label
// of every listed page joined with the manifest app name
// (metadata.Name — the existing project-name field) as
// "<label> · <app name>". nil without a sitemap — no title effects are
// then emitted, keeping the sitemap-absent output byte-identical; pages
// not listed in the sitemap simply have no entry, with the same effect.
func collectDocumentTitles(fs *yongol.Fullstack) map[string]string {
	if fs == nil || fs.Sitemap == nil {
		return nil
	}
	appName := ""
	if fs.Manifest != nil {
		appName = fs.Manifest.Metadata.Name
	}
	titles := map[string]string{}
	for _, nav := range fs.Sitemap.Navs {
		addDocumentTitles(nav.Items, appName, titles)
	}
	return titles
}
