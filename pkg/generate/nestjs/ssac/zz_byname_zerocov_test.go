//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func bnPlan() *ir.ServicePlan {
	return &ir.ServicePlan{
		OperationID: "CreateOrder",
		HTTPMethod:  "POST",
		TriggerKind: ir.TriggerHTTP,
		URLPath:     "/orders/:id/items",
		BodyFields:  []ir.BodyFieldMeta{{Name: "qty"}},
		PathParams:  []string{"id"},
		QueryParams: []ir.QueryParamMeta{{Name: "page", Type: "integer"}},
		Ops: []ir.Op{
			{Kind: ir.OpAuth, Auth: &ir.AuthOp{}},
			{Kind: ir.OpPublish, Publish: &ir.PublishOp{}},
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing"}},
		},
	}
}

func TestAddExternalOpPackage_ZeroCov(t *testing.T) {
	seen := map[string]bool{}
	addExternalOpPackage(seen, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing"}})
	addExternalOpPackage(seen, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "audit"}})
	if !seen["billing"] || !seen["audit"] {
		t.Errorf("packages not collected: %v", seen)
	}
}

func TestCollectExternalOpsPackages_ZeroCov(t *testing.T) {
	ops := []ir.Op{
		{Kind: ir.OpCall, Call: &ir.CallOp{Package: "zeta"}},
		{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "alpha"}},
	}
	got := collectExternalOpsPackages(ops)
	if len(got) != 2 || got[0] != "alpha" || got[1] != "zeta" {
		t.Errorf("sorted packages wrong: %v", got)
	}
}

func TestControllerRouteSuffix_ZeroCov(t *testing.T) {
	plan := &ir.ServicePlan{URLPath: "/courses/{id}/enroll"}
	if got := controllerRouteSuffix(plan); got != ":id/enroll" {
		t.Errorf("suffix=%q", got)
	}
	if got := controllerRouteSuffix(&ir.ServicePlan{URLPath: "/courses"}); got != "" {
		t.Errorf("single segment should be empty, got %q", got)
	}
	if got := controllerRouteSuffix(&ir.ServicePlan{}); got != "" {
		t.Errorf("empty path should be empty")
	}
}

func TestHasAuthOp_ZeroCov(t *testing.T) {
	if !hasAuthOp([]ir.Op{{Kind: ir.OpAuth}}) {
		t.Error("expected auth op")
	}
	if hasAuthOp([]ir.Op{{Kind: ir.OpGet}}) {
		t.Error("unexpected auth op")
	}
}

func TestHasPublishOp_ZeroCov(t *testing.T) {
	if !hasPublishOp([]ir.Op{{Kind: ir.OpPublish}}) {
		t.Error("expected publish op")
	}
	if hasPublishOp([]ir.Op{{Kind: ir.OpGet}}) {
		t.Error("unexpected publish op")
	}
}

func TestHasVerifyPasswordOp_ZeroCov(t *testing.T) {
	if !hasVerifyPasswordOp([]ir.Op{{Kind: ir.OpVerifyPassword}}) {
		t.Error("expected verify-password op")
	}
	if hasVerifyPasswordOp([]ir.Op{{Kind: ir.OpGet}}) {
		t.Error("unexpected verify-password op")
	}
}

func TestHttpStatusConst_ZeroCov(t *testing.T) {
	cases := map[int]string{
		400: "BAD_REQUEST", 401: "UNAUTHORIZED", 403: "FORBIDDEN",
		404: "NOT_FOUND", 409: "CONFLICT", 422: "UNPROCESSABLE_ENTITY",
		500: "INTERNAL_SERVER_ERROR",
	}
	for code, want := range cases {
		if got := httpStatusConst(code); got != want {
			t.Errorf("httpStatusConst(%d)=%q want %q", code, got, want)
		}
	}
	if got := httpStatusConst(418); !strings.Contains(got, "BAD_REQUEST") {
		t.Errorf("default branch: %q", got)
	}
}

