//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what childStateConditions — 단일 ChildNode에서 data-state 조건 수집·재귀 (state/fetch/each/default) 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestChildStateConditions(t *testing.T) {
	tests := []struct {
		name string
		node stml.ChildNode
		want []string
	}{
		{
			name: "state with nested state child",
			node: stml.ChildNode{
				Kind: "state",
				State: &stml.StateBind{
					Condition: "a.b=x",
					Children: []stml.ChildNode{
						{Kind: "state", State: &stml.StateBind{Condition: "c.d=y"}},
					},
				},
			},
			want: []string{"a.b=x", "c.d=y"},
		},
		{
			name: "fetch recurses into children",
			node: stml.ChildNode{
				Kind: "fetch",
				Fetch: &stml.FetchBlock{
					Children: []stml.ChildNode{
						{Kind: "state", State: &stml.StateBind{Condition: "f.g=z"}},
					},
				},
			},
			want: []string{"f.g=z"},
		},
		{
			name: "each recurses into children",
			node: stml.ChildNode{
				Kind: "each",
				Each: &stml.EachBlock{
					Children: []stml.ChildNode{
						{Kind: "state", State: &stml.StateBind{Condition: "e.h=w"}},
					},
				},
			},
			want: []string{"e.h=w"},
		},
		{
			name: "default kind yields nil",
			node: stml.ChildNode{Kind: "bind", Bind: &stml.FieldBind{}},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := childStateConditions(tt.node)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("childStateConditions() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
