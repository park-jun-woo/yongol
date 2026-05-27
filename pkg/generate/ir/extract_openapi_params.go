//ff:func feature=gen-ir type=util control=iteration dimension=2
//ff:what extractOpenAPIParams -- operationId 로 OpenAPI operation 탐색 → path/query/body 분류 + SuccessStatus 이식

package ir

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
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
		ops := map[string]*openapi3.Operation{
			"GET": pathItem.Get, "POST": pathItem.Post,
			"PUT": pathItem.Put, "DELETE": pathItem.Delete,
			"PATCH": pathItem.Patch,
		}
		for method, op := range ops {
			if op == nil || op.OperationID != operationID {
				continue
			}

			// Found matching operation.
			plan.HTTPMethod = method
			plan.URLPath = path
			if s := oapiparser.DeriveSuccessStatus(op, method); s != 0 {
				plan.SuccessStatus = s
			}

			// Collect path-level + operation-level parameters.
			allParams := append(pathItem.Parameters, op.Parameters...)
			for _, p := range allParams {
				if p == nil || p.Value == nil {
					continue
				}
				switch p.Value.In {
				case "path":
					pathParams[p.Value.Name] = true
					plan.PathParams = append(plan.PathParams, p.Value.Name)
				case "query":
					queryParams[p.Value.Name] = true
					typ := ""
					if p.Value.Schema != nil && p.Value.Schema.Value != nil && p.Value.Schema.Value.Type != nil {
						typ = p.Value.Schema.Value.Type.Slice()[0]
					}
					plan.QueryParams = append(plan.QueryParams, QueryParamMeta{
						Name:     p.Value.Name,
						Type:     typ,
						Required: p.Value.Required,
					})
				}
			}
			sort.Strings(plan.PathParams)
			sort.Slice(plan.QueryParams, func(i, j int) bool {
				return plan.QueryParams[i].Name < plan.QueryParams[j].Name
			})

			// Collect body fields.
			plan.BodyFields = extractBodyFields(op)
			sort.Slice(plan.BodyFields, func(i, j int) bool {
				return plan.BodyFields[i].Name < plan.BodyFields[j].Name
			})

			return pathParams, queryParams
		}
	}

	return pathParams, queryParams
}
