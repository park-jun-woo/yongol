//ff:func feature=gen-fastapi type=util control=sequence
//ff:what planHasRequestBody — plan 이 POST/PUT/PATCH + body 필드를 가진 요청 본문 보유 여부

package fastapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// planHasRequestBody reports whether the plan is a POST/PUT/PATCH operation
// that carries at least one request body field.
func planHasRequestBody(plan *ir.ServicePlan) bool {
	method := strings.ToUpper(plan.HTTPMethod)
	mutating := method == "POST" || method == "PUT" || method == "PATCH"
	return mutating && len(plan.BodyFields) > 0
}