func TestLcFirst_ZeroCov(t *testing.T) {
	if lcFirst("Hello") != "hello" {
		t.Error("lcFirst failed")
	}
	if lcFirst("") != "" {
		t.Error("empty failed")
	}
}

func TestNestHTTPDecorator_ZeroCov(t *testing.T) {
	for m, want := range map[string]string{"get": "Get", "POST": "Post", "put": "Put", "delete": "Delete", "patch": "Patch", "weird": "Get"} {
		if got := nestHTTPDecorator(m); got != want {
			t.Errorf("nestHTTPDecorator(%q)=%q want %q", m, got, want)
		}
	}
}

func TestNestURLPath_ZeroCov(t *testing.T) {
	if got := nestURLPath("/orders/{id}/items/{itemId}"); got != "/orders/:id/items/:itemId" {
		t.Errorf("nestURLPath=%q", got)
	}
	if got := nestURLPath("/plain"); got != "/plain" {
		t.Errorf("plain path changed: %q", got)
	}
}

func TestRenderArgValueLegacy_ZeroCov(t *testing.T) {
	if got := renderArgValueLegacy(ir.FieldArg{Source: "request"}, ""); got != "params" {
		t.Errorf("request empty col: %q", got)
	}
	if got := renderArgValueLegacy(ir.FieldArg{Source: "currentUser"}, ""); got != "user" {
		t.Errorf("currentUser empty col: %q", got)
	}
	if got := renderArgValueLegacy(ir.FieldArg{Source: "request"}, "name"); got != "body.name" {
		t.Errorf("request col: %q", got)
	}
	if got := renderArgValueLegacy(ir.FieldArg{Source: "var"}, "x"); got != "var.x" {
		t.Errorf("var col: %q", got)
	}
}

func TestFieldName_ZeroCov(t *testing.T) {
	if got := fieldName(ir.FieldArg{Field: ".ID"}); got != "ID" {
		t.Errorf("dotted: %q", got)
	}
	if got := fieldName(ir.FieldArg{Field: "Name"}); got != "Name" {
		t.Errorf("plain: %q", got)
	}
}

func TestWriteConstructorParams_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeConstructorParams(&b, bnPlan())
	out := b.String()
	if !strings.Contains(out, "constructor(") || !strings.Contains(out, "QueueService") || !strings.Contains(out, "AuthzService") {
		t.Errorf("constructor params missing pieces: %s", out)
	}
}

func TestWriteControllerClass_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeControllerClass(&b, bnPlan())
	if !strings.Contains(b.String(), "@Controller(") {
		t.Errorf("controller class missing: %s", b.String())
	}
}

func TestWriteControllerImports_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeControllerImports(&b, bnPlan())
	out := b.String()
	if !strings.Contains(out, "Controller,") || !strings.Contains(out, "@nestjs/common") {
		t.Errorf("imports missing: %s", out)
	}
}

func TestWriteMethodBody_ZeroCov(t *testing.T) {
	var b strings.Builder
	p := bnPlan()
	p.UsesTransaction = true
	writeMethodBody(&b, p)
	if !strings.Contains(b.String(), "$transaction") {
		t.Errorf("transaction body missing: %s", b.String())
	}
	var b2 strings.Builder
	sub := bnPlan()
	sub.TriggerKind = ir.TriggerSubscribe
	writeMethodBody(&b2, sub)
	if !strings.Contains(b2.String(), "const message = payload") {
		t.Errorf("subscribe alias missing: %s", b2.String())
	}
}

func TestWriteServiceClass_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeServiceClass(&b, bnPlan())
	out := b.String()
	if !strings.Contains(out, "@Injectable()") || !strings.Contains(out, "Service {") {
		t.Errorf("service class missing: %s", out)
	}
}

func TestWriteSubscribeHandler_ZeroCov(t *testing.T) {
	var b strings.Builder
	p := bnPlan()
	p.Topic = "order.completed"
	writeSubscribeHandler(&b, p)
	if !strings.Contains(b.String(), "order.completed") {
		t.Errorf("subscribe handler missing: %s", b.String())
	}
}
