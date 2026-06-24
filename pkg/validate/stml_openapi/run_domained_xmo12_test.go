//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestRunDomained_FiresXMO12AcrossDomains — 도메인 모드: 전체 도메인에 걸쳐 XMO-12 발동, XMO-10/11 미발동

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRunDomained_FiresXMO12AcrossDomains(t *testing.T) {
	// public domain declares a no-front op that its own page consumes → XMO-12.
	// A second domain with a consumed normal op keeps pages non-empty so XMO-11
	// stays silent, proving XMO-12 fires independently across domains.
	docs := map[string]*openapi3.T{
		"public": makeDoc(map[string]*openapi3.PathItem{
			"/users": {Get: &openapi3.Operation{OperationID: "ListUsers", Tags: []string{noFrontTag}}},
		}),
		"admin": makeDoc(map[string]*openapi3.PathItem{
			"/admin/users": {Get: &openapi3.Operation{OperationID: "ListAdminUsers"}},
		}),
	}
	pages := map[string][]stml.PageSpec{
		"public": {{FileName: "public/users.html", Fetches: []stml.FetchBlock{{OperationID: "ListUsers"}}}},
		"admin":  {{FileName: "admin/users.html", Fetches: []stml.FetchBlock{{OperationID: "ListAdminUsers"}}}},
	}
	fs := domainedFS(docs, pages)
	// A manifest domain with no pre-parsed doc exercises the nil-doc skip in
	// the XMO-12 emit loop without affecting the result.
	fs.Manifest.Domains["ghost"] = manifest.DomainConfig{}

	diags := Run(fs)
	if n := countDiag(diags, "[XMO-12]"); n != 1 {
		t.Fatalf("expected 1 XMO-12 across domains, got %d (%v)", n, diags)
	}
	if n := countDiag(diags, "[XMO-11]"); n != 0 {
		t.Errorf("XMO-11 should not fire when pages exist, got %d", n)
	}
	if n := countDiag(diags, "[XMO-10]"); n != 0 {
		t.Errorf("XMO-10 must be skipped in domain mode, got %d", n)
	}
}
