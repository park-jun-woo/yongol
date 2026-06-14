//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what subtestTestCanonicalResponseReprNonEntitySkipped — DDL 테이블 없는 model(token) 수렴은 그룹 미형성 → 무진단

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func subtestTestCanonicalResponseReprNonEntitySkipped(t *testing.T) {
	ruleTables := []ddl.Table{tbl("rules", "id", "name", "created_at")}
	ruleSchemas := openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name", "created_at")}}
	// Both ops converge to model "token" which has no DDL table → no group,
	// no diagnostic even though representations differ.
	fs := buildCanonicalFS(
		ruleTables, ruleSchemas,
		[]canonOp{
			{"POST", "Login", jsonResp("201", inlineRef("access"))},
			{"POST", "Refresh", jsonResp("201", inlineRef("access", "refresh"))},
		},
		[]ssac.ServiceFunc{
			{Name: "Login", Sequences: []ssac.Sequence{
				{Type: "call", Result: &ssac.Result{Var: "token", Type: "Token"}},
				{Type: "response", Fields: map[string]string{"access": "token.Access"}},
			}},
			{Name: "Refresh", Sequences: []ssac.Sequence{
				{Type: "call", Result: &ssac.Result{Var: "token", Type: "Token"}},
				{Type: "response", Fields: map[string]string{"access": "token.Access", "refresh": "token.Refresh"}},
			}},
		},
		map[string]string{
			"SSaC.var.Login.token":   "Token",
			"SSaC.var.Refresh.token": "Token",
		},
	)
	diags := canonicalResponseRepr(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics (non-entity), got %d: %+v", len(diags), diags)
	}
}
