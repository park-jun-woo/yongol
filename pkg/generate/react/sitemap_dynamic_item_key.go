//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what sitemapDynamicItemKey — 동적 그룹 항목의 React key 표현식 (첫 item.* link-param 소스, 없으면 "" = 인덱스)

package react

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// sitemapDynamicItemKey picks the React key expression for a dynamic menu
// group's items (plans/stml/sitemap Phase007): the first item.<Field>
// data-link-params source — the value that fills the target route segment,
// so it identifies the row. Without any (a target route without required
// segments needs no params) the emitter falls back to the positional
// index, the page each-emitter's own fallback.
func sitemapDynamicItemKey(params []stml.LinkParamBind) string {
	for _, p := range params {
		if strings.HasPrefix(p.Source, "item.") {
			return p.Source
		}
	}
	return ""
}
