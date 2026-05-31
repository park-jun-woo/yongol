//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what collectSchemaModels — body 필드를 가진 plan 들의 Pydantic request 모델 클래스명 수집

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// collectSchemaModels returns the list of Pydantic request model class names
// needed by plans that have request body fields.
func collectSchemaModels(plans []*ir.ServicePlan) []string {
	var models []string
	seen := make(map[string]bool)
	for _, plan := range plans {
		method := strings.ToUpper(plan.HTTPMethod)
		hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0
		name := pascalCase(plan.OperationID) + "Request"
		if hasBody && !seen[name] {
			seen[name] = true
			models = append(models, name)
		}
	}
	return models
}
