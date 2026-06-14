//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what removeOps 목록으로 queryClient.removeQueries 코드 문자열을 생성한다 (delete onSuccess: 자기 GET 캐시 제거)
package stml

import (
	"fmt"
	"strings"
)

// renderRemoveQueriesExpr builds the removeQueries expression for a delete
// mutation's onSuccess handler. A delete drops the deleted item's own GET
// from the cache (instead of invalidating it, which would refetch a 404),
// then navigates away (BUG-132 132-2). An empty removeOps yields "".
func renderRemoveQueriesExpr(removeOps []string) string {
	if len(removeOps) == 0 {
		return ""
	}
	var parts []string
	for _, op := range removeOps {
		parts = append(parts, fmt.Sprintf("queryClient.removeQueries({ queryKey: ['%s'] })", op))
	}
	return strings.Join(parts, "\n      ")
}
