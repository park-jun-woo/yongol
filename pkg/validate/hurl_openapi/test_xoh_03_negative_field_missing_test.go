//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh03_Negative_FieldMissing — body 필드 오타 → [XOH-03]

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh03_Negative_FieldMissing(t *testing.T) {
	op := withRequestBody("Register", map[string]*openapi3.Schema{
		"email":    {Type: &openapi3.Types{"string"}},
		"password": {Type: &openapi3.Types{"string"}},
	})
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/register": {"POST": op},
		}),
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/register", StatusCode: "201",
				File: "t.hurl", Line: 1,
				BodyFields: []string{"emale"}, // typo
			},
		},
	}
	diags := xoh03RequestFieldInSchema(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "emale") || !strings.Contains(diags[0].Advice, "email") {
		t.Fatalf("advice/msg missing field hints: %+v", diags[0])
	}
}
