//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderQueryParam — 단일 query 파라미터 메타 → Python 타입 파라미터 문자열 (필수/옵셔널)

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderQueryParam returns the Python parameter declaration for a single query
// parameter; optional parameters default to None.
func renderQueryParam(qp ir.QueryParamMeta) string {
	pyType := openAPITypeToPython(qp.Type)
	if qp.Required {
		return fmt.Sprintf("%s: %s", qp.Name, pyType)
	}
	return fmt.Sprintf("%s: %s | None = None", qp.Name, pyType)
}
