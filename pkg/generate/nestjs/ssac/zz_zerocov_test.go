//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestControllerRoutePrefix(t *testing.T) {
	if got := controllerRoutePrefix(&ir.ServicePlan{URLPath: "/courses/:id"}); got != "courses" {
		t.Errorf("URLPath = %q, want courses", got)
	}
	if got := controllerRoutePrefix(&ir.ServicePlan{URLPath: "/courses"}); got != "courses" {
		t.Errorf("single-seg = %q, want courses", got)
	}
	got := controllerRoutePrefix(&ir.ServicePlan{URLPath: "", Feature: "Auth"})
	if got != "auth" {
		t.Errorf("empty path fallback = %q, want auth", got)
	}
}

func TestFormatCallTarget(t *testing.T) {
	if got := formatCallTarget("", "HoldEscrow"); got != "holdEscrow" {
		t.Errorf("local = %q", got)
	}
	if got := formatCallTarget("billing", "HoldEscrow"); got != "this.billingService.holdEscrow" {
		t.Errorf("di = %q", got)
	}
}

func TestIsDeleteByPK(t *testing.T) {
	if !isDeleteByPK(nil) {
		t.Errorf("empty args should be PK")
	}
	if !isDeleteByPK([]ir.FieldArg{{Key: "id", IsPK: true}}) {
		t.Errorf("IsPK flag should be PK")
	}
	if !isDeleteByPK([]ir.FieldArg{{Key: "id"}}) {
		t.Errorf("key=id heuristic should be PK")
	}
	if isDeleteByPK([]ir.FieldArg{{Key: "slug"}}) {
		t.Errorf("non-pk key should not be PK")
	}
}

func TestRenderCallArgs(t *testing.T) {
	got := renderCallArgs([]ir.FieldArg{
		{Literal: "1"},
		{Literal: "x", IsQuoted: true},
	})
	if got != "1, 'x'" {
		t.Errorf("renderCallArgs = %q", got)
	}
}

func TestRenderCallOp(t *testing.T) {
	var b strings.Builder
	renderCallOp(&b, nil, "  ")
	if b.String() != "" {
		t.Errorf("nil op should be empty")
	}
	b.Reset()
	renderCallOp(&b, &ir.CallOp{Function: "DoIt", ResultVar: "r", Args: []ir.FieldArg{{Literal: "1"}}}, "  ")
	if !strings.Contains(b.String(), "const r = await doIt(1);") {
		t.Errorf("result-bound call = %q", b.String())
	}
	b.Reset()
	renderCallOp(&b, &ir.CallOp{Package: "billing", Function: "DoIt"}, "  ")
	if !strings.Contains(b.String(), "await this.billingService.doIt();") {
		t.Errorf("void di call = %q", b.String())
	}
}

func TestRenderDeleteOp(t *testing.T) {
	var b strings.Builder
	renderDeleteOp(&b, nil, "  ", "this.prisma")
	if b.String() != "" {
		t.Errorf("nil op")
	}
	b.Reset()
	renderDeleteOp(&b, &ir.DeleteOp{Model: "Course", Args: []ir.FieldArg{{Key: "id", IsPK: true}}}, "  ", "this.prisma")
	if !strings.Contains(b.String(), "this.prisma.course.delete(") {
		t.Errorf("pk delete = %q", b.String())
	}
	b.Reset()
	renderDeleteOp(&b, &ir.DeleteOp{Model: "Course", Args: []ir.FieldArg{{Key: "slug"}}}, "  ", "this.prisma")
	if !strings.Contains(b.String(), "deleteMany(") {
		t.Errorf("non-pk deleteMany = %q", b.String())
	}
}

