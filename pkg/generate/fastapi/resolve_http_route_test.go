//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiHelpers — fastapi plan/package/route 헬퍼 검증 (Op 종류·외부 패키지 수집·라우트 해석)
package fastapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveHTTPRoute(t *testing.T) {
	// Nil OpenAPI doc -> no mutation.
	plan := &ir.ServicePlan{OperationID: "ListItems"}
	resolveHTTPRoute(plan, &yongol.Fullstack{})
	if plan.HTTPMethod != "" || plan.URLPath != "" {
		t.Errorf("nil doc should not mutate plan: %+v", plan)
	}

	// Matching operationId -> method + path resolved.
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/items/{id}", &openapi3.PathItem{
				Get: &openapi3.Operation{OperationID: "GetItem"},
			}),
		),
	}
	fs := &yongol.Fullstack{OpenAPIDoc: doc}
	p2 := &ir.ServicePlan{OperationID: "GetItem"}
	resolveHTTPRoute(p2, fs)
	if p2.HTTPMethod != "GET" || p2.URLPath != "/items/{id}" {
		t.Errorf("unexpected route resolution: %+v", p2)
	}

	// No match -> unchanged.
	p3 := &ir.ServicePlan{OperationID: "Nope"}
	resolveHTTPRoute(p3, fs)
	if p3.HTTPMethod != "" {
		t.Errorf("non-matching op should not mutate: %+v", p3)
	}
}
