//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what resolveSSaCEntity — 전략 A(Target var)/B-1(Fields) 로 SSaC 응답 엔티티 해석, 컬렉션/빈타입/미해결은 ""

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestResolveSSaCEntity(t *testing.T) {
	tables := []ddl.Table{tbl("rules", "id", "name")}
	schemas := openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name")}}

	// Strategy A: Target single var → component
	fnA := ssac.ServiceFunc{Name: "GetRule", Sequences: []ssac.Sequence{
		{Type: "get", Result: &ssac.Result{Var: "rule", Type: "Rule"}},
		{Type: "response", Target: "rule"},
	}}
	idxA := idxFor(tables, schemas, []ssac.ServiceFunc{fnA},
		map[string]string{"SSaC.var.GetRule.rule": "Rule"})
	fnAp := idxA.funcByName["GetRule"]
	if got := resolveSSaCEntity(idxA, fnAp, responseSeqOf(fnAp)); got != "Rule" {
		t.Errorf("strategy A → %q, want Rule", got)
	}

	// Target bound to collection → ""
	fnC := ssac.ServiceFunc{Name: "ListRules", Sequences: []ssac.Sequence{
		{Type: "get", Result: &ssac.Result{Var: "rules", Type: "Rule", Wrapper: "Page"}},
		{Type: "response", Target: "rules"},
	}}
	idxC := idxFor(tables, schemas, []ssac.ServiceFunc{fnC},
		map[string]string{"SSaC.var.ListRules.rules": "Rule"})
	fnCp := idxC.funcByName["ListRules"]
	if got := resolveSSaCEntity(idxC, fnCp, responseSeqOf(fnCp)); got != "" {
		t.Errorf("collection Target → %q, want \"\"", got)
	}

	// Target with empty Ground type → ""
	fnE := ssac.ServiceFunc{Name: "Ghost", Sequences: []ssac.Sequence{
		{Type: "response", Target: "rule"},
	}}
	idxE := idxFor(tables, schemas, []ssac.ServiceFunc{fnE}, nil)
	fnEp := idxE.funcByName["Ghost"]
	if got := resolveSSaCEntity(idxE, fnEp, responseSeqOf(fnEp)); got != "" {
		t.Errorf("empty type Target → %q, want \"\"", got)
	}

	// Target with collection raw type → ""
	idxCol := idxFor(tables, schemas, []ssac.ServiceFunc{fnE},
		map[string]string{"SSaC.var.Ghost.rule": "[]Rule"})
	fnColp := idxCol.funcByName["Ghost"]
	if got := resolveSSaCEntity(idxCol, fnColp, responseSeqOf(fnColp)); got != "" {
		t.Errorf("collection raw type Target → %q, want \"\"", got)
	}

	// Strategy B-1: Fields → component
	fnB := ssac.ServiceFunc{Name: "MakeRule", Sequences: []ssac.Sequence{
		{Type: "post", Result: &ssac.Result{Var: "rule", Type: "Rule"}},
		{Type: "response", Fields: map[string]string{"id": "rule.ID", "name": "rule.Name"}},
	}}
	idxB := idxFor(tables, schemas, []ssac.ServiceFunc{fnB},
		map[string]string{"SSaC.var.MakeRule.rule": "Rule"})
	fnBp := idxB.funcByName["MakeRule"]
	if got := resolveSSaCEntity(idxB, fnBp, responseSeqOf(fnBp)); got != "Rule" {
		t.Errorf("strategy B-1 → %q, want Rule", got)
	}

	// Neither Target nor Fields → ""
	fnN := ssac.ServiceFunc{Name: "Nada", Sequences: []ssac.Sequence{{Type: "response"}}}
	idxN := idxFor(tables, schemas, []ssac.ServiceFunc{fnN}, nil)
	fnNp := idxN.funcByName["Nada"]
	if got := resolveSSaCEntity(idxN, fnNp, responseSeqOf(fnNp)); got != "" {
		t.Errorf("empty response → %q, want \"\"", got)
	}
}
