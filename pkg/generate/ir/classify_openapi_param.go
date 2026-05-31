//ff:func feature=gen-ir type=util control=selection
//ff:what classifyOpenAPIParam -- 단일 OpenAPI 파라미터를 in=path/query 에 따라 plan 에 분류 추가

package ir

import "github.com/getkin/kin-openapi/openapi3"

// classifyOpenAPIParam adds a single OpenAPI parameter to plan and to the
// appropriate set (pathParams/queryParams) based on its `in` location.
func classifyOpenAPIParam(p *openapi3.ParameterRef, plan *ServicePlan, pathParams, queryParams map[string]bool) {
	if p == nil || p.Value == nil {
		return
	}
	switch p.Value.In {
	case "path":
		pathParams[p.Value.Name] = true
		plan.PathParams = append(plan.PathParams, p.Value.Name)
	case "query":
		queryParams[p.Value.Name] = true
		plan.QueryParams = append(plan.QueryParams, QueryParamMeta{
			Name:     p.Value.Name,
			Type:     paramSchemaType(p),
			Required: p.Value.Required,
		})
	}
}
