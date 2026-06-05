//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what guardCompareToJSX — ref op value 비교 노드를 JSX 비교식으로 변환 검증

package stml

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGuardCompareToJSX(t *testing.T) {
	tests := []struct {
		name    string
		ref     stml.GuardRef
		op      string
		value   string
		dataVar string
		want    string
	}{
		{
			name:    "equality",
			ref:     stml.GuardRef{Model: "workflow", Field: "status"},
			op:      "=",
			value:   "active",
			dataVar: "data",
			want:    "data.workflow?.status === 'active'",
		},
		{
			name:    "inequality",
			ref:     stml.GuardRef{Model: "workflow", Field: "status"},
			op:      "!=",
			value:   "draft",
			dataVar: "data",
			want:    "data.workflow?.status !== 'draft'",
		},
		{
			name:    "relational greater equal",
			ref:     stml.GuardRef{Model: "order", Field: "count"},
			op:      ">=",
			value:   "3",
			dataVar: "data",
			want:    "data.order?.count >= '3'",
		},
		{
			name:    "custom data var",
			ref:     stml.GuardRef{Model: "user", Field: "role"},
			op:      "=",
			value:   "owner",
			dataVar: "ctx",
			want:    "ctx.user?.role === 'owner'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertGuardCompareToJSX(t, tt.ref, tt.op, tt.value, tt.dataVar, tt.want)
		})
	}
}
