//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what extractOpenAPIParams -- operationId 로 OpenAPI operation 탐색 → path/query/body 분류 + SuccessStatus 이식

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// extractOpenAPIParams finds the OpenAPI operation matching the given
// operationId and populates plan.SuccessStatus, plan.PathParams,
// plan.QueryParams, and plan.BodyFields. Returns the path and query
// param sets for use by enrichFieldArgLocations.
func extractOpenAPIParams(fs *yongol.Fullstack, operationID string, plan *ServicePlan) (pathParams map[string]bool, queryParams map[string]bool) {
	pathParams = make(map[string]bool)
	queryParams = make(map[string]bool)

	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return pathParams, queryParams
	}

	for path, pathItem := range fs.OpenAPIDoc.Paths.Map() {
		if matchAndPopulateOperation(operationID, path, pathItem, plan, pathParams, queryParams) {
			return pathParams, queryParams
		}
	}

	return pathParams, queryParams
}
