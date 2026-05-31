//ff:func feature=gen-fastapi type=util control=sequence
//ff:what handlerQueryParamDecl — 단일 query 파라미터 메타 → handler 시그니처용 타입 선언 문자열

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// handlerQueryParamDecl returns the handler signature declaration for a single
// query parameter; optional parameters default to None.
func handlerQueryParamDecl(qp ir.QueryParamMeta) string {
	pyType := openAPITypeToPython(qp.Type)
	if qp.Required {
		return fmt.Sprintf("%s: %s", qp.Name, pyType)
	}
	return fmt.Sprintf("%s: %s | None = None", qp.Name, pyType)
}
