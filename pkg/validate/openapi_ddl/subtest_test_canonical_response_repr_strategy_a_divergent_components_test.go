//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what subtestTestCanonicalResponseReprStrategyADivergentComponents — 전략 A shorthand @response var, 발산 component → XDO-11

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func subtestTestCanonicalResponseReprStrategyADivergentComponents(t *testing.T) {
	ruleTables := []ddl.Table{tbl("rules", "id", "name", "created_at")}
	schemas := openapi3.Schemas{
		"Rule":       &openapi3.SchemaRef{Value: schemaOf("id", "name", "created_at")},
		"RuleDetail": &openapi3.SchemaRef{Value: schemaOf("id", "name")},
	}
	fs := buildCanonicalFS(
		ruleTables, schemas,
		[]canonOp{
			{"GET", "GetRule", jsonResp("200", compRef("RuleDetail", "id", "name"))},
			{"PUT", "UpdateRule", jsonResp("200", compRef("Rule", "id", "name", "created_at"))},
		},
		[]ssac.ServiceFunc{
			{Name: "GetRule", Sequences: []ssac.Sequence{
				{Type: "get", Result: &ssac.Result{Var: "rule", Type: "Rule"}},
				{Type: "response", Target: "rule"},
			}},
			{Name: "UpdateRule", Sequences: []ssac.Sequence{
				{Type: "post", Result: &ssac.Result{Var: "rule", Type: "Rule"}},
				{Type: "response", Target: "rule"},
			}},
		},
		map[string]string{
			"SSaC.var.GetRule.rule":    "Rule",
			"SSaC.var.UpdateRule.rule": "Rule",
		},
	)
	diags := canonicalResponseRepr(fs)
	if got := countLevel(diags, "[XDO-11]"); got != 1 {
		t.Fatalf("expected 1 XDO-11, got %d: %+v", got, diags)
	}
}