func TestRenderEmptyExistsOp(t *testing.T) {
	var b strings.Builder
	renderEmptyOp(&b, nil, "  ")
	renderExistsOp(&b, nil, "  ")
	if b.String() != "" {
		t.Errorf("nil ops should be empty")
	}
	b.Reset()
	renderEmptyOp(&b, &ir.EmptyOp{VarName: "course", Message: "not found", StatusCode: 404}, "  ")
	if !strings.Contains(b.String(), "if (!course) {") || !strings.Contains(b.String(), "not found") {
		t.Errorf("empty op = %q", b.String())
	}
	b.Reset()
	renderExistsOp(&b, &ir.ExistsOp{VarName: "dup", Message: "conflict", StatusCode: 409}, "  ")
	if !strings.Contains(b.String(), "if (dup) {") || !strings.Contains(b.String(), "conflict") {
		t.Errorf("exists op = %q", b.String())
	}
}

func TestRenderEvalOp(t *testing.T) {
	var b strings.Builder
	renderEvalOp(&b, nil, "  ")
	if b.String() != "" {
		t.Errorf("nil eval")
	}
	b.Reset()
	renderEvalOp(&b, &ir.EvalOp{Function: "IsExpired", Message: "expired", StatusCode: 400}, "  ")
	if !strings.Contains(b.String(), "if (await isExpired()) {") {
		t.Errorf("eval op = %q", b.String())
	}
}

func TestRenderPostOp(t *testing.T) {
	var b strings.Builder
	renderPostOp(&b, nil, "  ", "this.prisma")
	if b.String() != "" {
		t.Errorf("nil post")
	}
	b.Reset()
	renderPostOp(&b, &ir.PostOp{VarName: "course", Model: "Course", Args: []ir.FieldArg{{Key: "title", ColumnName: "title", Source: "body.title"}}}, "  ", "this.prisma")
	if !strings.Contains(b.String(), "const course = await this.prisma.course.create(") {
		t.Errorf("post op = %q", b.String())
	}
}

func TestRenderPublishOp(t *testing.T) {
	var b strings.Builder
	renderPublishOp(&b, nil, "  ")
	if b.String() != "" {
		t.Errorf("nil publish")
	}
	b.Reset()
	renderPublishOp(&b, &ir.PublishOp{Topic: "order.completed", Payload: []ir.FieldArg{{Key: "id", Literal: "1"}}}, "  ")
	out := b.String()
	if !strings.Contains(out, "this.queue.publish('order.completed'") || !strings.Contains(out, "id: 1") {
		t.Errorf("publish op = %q", out)
	}
}

func TestRenderVerifyPasswordOp(t *testing.T) {
	var b strings.Builder
	renderVerifyPasswordOp(&b, nil, "  ", "this.prisma")
	if b.String() != "" {
		t.Errorf("nil verify")
	}
	b.Reset()
	renderVerifyPasswordOp(&b, &ir.VerifyPasswordOp{
		Model:        "User",
		EmailCol:     "Email",
		EmailExpr:    "request.email",
		HashCol:      "PasswordHash",
		PasswordExpr: "request.password",
		ResultVar:    "user",
		Message:      "invalid",
	}, "  ", "this.prisma")
	out := b.String()
	if !strings.Contains(out, "findUnique({ where: { email: body.email } })") {
		t.Errorf("verify lookup = %q", out)
	}
	if !strings.Contains(out, "bcrypt.compare(body.password") {
		t.Errorf("verify bcrypt = %q", out)
	}
}

func TestResolveNestJSExpr(t *testing.T) {
	if got := resolveNestJSExpr("request.email"); got != "body.email" {
		t.Errorf("request rewrite = %q", got)
	}
	if got := resolveNestJSExpr("user.id"); got != "user.id" {
		t.Errorf("passthrough = %q", got)
	}
}

