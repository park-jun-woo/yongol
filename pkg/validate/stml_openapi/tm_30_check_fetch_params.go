//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-30 보조 — data-fetch 파라미터의 item.* 소스를 each 외부 사용으로 진단

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm30CheckFetchParams flags item.* sources on a data-fetch element: a
// fetch is never row-scoped (the query runs at page level, before any row
// exists), so every item.* source there is an outside-each error.
func tm30CheckFetchParams(f stml.FetchBlock, file string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, p := range f.Params {
		if _, ok := itemParamField(p.Source); ok {
			diags = append(diags, tm30OutsideEachDiag(file, f.OperationID, p.Source))
		}
	}
	return diags
}
