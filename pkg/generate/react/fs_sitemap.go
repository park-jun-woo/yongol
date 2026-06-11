//ff:func feature=gen-react type=accessor control=sequence
//ff:what fsSitemap — Fullstack.Sitemap 접근자 (nil-safe)

package react

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// fsSitemap returns the parsed frontend/sitemap.html or nil (absent file
// or nil Fullstack — every sitemap-keyed behavior then stays off).
func fsSitemap(fs *yongol.Fullstack) *stml.SitemapSpec {
	if fs == nil {
		return nil
	}
	return fs.Sitemap
}
