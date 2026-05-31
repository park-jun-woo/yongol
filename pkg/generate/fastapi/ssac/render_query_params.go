//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderQueryParams — query 파라미터 메타 목록 → Python 타입(필수/옵셔널) 파라미터 문자열 목록

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderQueryParams returns one Python parameter declaration per query
// parameter, with optional parameters defaulting to None.
func renderQueryParams(queryParams []ir.QueryParamMeta) []string {
	out := make([]string, 0, len(queryParams))
	for _, qp := range queryParams {
		out = append(out, renderQueryParam(qp))
	}
	return out
}
