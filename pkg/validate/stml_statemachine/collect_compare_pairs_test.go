//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what collectComparePairs — nil/비교/이항/단항/그룹/생명주기 노드별 (model,value) DOM순 수집 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectComparePairs(t *testing.T) {
	tests := []struct {
		name string
		expr *stml.GuardExpr
		want []comparePair
	}{
		{name: "nil node yields nil", expr: nil, want: nil},
		{
			name: "single compare leaf",
			expr: cmp("workflow", "draft"),
			want: []comparePair{{Model: "workflow", Value: "draft"}},
		},
		{
			name: "lifecycle leaf contributes nothing",
			expr: &stml.GuardExpr{Kind: stml.GuardLifecycle, Ref: stml.GuardRef{Model: "workflow", Field: "status"}, Lifecycle: "loading"},
			want: nil,
		},
		{
			name: "binary recurses left then right in DOM order",
			expr: &stml.GuardExpr{Kind: stml.GuardBinary, Op: "&&", Left: cmp("a", "x"), Right: cmp("b", "y")},
			want: []comparePair{{Model: "a", Value: "x"}, {Model: "b", Value: "y"}},
		},
		{
			name: "unary recurses into operand",
			expr: &stml.GuardExpr{Kind: stml.GuardUnary, Op: "!", Operand: cmp("a", "x")},
			want: []comparePair{{Model: "a", Value: "x"}},
		},
		{
			name: "group recurses into operand",
			expr: &stml.GuardExpr{Kind: stml.GuardGroup, Operand: cmp("g", "z")},
			want: []comparePair{{Model: "g", Value: "z"}},
		},
		{
			name: "nested binary inside group preserves order",
			expr: &stml.GuardExpr{Kind: stml.GuardGroup, Operand: &stml.GuardExpr{
				Kind: stml.GuardBinary, Op: "||", Left: cmp("a", "x"), Right: cmp("b", "y"),
			}},
			want: []comparePair{{Model: "a", Value: "x"}, {Model: "b", Value: "y"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertComparePairs(t, collectComparePairs(tt.expr), tt.want)
		})
	}
}
