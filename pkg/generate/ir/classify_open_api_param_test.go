//ff:func feature=gen-ir type=test control=sequence
//ff:what TestClassifyOpenAPIParam — TestClassifyOpenAPIParam -- OpenAPI 파라미터 path/query 분류·nil/기타 위치 분기 검증

package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestClassifyOpenAPIParam(t *testing.T) {
	t.Run("Path", func(t *testing.T) {
		plan := &ServicePlan{}
		pathSet := map[string]bool{}
		querySet := map[string]bool{}
		classifyOpenAPIParam(paramRef("id", "path", true, "integer"), plan, pathSet, querySet)
		if !pathSet["id"] || len(plan.PathParams) != 1 || plan.PathParams[0] != "id" {
			t.Errorf("path not classified: set=%v plan=%v", pathSet, plan.PathParams)
		}
	})

	t.Run("Query", func(t *testing.T) {
		plan := &ServicePlan{}
		pathSet := map[string]bool{}
		querySet := map[string]bool{}
		classifyOpenAPIParam(paramRef("limit", "query", true, "integer"), plan, pathSet, querySet)
		if !querySet["limit"] || len(plan.QueryParams) != 1 {
			t.Fatalf("query not classified: set=%v plan=%v", querySet, plan.QueryParams)
		}
		qp := plan.QueryParams[0]
		if qp.Name != "limit" || qp.Type != "integer" || !qp.Required {
			t.Errorf("QueryParam = %+v, want {limit integer true}", qp)
		}
	})

	t.Run("NilRefNoop", func(t *testing.T) {
		plan := &ServicePlan{}
		classifyOpenAPIParam(nil, plan, map[string]bool{}, map[string]bool{})
		classifyOpenAPIParam(&openapi3.ParameterRef{Value: nil}, plan, map[string]bool{}, map[string]bool{})
		if len(plan.PathParams) != 0 || len(plan.QueryParams) != 0 {
			t.Error("nil ref should be a noop")
		}
	})

	t.Run("OtherLocationIgnored", func(t *testing.T) {
		plan := &ServicePlan{}
		classifyOpenAPIParam(paramRef("X-Token", "header", false, "string"), plan, map[string]bool{}, map[string]bool{})
		if len(plan.PathParams) != 0 || len(plan.QueryParams) != 0 {
			t.Error("header param should be ignored")
		}
	})
}
