//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what subtestTestCanonicalResponseReprPaginatedListExcluded — 페이지네이션 list 응답은 그룹화 제외 → 무진단

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func subtestTestCanonicalResponseReprPaginatedListExcluded(t *testing.T) {
	ruleTables := []ddl.Table{tbl("rules", "id", "name", "created_at")}
	ruleSchemas := openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name", "created_at")}}
	fs := buildCanonicalFS(
		ruleTables, ruleSchemas,
		[]canonOp{
			{"GET", "GetRule", jsonResp("200", compRef("Rule", "id", "name", "created_at"))},
			{"GET", "ListRules", jsonResp("200", inlineRef("items", "total"))},
		},
		[]ssac.ServiceFunc{
			{Name: "GetRule", Sequences: []ssac.Sequence{
				{Type: "get", Result: &ssac.Result{Var: "rule", Type: "Rule"}},
				{Type: "response", Target: "rule"},
			}},
			{Name: "ListRules", Sequences: []ssac.Sequence{
				{Type: "get", Result: &ssac.Result{Var: "rules", Type: "Rule", Wrapper: "Page"}},
				{Type: "response", Target: "rules"},
			}},
		},
		map[string]string{
			"SSaC.var.GetRule.rule":    "Rule",
			"SSaC.var.ListRules.rules": "Rule",
		},
	)
	diags := canonicalResponseRepr(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics (list excluded; single GET), got %d: %+v", len(diags), diags)
	}
}
