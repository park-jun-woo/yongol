//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCheckActionClass — ActionBlock 자체/Fields/Children class 진단 누적 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCheckActionClass(t *testing.T) {
	// Empty action block → no diagnostics.
	if got := checkActionClass(stml.ActionBlock{OperationID: "Create"}, "p.stml"); got != nil {
		t.Fatalf("empty block: expected nil, got %+v", got)
	}

	ab := stml.ActionBlock{
		OperationID: "CreateItem",
		ClassName:   "bg-blue",
		Fields: []stml.FieldBind{
			{Name: "title", ClassName: "text-lg"},
			{Name: "noclass"},
		},
		Children: []stml.ChildNode{
			{Kind: "bind", Bind: &stml.FieldBind{Name: "child", ClassName: "p-2"}},
		},
	}
	diags := checkActionClass(ab, "p.stml")
	// ActionBlock class (1) + title field (1) + child bind (1) = 3.
	if len(diags) != 3 {
		t.Fatalf("expected 3 diagnostics, got %d: %+v", len(diags), diags)
	}
}
