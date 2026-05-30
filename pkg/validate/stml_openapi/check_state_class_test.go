//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCheckStateClass — StateBind 자체/Children class 검사

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCheckStateClass(t *testing.T) {
	// Minimal → no diagnostics.
	if got := checkStateClass(stml.StateBind{Condition: "empty"}, "p.stml"); got != nil {
		t.Fatalf("empty state: expected nil, got %+v", got)
	}

	sb := stml.StateBind{
		Condition: "empty",
		ClassName: "muted",
		Children: []stml.ChildNode{
			{Kind: "bind", Bind: &stml.FieldBind{Name: "c", ClassName: "p-1"}},
		},
	}
	diags := checkStateClass(sb, "p.stml")
	// state class(1) + child(1) = 2.
	if len(diags) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d: %+v", len(diags), diags)
	}
}
