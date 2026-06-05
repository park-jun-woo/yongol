//ff:func feature=validate type=test control=iteration dimension=5 topic=stml-openapi
//ff:what TestCollectChildComponentNames — Component/Fetch/Each/Action/State 다섯 분기별 이름 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectChildComponentNames(t *testing.T) {
	tests := []struct {
		name  string
		child stml.ChildNode
		want  string
	}{
		{"component", stml.ChildNode{Kind: "component", Component: &stml.ComponentRef{Name: "DirectComp"}}, "DirectComp"},
		{"fetch", stml.ChildNode{Kind: "fetch", Fetch: &stml.FetchBlock{Components: []stml.ComponentRef{{Name: "FetchComp"}}}}, "FetchComp"},
		{"each", stml.ChildNode{Kind: "each", Each: &stml.EachBlock{Components: []stml.ComponentRef{{Name: "EachComp"}}}}, "EachComp"},
		{"action", stml.ChildNode{Kind: "action", Action: &stml.ActionBlock{Fields: []stml.FieldBind{{Tag: "data-component:ActionComp"}}}}, "ActionComp"},
		{"state", stml.ChildNode{Kind: "state", State: &stml.StateBind{Children: []stml.ChildNode{{Kind: "component", Component: &stml.ComponentRef{Name: "StateComp"}}}}}, "StateComp"},
	}
	for _, tt := range tests {
		out := map[string]struct{}{}
		collectChildComponentNames(tt.child, out)
		if _, ok := out[tt.want]; !ok || len(out) != 1 {
			t.Errorf("%s: want {%q}, got %+v", tt.name, tt.want, out)
		}
	}
}
