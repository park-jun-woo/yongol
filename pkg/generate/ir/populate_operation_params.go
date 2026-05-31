//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what populateOperationParams -- operation 의 path/query 파라미터 + body 필드를 plan 에 수집·정렬

package ir

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

// populateOperationParams collects path-level and operation-level parameters
// plus body fields into plan, then sorts each collection deterministically.
func populateOperationParams(plan *ServicePlan, pathItem *openapi3.PathItem, op *openapi3.Operation, pathParams, queryParams map[string]bool) {
	allParams := append(pathItem.Parameters, op.Parameters...)
	for _, p := range allParams {
		classifyOpenAPIParam(p, plan, pathParams, queryParams)
	}
	sort.Strings(plan.PathParams)
	sort.Slice(plan.QueryParams, func(i, j int) bool {
		return plan.QueryParams[i].Name < plan.QueryParams[j].Name
	})

	plan.BodyFields = extractBodyFields(op)
	sort.Slice(plan.BodyFields, func(i, j int) bool {
		return plan.BodyFields[i].Name < plan.BodyFields[j].Name
	})
}
