//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what collectStateConditions — 자식 트리 순회로 모든 data-state 조건 수집 (빈/다중) 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectStateConditions(t *testing.T) {
	tests := []struct {
		name     string
		children []stml.ChildNode
		want     []string
	}{
		{
			name:     "empty list yields nil",
			children: nil,
			want:     nil,
		},
		{
			name: "multiple children aggregated in DOM order",
			children: []stml.ChildNode{
				{Kind: "state", State: &stml.StateBind{Condition: "a.b=x"}},
				{Kind: "bind", Bind: &stml.FieldBind{}},
				{Kind: "state", State: &stml.StateBind{Condition: "c.d=y"}},
			},
			want: []string{"a.b=x", "c.d=y"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := collectStateConditions(tt.children)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collectStateConditions() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
