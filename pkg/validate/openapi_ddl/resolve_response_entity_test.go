//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what resolveResponseEntity — @response 있으면 SSaC 경로, 없거나 fn nil 이면 fallback 으로 엔티티 해석

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestResolveResponseEntity(t *testing.T) {
	tables := []ddl.Table{tbl("rules", "id", "name")}
	schemas := openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name")}}

	// fn with @response seq → SSaC path
	fn := ssac.ServiceFunc{Name: "GetRule", Sequences: []ssac.Sequence{
		{Type: "get", Result: &ssac.Result{Var: "rule", Type: "Rule"}},
		{Type: "response", Target: "rule"},
	}}
	idx := idxFor(tables, schemas, []ssac.ServiceFunc{fn},
		map[string]string{"SSaC.var.GetRule.rule": "Rule"})
	if got := resolveResponseEntity(idx, idx.funcByName["GetRule"], inlineRef("id")); got != "Rule" {
		t.Errorf("SSaC path → %q, want Rule", got)
	}

	// fn != nil but no @response seq → fallback used
	fnNoResp := ssac.ServiceFunc{Name: "PlainGet", Sequences: []ssac.Sequence{{Type: "get"}}}
	idx2 := idxFor(tables, schemas, []ssac.ServiceFunc{fnNoResp}, nil)
	if got := resolveResponseEntity(idx2, idx2.funcByName["PlainGet"], compRef("Rule", "id")); got != "Rule" {
		t.Errorf("no-response fn → fallback %q, want Rule", got)
	}

	// fn == nil → fallback
	if got := resolveResponseEntity(idx, nil, compRef("Rule", "id")); got != "Rule" {
		t.Errorf("nil fn → fallback %q, want Rule", got)
	}
}
