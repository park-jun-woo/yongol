//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what subtestTestCanonicalResponseReprXdo12FiresConsistentInline — 일관되지만 inline(non-SSaC 컬럼매칭) → op 당 XDO-12

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func subtestTestCanonicalResponseReprXdo12FiresConsistentInline(t *testing.T) {
	fs := buildCanonicalFS(
		[]ddl.Table{tbl("users", "id", "name")},
		openapi3.Schemas{"User": &openapi3.SchemaRef{Value: schemaOf("id", "name")}},
		[]canonOp{
			{"GET", "GetUser", jsonResp("200", inlineRef("id", "name"))},
			{"POST", "CreateUser", jsonResp("201", inlineRef("id", "name"))},
		},
		nil, nil,
	)
	diags := canonicalResponseRepr(fs)
	if got := countLevel(diags, "[XDO-11]"); got != 0 {
		t.Fatalf("expected 0 XDO-11, got %d: %+v", got, diags)
	}
	if got := countLevel(diags, "[XDO-12]"); got != 2 {
		t.Fatalf("expected 2 XDO-12, got %d: %+v", got, diags)
	}
}
