//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what XOH-01 positive/negative 테스트 — path + method 매칭

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh01_Negative_PathMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/login": {"POST": {OperationID: "Login"}},
		}),
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/gigs", File: "t.hurl", Line: 1},
		},
	}
	diags := xoh01URLMethod(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XOH-01]") || !strings.Contains(diags[0].Message, "path not declared") {
		t.Fatalf("unexpected msg: %q", diags[0].Message)
	}
}

func TestXoh01_Negative_MethodMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/gigs": {"GET": {OperationID: "ListGigs"}},
		}),
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/gigs", File: "t.hurl", Line: 1},
		},
	}
	diags := xoh01URLMethod(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "method not declared") {
		t.Fatalf("unexpected msg: %q", diags[0].Message)
	}
}

func TestXoh01_Positive_ExactMatch(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/gigs/{id}": {"GET": {OperationID: "GetGig"}},
		}),
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/gigs/{{id}}", File: "t.hurl", Line: 1},
		},
	}
	if diags := xoh01URLMethod(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d: %+v", len(diags), diags)
	}
}

// newDoc builds a minimal *openapi3.T whose paths match the given
// table. Operations are attached by HTTP method verb. Used by every
// hurl_openapi rule test to keep fixtures inline and readable.
func newDoc(paths map[string]map[string]*openapi3.Operation) *openapi3.T {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	for p, methods := range paths {
		pi := &openapi3.PathItem{}
		for m, op := range methods {
			if op.Responses == nil {
				op.Responses = openapi3.NewResponses()
			}
			setOp(pi, m, op)
		}
		doc.Paths.Set(p, pi)
	}
	return doc
}

// setOp attaches op to pi under the given HTTP method. openapi3's
// PathItem exposes dedicated fields per verb; centralising the switch
// makes fixture code small.
func setOp(pi *openapi3.PathItem, method string, op *openapi3.Operation) {
	switch strings.ToUpper(method) {
	case "GET":
		pi.Get = op
	case "POST":
		pi.Post = op
	case "PUT":
		pi.Put = op
	case "DELETE":
		pi.Delete = op
	case "PATCH":
		pi.Patch = op
	}
}
