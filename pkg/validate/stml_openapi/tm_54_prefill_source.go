//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-54 — 페이지 폼들의 data-prefill 소스/필드 커버리지 검증 (ERROR/WARNING)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm54PrefillSource validates every data-prefill declaration on a page
// (plans/gen/frontend Phase035, BUG-124). It walks all forms actually rendered
// — CollectChildActions(page.Children), since an edit form is typically nested
// inside the GET-by-id data-fetch that page.Actions misses — and delegates each
// to tm54PrefillSourceForAction. The valid prefill sources are the page's
// data-fetch operationIds (nested included), the same scope the react emitter
// gives the fetch data variables.
func tm54PrefillSource(page stml.PageSpec, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	pageFetchOps := collectPageFetchOps(page)
	var diags []diagnostic.Diagnostic
	for _, a := range stml.CollectChildActions(page.Children) {
		diags = append(diags, tm54PrefillSourceForAction(a, page.FileName, pageFetchOps, opMap)...)
	}
	return diags
}
