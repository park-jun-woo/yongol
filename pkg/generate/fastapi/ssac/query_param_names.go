//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what queryParamNames — query 파라미터 메타 목록에서 이름만 추출

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// queryParamNames extracts the Name of each query parameter.
func queryParamNames(queryParams []ir.QueryParamMeta) []string {
	out := make([]string, 0, len(queryParams))
	for _, qp := range queryParams {
		out = append(out, qp.Name)
	}
	return out
}
