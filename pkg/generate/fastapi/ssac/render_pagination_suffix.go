//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderPaginationSuffix — PaginationArgs → SQLAlchemy .limit()/.offset() 접미사 문자열

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPaginationSuffix builds the SQLAlchemy .limit()/.offset() suffix from
// the given pagination args.
func renderPaginationSuffix(pagArgs []ir.FieldArg) string {
	var suffix string
	for _, pa := range pagArgs {
		val := renderArgValue(pa)
		switch resolveArgKey(pa) {
		case "per_page", "limit":
			suffix += fmt.Sprintf(".limit(%s)", val)
		case "page_offset", "offset":
			suffix += fmt.Sprintf(".offset(%s)", val)
		}
	}
	return suffix
}
