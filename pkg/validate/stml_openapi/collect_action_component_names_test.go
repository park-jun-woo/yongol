//ff:func feature=validate type=test control=iteration dimension=2 topic=stml-openapi
//ff:what TestCollectActionComponentNames — Fields(data-component:)와 children 재귀 두 경로 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectActionComponentNames(t *testing.T) {
	a := stml.ActionBlock{
		OperationID: "CreateItem",
		Fields: []stml.FieldBind{
			{Name: "StartAt", Tag: "data-component:FieldComp"},
			{Name: "Plain", Tag: "data-field"},
		},
		Children: []stml.ChildNode{
			{Kind: "component", Component: &stml.ComponentRef{Name: "ChildComp"}},
		},
	}

	out := map[string]struct{}{}
	collectActionComponentNames(a, out)

	want := []string{"FieldComp", "ChildComp"}
	for _, name := range want {
		if _, ok := out[name]; !ok {
			t.Errorf("missing component name %q", name)
		}
	}
	if _, ok := out["Plain"]; ok {
		t.Errorf("non-component tag should not yield a name, got %+v", out)
	}
	if len(out) != len(want) {
		t.Errorf("expected %d names, got %d: %+v", len(want), len(out), out)
	}
}
