//ff:func feature=stml-gen type=util control=sequence
//ff:what pageCrumbField — 페이지의 동적 crumb 라벨 필드 판정 (CrumbFields 등재 + fetch 보유 시에만)

package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// pageCrumbField returns the data-crumb-field of a page when the dynamic
// crumb-label wiring applies (plans/stml/sitemap Phase006): the sitemap
// lists the page with the attribute AND the page has a data-fetch to read
// the value from. A fetch-less declaration is TM-50's ERROR — the guard
// here only keeps emission total. "" means no wiring, keeping the
// Phase004/005 page output byte-identical.
func pageCrumbField(page stmlparser.PageSpec, crumbFields map[string]string) string {
	if len(page.Fetches) == 0 {
		return ""
	}
	return crumbFields[page.Name]
}
