//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what fetchOps 목록으로 queryClient.invalidateQueries 코드 문자열을 생성한다
package stml

import (
	"fmt"
	"strings"
)

// renderInvalidateExpr builds the invalidateQueries expression for the
// mutation's onSuccess handler. It joins one per-query invalidation per
// fetch op. When fetchOps is empty it returns "" — there is no key to
// invalidate, so it emits nothing rather than a keyless
// invalidateQueries() that would wipe the entire app cache (BUG-132 132-3).
func renderInvalidateExpr(fetchOps []string) string {
	if len(fetchOps) == 0 {
		return ""
	}
	var parts []string
	for _, op := range fetchOps {
		parts = append(parts, fmt.Sprintf("queryClient.invalidateQueries({ queryKey: ['%s'] })", op))
	}
	return strings.Join(parts, "\n      ")
}
