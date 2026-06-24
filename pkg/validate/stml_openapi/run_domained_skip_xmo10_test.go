//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestRunDomained_SkipsXMO10ButFiresXMO11 — 도메인 모드: XMO-10 미발동 + 전체 페이지 0개 시 XMO-11 발동

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRunDomained_SkipsXMO10ButFiresXMO11(t *testing.T) {
	// Frontend ON, two domains, but zero STML pages anywhere. A normal
	// (non-no-front) op is left unconsumed: single-site would emit XMO-10, but
	// domain mode must skip it. XMO-11 (no pages) must still fire exactly once.
	docs := map[string]*openapi3.T{
		"public": makeDoc(map[string]*openapi3.PathItem{
			"/users": {Get: &openapi3.Operation{OperationID: "ListUsers"}},
		}),
		"admin": makeDoc(map[string]*openapi3.PathItem{
			"/admin/users": {Get: &openapi3.Operation{OperationID: "ListAdminUsers"}},
		}),
	}
	fs := domainedFS(docs, nil)

	diags := Run(fs)
	if n := countDiag(diags, "[XMO-10]"); n != 0 {
		t.Errorf("XMO-10 must be skipped in domain mode, got %d", n)
	}
	if n := countDiag(diags, "[XMO-11]"); n != 1 {
		t.Fatalf("expected 1 XMO-11 in domain mode, got %d (%v)", n, diags)
	}
}
