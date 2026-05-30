//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiHelpers — fastapi plan/package/route 헬퍼 검증 (Op 종류·외부 패키지 수집·라우트 해석)

package fastapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestContainsOpKind(t *testing.T) {
	ops := []ir.Op{{Kind: ir.OpCall}, {Kind: ir.OpAuth}}
	if !containsOpKind(ops, ir.OpAuth) {
		t.Error("expected auth op present")
	}
	if containsOpKind(ops, ir.OpPublish) {
		t.Error("publish op should be absent")
	}
	if containsOpKind(nil, ir.OpAuth) {
		t.Error("nil ops should be false")
	}
}

func TestEnsurePkgMap(t *testing.T) {
	pm := map[string]map[string]bool{}
	ensurePkgMap(pm, "billing")
	if pm["billing"] == nil {
		t.Fatal("expected sub-map created")
	}
	pm["billing"]["X"] = true
	ensurePkgMap(pm, "billing") // idempotent
	if !pm["billing"]["X"] {
		t.Error("existing sub-map should be preserved")
	}
}

func TestHasAuthPlans(t *testing.T) {
	withAuth := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{{Kind: ir.OpAuth}}}},
	}
	if !hasAuthPlans(withAuth) {
		t.Error("expected auth plans true")
	}
	without := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{{Kind: ir.OpCall}}}},
	}
	if hasAuthPlans(without) {
		t.Error("expected auth plans false")
	}
}

func TestHasPublishPlans(t *testing.T) {
	withPub := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{{Kind: ir.OpPublish}}}},
	}
	if !hasPublishPlans(withPub) {
		t.Error("expected publish plans true")
	}
	if hasPublishPlans(map[string][]*ir.ServicePlan{}) {
		t.Error("empty map should be false")
	}
}

func TestAddOpPackageRefAndCollect(t *testing.T) {
	plans := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Hold"}},
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "", Function: "Local"}}, // empty pkg skipped
			{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "auth", Function: "IsExpired"}},
			{Kind: ir.OpEval, Eval: nil}, // nil eval skipped
			{Kind: ir.OpAuth},            // unhandled kind
		}}},
	}
	pkgs := collectExternalPackages(plans)
	if len(pkgs) != 2 {
		t.Fatalf("expected 2 packages, got %d: %+v", len(pkgs), pkgs)
	}
	// Sorted: auth before billing.
	if pkgs[0].Name != "auth" || pkgs[1].Name != "billing" {
		t.Errorf("unexpected package order: %+v", pkgs)
	}
	if len(pkgs[0].Methods) != 1 || pkgs[0].Methods[0] != "IsExpired" {
		t.Errorf("unexpected auth methods: %+v", pkgs[0])
	}
}

func TestResolveProjectID(t *testing.T) {
	// No manifest -> "app".
	if got := resolveProjectID(&yongol.Fullstack{}); got != "app" {
		t.Errorf("expected app fallback, got %q", got)
	}
	// Manifest with name.
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Metadata: manifest.Metadata{Name: "zenflow"},
	}}
	if got := resolveProjectID(fs); got != "zenflow" {
		t.Errorf("expected zenflow, got %q", got)
	}
	// Manifest with empty name -> fallback.
	fs2 := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if got := resolveProjectID(fs2); got != "app" {
		t.Errorf("expected app for empty name, got %q", got)
	}
}

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
