//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveHTTPRoute_ZeroCov(t *testing.T) {
	// nil doc → no-op
	plan := &ir.ServicePlan{OperationID: "getUser"}
	resolveHTTPRoute(plan, &yongol.Fullstack{})
	if plan.HTTPMethod != "" {
		t.Error("expected no route with nil doc")
	}
	// matching doc → resolves method + path
	op := &openapi3.Operation{OperationID: "getUser"}
	paths := openapi3.NewPaths(openapi3.WithPath("/users/{id}", &openapi3.PathItem{Get: op}))
	fs := &yongol.Fullstack{OpenAPIDoc: &openapi3.T{Paths: paths}}
	resolveHTTPRoute(plan, fs)
	if plan.HTTPMethod != "GET" || plan.URLPath != "/users/:id" {
		t.Errorf("resolveHTTPRoute = %q %q", plan.HTTPMethod, plan.URLPath)
	}
}
