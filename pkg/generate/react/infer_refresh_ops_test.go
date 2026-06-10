//ff:func feature=gen-react type=test control=sequence
//ff:what inferRefreshOps — token+refresh 보유 op 수집, 경로파라미터·capture선언·필드부족·opID없음 제외 + 정렬

package react

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestInferRefreshOps(t *testing.T) {
	tokenResp := []string{"access_token", "refresh_token"}
	fa := &manifest.FrontendAuth{TokenField: "access_token", RefreshField: "refresh_token"}

	doc := &openapi3.T{Paths: openapi3.NewPaths(
		// qualifies, but capture-declared (excluded via captured map below)
		openapi3.WithPath("/auth/login", &openapi3.PathItem{Post: buildTokenOp("Login", tokenResp, []string{"email"})}),
		// qualifies -> candidate
		openapi3.WithPath("/auth/renew", &openapi3.PathItem{Post: buildTokenOp("Renew", tokenResp, []string{"refresh_token"})}),
		// qualifies -> candidate (proves sort: Refresh < Renew by opID)
		openapi3.WithPath("/auth/refresh", &openapi3.PathItem{Post: buildTokenOp("Refresh", tokenResp, []string{"refresh_token"})}),
		// path parameter -> skipped wholesale
		openapi3.WithPath("/auth/refresh/{id}", &openapi3.PathItem{Post: buildTokenOp("RefreshById", tokenResp, nil)}),
		// missing refresh field -> excluded
		openapi3.WithPath("/auth/access", &openapi3.PathItem{Post: buildTokenOp("AccessOnly", []string{"access_token"}, nil)}),
		// empty operationId -> excluded
		openapi3.WithPath("/auth/blank", &openapi3.PathItem{Post: buildTokenOp("", tokenResp, nil)}),
	)}

	captured := map[string]bool{"Login": true}
	cands := inferRefreshOps(doc, fa, captured)

	if len(cands) != 2 {
		t.Fatalf("candidates = %+v, want exactly Refresh and Renew", cands)
	}
	// sorted by opID: Refresh before Renew
	if cands[0].opID != "Refresh" || cands[1].opID != "Renew" {
		t.Errorf("order = [%s, %s], want [Refresh, Renew]", cands[0].opID, cands[1].opID)
	}
	if cands[0].bodyKey != "refresh_token" {
		t.Errorf("Refresh bodyKey = %q, want refresh_token", cands[0].bodyKey)
	}

	// no qualifying ops -> nil
	noTok := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items", &openapi3.PathItem{Get: buildTokenOp("ListItems", []string{"items"}, nil)}),
	)}
	if got := inferRefreshOps(noTok, fa, nil); got != nil {
		t.Errorf("no candidates = %+v, want nil", got)
	}
}
