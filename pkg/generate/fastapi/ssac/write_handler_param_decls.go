//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeHandlerParamDecls — handler 함수의 path/body/query 파라미터 선언 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeHandlerParamDecls writes the path, body, and query parameter
// declarations of the FastAPI route handler signature.
func writeHandlerParamDecls(b *strings.Builder, plan *ir.ServicePlan, hasBody bool) {
	writeHandlerPathDecls(b, plan.PathParams)
	if hasBody {
		reqModel := pascalCase(plan.OperationID) + "Request"
		b.WriteString(fmt.Sprintf("    body: %s,\n", reqModel))
	}
	writeHandlerQueryDecls(b, plan.QueryParams)
}
