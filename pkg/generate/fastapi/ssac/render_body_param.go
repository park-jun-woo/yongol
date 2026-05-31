//ff:func feature=gen-fastapi type=util control=selection
//ff:what renderBodyParam — body 파라미터 선언 목록 생성 (fallback 시 dict, 아니면 Request 모델)

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderBodyParam returns the body parameter declaration list. When hasBody is
// false the list is empty. A fallback body is typed as dict; otherwise it is
// typed as the operation's Request model.
func renderBodyParam(plan *ir.ServicePlan, hasBody, bodyFallback bool) []string {
	switch {
	case !hasBody:
		return nil
	case bodyFallback:
		return []string{"body: dict"}
	default:
		reqModel := pascalCase(plan.OperationID) + "Request"
		return []string{fmt.Sprintf("body: %s", reqModel)}
	}
}
