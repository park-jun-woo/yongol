//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderServiceParams — ServicePlan 메타데이터 기반 메서드 파라미터 목록 생성

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderServiceParams produces the TypeScript parameter list for the service
// method based on the plan's trigger kind and OpenAPI metadata.
func renderServiceParams(plan *ir.ServicePlan) string {
	if plan.TriggerKind == ir.TriggerSubscribe {
		return "payload: any"
	}

	method := strings.ToUpper(plan.HTTPMethod)
	hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0
	hasQuery := len(plan.QueryParams) > 0
	hasPath := len(plan.PathParams) > 0

	// Defensive: SSaC Ops may reference body/path/query even when OpenAPI
	// metadata is absent or incomplete. Force flags when Ops contain FieldArgs
	// with the corresponding Location.
	if !hasBody && opsReferenceBody(plan.Ops) {
		hasBody = true
	}
	if !hasPath && opsReferencePath(plan.Ops) {
		hasPath = true
	}
	if !hasQuery && opsReferenceQuery(plan.Ops) {
		hasQuery = true
	}

	var params []string
	if hasPath {
		params = append(params, "params: any")
	}
	if hasBody {
		params = append(params, "body: any")
	}
	if hasQuery {
		params = append(params, "query: any")
	}
	params = append(params, "user?: any")

	return strings.Join(params, ", ")
}
