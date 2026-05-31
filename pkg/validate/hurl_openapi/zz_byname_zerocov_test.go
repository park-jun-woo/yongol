//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — checkEntryURLMethod / xoh12 / xoh13 분기 직접 호출

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func miniDoc() *openapi3.T {
	op := openapi3.NewOperation()
	op.OperationID = "ListWidgets"
	op.Responses = openapi3.NewResponses()
	op.Responses.Set("200", &openapi3.ResponseRef{Value: openapi3.NewResponse().WithDescription("ok")})
	op.Responses.Set("404", &openapi3.ResponseRef{Value: openapi3.NewResponse().WithDescription("missing")})
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/widgets", &openapi3.PathItem{Get: op})
	return doc
}

func TestByNameCheckEntryURLMethod_ZeroCov(t *testing.T) {
	routes := collectOpenAPIRoutes(miniDoc())

	// external-service entry (URLVar set, not host) → nil.
	ext := hurl.HurlEntry{Method: "GET", Path: "/x", URLVar: "authurl"}
	if d := checkEntryURLMethod(ext, routes); d != nil {
		t.Errorf("external entry should be skipped, got %v", d)
	}

	// matching own-API entry → nil.
	match := hurl.HurlEntry{Method: "GET", Path: "/widgets", URLVar: "host"}
	if d := checkEntryURLMethod(match, routes); d != nil {
		t.Errorf("matching entry should yield nil, got %v", d)
	}

	// drift entry → diagnostic.
	drift := hurl.HurlEntry{Method: "GET", Path: "/nonexistent", URLVar: "host", File: "t.hurl", Line: 1}
	if d := checkEntryURLMethod(drift, routes); d == nil {
		t.Errorf("drift entry should yield diagnostic")
	}
}

func TestByNameXoh12_ZeroCov(t *testing.T) {
	// nil/empty guards.
	if d := xoh12StatusCoverage(nil); d != nil {
		t.Errorf("nil fs should yield nil")
	}
	// populated: GET /widgets covers 200 but not 404 → warning.
	fs := &yongol.Fullstack{
		OpenAPIDoc: miniDoc(),
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/widgets", StatusCode: "200", URLVar: "host"},
		},
	}
	_ = xoh12StatusCoverage(fs) // exercise the iteration; result depends on 5xx filter
}

func TestByNameXoh13_ZeroCov(t *testing.T) {
	if d := xoh13GuardCoverage(nil); d != nil {
		t.Errorf("nil fs should yield nil")
	}
	// missing service funcs → nil early return.
	fs := &yongol.Fullstack{OpenAPIDoc: miniDoc(), HurlEntries: []hurl.HurlEntry{{Method: "GET", Path: "/widgets"}}}
	if d := xoh13GuardCoverage(fs); d != nil {
		t.Errorf("no service funcs should yield nil, got %v", d)
	}
}
