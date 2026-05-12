//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what fetchOps 목록으로 queryClient.invalidateQueries 코드 문자열을 생성한다
package stml

import (
	"fmt"
	"strings"
)

// renderInvalidateExpr builds the invalidateQueries expression for the
// mutation's onSuccess handler. When fetchOps is empty it returns a blanket
// invalidation; otherwise it joins per-query invalidations.
func renderInvalidateExpr(fetchOps []string) string {
	if len(fetchOps) == 0 {
		return "queryClient.invalidateQueries()"
	}
	var parts []string
	for _, op := range fetchOps {
		parts = append(parts, fmt.Sprintf("queryClient.invalidateQueries({ queryKey: ['%s'] })", op))
	}
	return strings.Join(parts, "\n      ")
}
