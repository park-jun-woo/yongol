//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what subtestTestCanonicalResponseReprXdo11FiresFlatInlineVsRef — flat-inline GET vs $ref Update 발산 → XDO-11 1건 (BUG-131)

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func subtestTestCanonicalResponseReprXdo11FiresFlatInlineVsRef(t *testing.T) {
	ruleTables := []ddl.Table{tbl("rules", "id", "name", "created_at")}
	ruleSchemas := openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name", "created_at")}}
	fs := buildCanonicalFS(
		ruleTables, ruleSchemas,
		[]canonOp{
			{"GET", "GetRule", jsonResp("200", inlineRef("id", "name"))},
			{"PUT", "UpdateRule", jsonResp("200", compRef("Rule", "id", "name", "created_at"))},
		},
		[]ssac.ServiceFunc{{
			Name: "GetRule",
			Sequences: []ssac.Sequence{
				{Type: "get", Result: &ssac.Result{Var: "rule", Type: "Rule"}},
				{Type: "response", Fields: map[string]string{"id": "rule.ID", "name": "rule.Name"}},
			},
		}},
		map[string]string{"SSaC.var.GetRule.rule": "Rule"},
	)
	diags := canonicalResponseRepr(fs)
	if got := countLevel(diags, "[XDO-11]"); got != 1 {
		t.Fatalf("expected 1 XDO-11, got %d: %+v", got, diags)
	}
}