func TestResolveDataKey(t *testing.T) {
	if got := resolveDataKey(ir.FieldArg{ColumnName: "user_id"}); got != "user_id" {
		t.Errorf("columnname = %q", got)
	}
	if got := resolveDataKey(ir.FieldArg{Key: "UserName"}); got != toSnake("UserName") {
		t.Errorf("key = %q", got)
	}
	if got := resolveDataKey(ir.FieldArg{Field: ".CourseId"}); got != toSnake("CourseId") {
		t.Errorf("field = %q", got)
	}
	if got := resolveDataKey(ir.FieldArg{}); got != "" {
		t.Errorf("empty = %q", got)
	}
}

func TestRenderPrismaDataWhere(t *testing.T) {
	if got := renderPrismaData(nil); got != "{ data: body }" {
		t.Errorf("empty data = %q", got)
	}
	got := renderPrismaData([]ir.FieldArg{{Key: "title", ColumnName: "title", Literal: "x", IsQuoted: true}})
	if !strings.Contains(got, "data: { title: 'x' }") {
		t.Errorf("data = %q", got)
	}
	// all keys empty → fallback
	if got := renderPrismaData([]ir.FieldArg{{}}); got != "{ data: body }" {
		t.Errorf("empty-key fallback = %q", got)
	}

	if got := renderPrismaWhere(nil); got != "{}" {
		t.Errorf("empty where = %q", got)
	}
	w := renderPrismaWhere([]ir.FieldArg{{Key: "id", ColumnName: "id", Literal: "1"}})
	if !strings.Contains(w, "where: { id: 1 }") {
		t.Errorf("where = %q", w)
	}
}

func TestRenderOneOpAndBody(t *testing.T) {
	var b strings.Builder
	ops := []ir.Op{
		{Kind: ir.OpEmpty, Empty: &ir.EmptyOp{VarName: "c", Message: "nf", StatusCode: 404}},
		{Kind: ir.OpExists, Exists: &ir.ExistsOp{VarName: "d", Message: "cf", StatusCode: 409}},
		{Kind: ir.OpCall, Call: &ir.CallOp{Function: "Run"}},
		{Kind: ir.OpEval, Eval: &ir.EvalOp{Function: "Check", Message: "bad", StatusCode: 400}},
		{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "t"}},
	}
	renderOpsBody(&b, ops, "  ", "this.prisma")
	out := b.String()
	for _, want := range []string{"if (!c)", "if (d)", "await run()", "await check()", "this.queue.publish('t'"} {
		if !strings.Contains(out, want) {
			t.Errorf("renderOpsBody missing %q in %q", want, out)
		}
	}
}

func TestRenderOneOp_RemainingKinds(t *testing.T) {
	// Dispatch coverage for the remaining renderOneOp branches.
	ops := []ir.Op{
		{Kind: ir.OpGet, Get: &ir.GetOp{VarName: "c", VarType: "Course", Model: "Course"}},
		{Kind: ir.OpPost, Post: &ir.PostOp{VarName: "c", Model: "Course"}},
		{Kind: ir.OpPut, Put: &ir.PutOp{Model: "Course", Args: []ir.FieldArg{{Key: "id", IsPK: true}}}},
		{Kind: ir.OpAuth, Auth: &ir.AuthOp{Action: "read", Resource: "course", Message: "denied", StatusCode: 403}},
		{Kind: ir.OpState, State: &ir.StateOp{Diagram: "d", Transition: "go", Message: "bad", StatusCode: 409}},
		{Kind: ir.OpVerifyPassword, VerifyPW: &ir.VerifyPasswordOp{Model: "User", EmailCol: "Email", EmailExpr: "request.email", HashCol: "PasswordHash", PasswordExpr: "request.password", ResultVar: "u", Message: "no"}},
		{Kind: ir.OpResponse, Response: &ir.ResponseOp{SingleVar: "c"}},
		{Kind: ir.OpDelete, Delete: &ir.DeleteOp{Model: "Course", Args: []ir.FieldArg{{Key: "id", IsPK: true}}}},
	}
	for _, op := range ops {
		var b strings.Builder
		// must not panic; each branch dispatches to its renderer
		renderOneOp(&b, op, "  ", "this.prisma")
	}
}
