//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCollectStateComponentNames — StateBind children 재귀 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectStateComponentNames(t *testing.T) {
	s := stml.StateBind{
		Children: []stml.ChildNode{
			{Kind: "component", Component: &stml.ComponentRef{Name: "StateChildComp"}},
		},
	}

	out := map[string]struct{}{}
	collectStateComponentNames(s, out)

	if _, ok := out["StateChildComp"]; !ok {
		t.Errorf("missing component name %q, got %+v", "StateChildComp", out)
	}
	if len(out) != 1 {
		t.Errorf("expected 1 name, got %d: %+v", len(out), out)
	}
}
