//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what subtestTestCanonicalResponseReprXdo11SilentAllShareRef — 모든 응답이 동일 $ref 공유 시 무진단

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func subtestTestCanonicalResponseReprXdo11SilentAllShareRef(t *testing.T) {
	ruleTables := []ddl.Table{tbl("rules", "id", "name", "created_at")}
	ruleSchemas := openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name", "created_at")}}
	fs := buildCanonicalFS(
		ruleTables, ruleSchemas,
		[]canonOp{
			{"GET", "GetRule", jsonResp("200", compRef("Rule", "id", "name", "created_at"))},
			{"PUT", "UpdateRule", jsonResp("200", compRef("Rule", "id", "name", "created_at"))},
		},
		nil, nil,
	)
	diags := canonicalResponseRepr(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
