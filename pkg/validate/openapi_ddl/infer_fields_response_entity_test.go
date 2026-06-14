//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what inferFieldsResponseEntity — @response Fields 의 base var 들을 단일 엔티티로 수렴, 컬렉션/미해결/혼합/리터럴은 ""

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestInferFieldsResponseEntity(t *testing.T) {
	tables := []ddl.Table{tbl("rules", "id", "name")}
	schemas := openapi3.Schemas{"Rule": &openapi3.SchemaRef{Value: schemaOf("id", "name")}}

	// converge to single model "rule" (type Rule)
	fn := &ssac.ServiceFunc{Name: "GetRule", Sequences: []ssac.Sequence{
		{Type: "get", Result: &ssac.Result{Var: "rule", Type: "Rule"}},
	}}
	idx := idxFor(tables, schemas, []ssac.ServiceFunc{*fn},
		map[string]string{"SSaC.var.GetRule.rule": "Rule"})
	fnIdx := idx.funcByName["GetRule"]
	if got := inferFieldsResponseEntity(idx, fnIdx, map[string]string{"id": "rule.ID", "name": "rule.Name", "k": "\"lit\""}); got != "Rule" {
		t.Errorf("converge → %q, want Rule", got)
	}

	// collection-bound base var → ""
	fnCol := &ssac.ServiceFunc{Name: "List", Sequences: []ssac.Sequence{
		{Type: "get", Result: &ssac.Result{Var: "rules", Type: "Rule", Wrapper: "Page"}},
	}}
	idxCol := idxFor(tables, schemas, []ssac.ServiceFunc{*fnCol},
		map[string]string{"SSaC.var.List.rules": "Rule"})
	if got := inferFieldsResponseEntity(idxCol, idxCol.funcByName["List"], map[string]string{"x": "rules.ID"}); got != "" {
		t.Errorf("collection-bound → %q, want \"\"", got)
	}

	// unresolvable var (no Ground type) → ""
	fnUnres := &ssac.ServiceFunc{Name: "Foo", Sequences: nil}
	idxUnres := idxFor(tables, schemas, []ssac.ServiceFunc{*fnUnres}, nil)
	if got := inferFieldsResponseEntity(idxUnres, idxUnres.funcByName["Foo"], map[string]string{"x": "ghost.ID"}); got != "" {
		t.Errorf("unresolvable → %q, want \"\"", got)
	}

	// mixed models → ""
	fnMix := &ssac.ServiceFunc{Name: "Mix", Sequences: nil}
	idxMix := idxFor(tables, schemas, []ssac.ServiceFunc{*fnMix},
		map[string]string{"SSaC.var.Mix.a": "Rule", "SSaC.var.Mix.b": "Note"})
	if got := inferFieldsResponseEntity(idxMix, idxMix.funcByName["Mix"], map[string]string{"x": "a.ID", "y": "b.ID"}); got != "" {
		t.Errorf("mixed models → %q, want \"\"", got)
	}

	// only literals → "" (converged stays empty)
	idxLit := idxFor(tables, schemas, []ssac.ServiceFunc{{Name: "Lit"}}, nil)
	if got := inferFieldsResponseEntity(idxLit, idxLit.funcByName["Lit"], map[string]string{"x": "42", "y": "\"s\""}); got != "" {
		t.Errorf("literals → %q, want \"\"", got)
	}
}
