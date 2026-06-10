//ff:func feature=gen-react type=test control=sequence
//ff:what resolveRefreshPlan — 선언 우선 / 구조 추론 / capture 제외 / 모호(2+) ERROR / 다운그레이드 검증

package react

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveRefreshPlan(t *testing.T) {
	newFS := func(fa *manifest.FrontendAuth, doc *openapi3.T, pages []stml.PageSpec) *yongol.Fullstack {
		return &yongol.Fullstack{
			Manifest: &manifest.ProjectConfig{
				Backend:  manifest.Backend{Auth: &manifest.Auth{Mode: "bearer"}},
				Frontend: manifest.Frontend{Auth: fa},
			},
			OpenAPIDoc: doc,
			STMLPages:  pages,
		}
	}
	fa := &manifest.FrontendAuth{TokenField: "access_token", RefreshField: "refresh_token"}
	loginCapture := []stml.PageSpec{{Actions: []stml.ActionBlock{{
		OperationID: "Login",
		Captures:    []stml.CaptureBind{{RespField: "access_token", Sink: "auth.token"}},
	}}}}
	tokenResp := []string{"access_token", "refresh_token"}

	// refresh_field undeclared -> nil (explicit downgrade)
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/refresh", &openapi3.PathItem{Post: buildTokenOp("Refresh", tokenResp, []string{"refresh_token"})}),
	)}
	if rp, err := resolveRefreshPlan(newFS(&manifest.FrontendAuth{TokenField: "access_token"}, doc, nil)); rp != nil || err != nil {
		t.Errorf("no refresh_field = (%v,%v), want (nil,nil)", rp, err)
	}

	// declared refresh_op wins, body key matched from requestBody
	declared := &manifest.FrontendAuth{TokenField: "access_token", RefreshField: "refresh_token", RefreshOp: "Refresh"}
	rp, err := resolveRefreshPlan(newFS(declared, doc, nil))
	if err != nil || rp == nil {
		t.Fatalf("declared = (%v,%v), want plan", rp, err)
	}
	if rp.opID != "Refresh" || rp.method != "POST" || rp.path != "/auth/refresh" || rp.bodyKey != "refresh_token" {
		t.Errorf("declared plan = %+v", rp)
	}

	// declared refresh_op not present in the doc -> nil (XON-60 reports the
	// missing operationId; emitter degrades to the no-refresh client)
	missingOp := &manifest.FrontendAuth{TokenField: "access_token", RefreshField: "refresh_token", RefreshOp: "Nonexistent"}
	if rp, err := resolveRefreshPlan(newFS(missingOp, doc, nil)); rp != nil || err != nil {
		t.Errorf("missing declared op = (%v,%v), want (nil,nil)", rp, err)
	}

	// inference: Login (capture-declared) excluded -> unique candidate Refresh
	doc2 := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/login", &openapi3.PathItem{Post: buildTokenOp("Login", tokenResp, []string{"email", "password"})}),
		openapi3.WithPath("/auth/refresh", &openapi3.PathItem{Post: buildTokenOp("Refresh", tokenResp, []string{"refresh_token"})}),
	)}
	rp, err = resolveRefreshPlan(newFS(fa, doc2, loginCapture))
	if err != nil || rp == nil || rp.opID != "Refresh" {
		t.Fatalf("inferred = (%+v,%v), want Refresh", rp, err)
	}

	// inference ambiguity: Login has no data-capture -> 2 candidates -> ERROR
	_, err = resolveRefreshPlan(newFS(fa, doc2, nil))
	if err == nil || !strings.Contains(err.Error(), "refresh_op") {
		t.Errorf("ambiguous = %v, want refresh_op advice error", err)
	}

	// zero candidates -> nil (no refresh endpoint in this project)
	noTok := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items", &openapi3.PathItem{Get: buildTokenOp("ListItems", []string{"items"}, nil)}),
	)}
	if rp, err := resolveRefreshPlan(newFS(fa, noTok, nil)); rp != nil || err != nil {
		t.Errorf("zero candidates = (%v,%v), want (nil,nil)", rp, err)
	}

	// declared op with path params -> generate ERROR
	paramDoc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/refresh/{id}", &openapi3.PathItem{Post: buildTokenOp("Refresh", tokenResp, nil)}),
	)}
	if _, err := resolveRefreshPlan(newFS(declared, paramDoc, nil)); err == nil {
		t.Errorf("declared op with path params: want error, got nil")
	}
}
