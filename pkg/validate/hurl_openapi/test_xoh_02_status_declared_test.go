//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what XOH-02 positive/negative 테스트 — response status 선언

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh02_Negative_StatusMissing(t *testing.T) {
	op := &openapi3.Operation{
		OperationID: "Login",
		Responses:   openapi3.NewResponses(),
	}
	op.Responses.Set("200", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
	op.Responses.Set("401", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/login": {"POST": op},
		}),
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/auth/login", StatusCode: "500", File: "t.hurl", Line: 1},
		},
	}
	diags := xoh02StatusDeclared(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XOH-02]") {
		t.Fatalf("unexpected msg: %q", diags[0].Message)
	}
}

func TestXoh02_Positive_StatusDeclared(t *testing.T) {
	op := &openapi3.Operation{
		OperationID: "Login",
		Responses:   openapi3.NewResponses(),
	}
	op.Responses.Set("200", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/login": {"POST": op},
		}),
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/auth/login", StatusCode: "200", File: "t.hurl", Line: 1},
		},
	}
	if diags := xoh02StatusDeclared(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d: %+v", len(diags), diags)
	}
}
