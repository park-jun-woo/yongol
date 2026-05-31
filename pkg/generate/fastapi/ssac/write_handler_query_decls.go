//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what writeHandlerQueryDecls — handler 함수의 query 파라미터 선언 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeHandlerQueryDecls writes one typed query parameter declaration per
// query parameter.
func writeHandlerQueryDecls(b *strings.Builder, queryParams []ir.QueryParamMeta) {
	for _, qp := range queryParams {
		b.WriteString(fmt.Sprintf("    %s,\n", handlerQueryParamDecl(qp)))
	}
}
