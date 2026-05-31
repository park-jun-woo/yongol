//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what splitPaginationArgs -- FieldArg 목록을 where-clause 인자와 pagination 인자로 분리

package ir

import "github.com/park-jun-woo/yongol/pkg/util/caseconv"

// splitPaginationArgs separates pagination args from where-clause args. Keys
// from SSaC may be PascalCase (e.g. "PerPage"), so they are converted to
// snake_case for matching against paginationKeys.
func splitPaginationArgs(allArgs []FieldArg) (whereArgs, pagArgs []FieldArg) {
	for _, a := range allArgs {
		snake := caseconv.PascalToSnake(a.Key)
		if paginationKeys[snake] {
			pagArgs = append(pagArgs, a)
		} else {
			whereArgs = append(whereArgs, a)
		}
	}
	return whereArgs, pagArgs
}
