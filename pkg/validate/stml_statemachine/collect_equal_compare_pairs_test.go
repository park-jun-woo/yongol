//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what collectEqualComparePairs — nil/&&/||/그룹/=/비등호/단항별 무조건 "=" 쌍 수집 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectEqualComparePairs(t *testing.T) {
	eq := func(model, value string) *stml.GuardExpr {
		return &stml.GuardExpr{Kind: stml.GuardCompare, Op: "=", Ref: stml.GuardRef{Model: model, Field: "status"}, Value: value}
	}
	tests := []struct {
		name string
		expr *stml.GuardExpr
		want []comparePair
	}{
		{name: "nil node yields nil", expr: nil, want: nil},
		{
			name: "equal compare leaf yields pair",
			expr: eq("workflow", "draft"),
			want: []comparePair{{Model: "workflow", Value: "draft"}},
		},
		{
			name: "non-equal compare yields nothing",
			expr: &stml.GuardExpr{Kind: stml.GuardCompare, Op: "!=", Ref: stml.GuardRef{Model: "a", Field: "status"}, Value: "x"},
			want: nil,
		},
		{
			name: "and recurses into both sides",
			expr: &stml.GuardExpr{Kind: stml.GuardBinary, Op: "&&", Left: eq("a", "x"), Right: eq("b", "y")},
			want: []comparePair{{Model: "a", Value: "x"}, {Model: "b", Value: "y"}},
		},
		{
			name: "or yields nothing",
			expr: &stml.GuardExpr{Kind: stml.GuardBinary, Op: "||", Left: eq("a", "x"), Right: eq("b", "y")},
			want: nil,
		},
		{
			name: "group recurses into operand",
			expr: &stml.GuardExpr{Kind: stml.GuardGroup, Operand: eq("g", "z")},
			want: []comparePair{{Model: "g", Value: "z"}},
		},
		{
			name: "unary negation yields nothing",
			expr: &stml.GuardExpr{Kind: stml.GuardUnary, Op: "!", Operand: eq("a", "x")},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertComparePairs(t, collectEqualComparePairs(tt.expr), tt.want)
		})
	}
}
