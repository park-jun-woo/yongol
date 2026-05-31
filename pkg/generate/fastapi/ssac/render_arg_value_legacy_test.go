//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderArgValueLegacy — Location 미설정 시 source 기반 매핑 분기 전체 커버
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderArgValueLegacy(t *testing.T) {
	tests := []struct {
		name string
		arg  ir.FieldArg
		col  string
		want string
	}{
		{"EmptyColRequest", ir.FieldArg{Source: "request"}, "", "params"},
		{"EmptyColCurrentUser", ir.FieldArg{Source: "currentUser"}, "", "current_user"},
		{"EmptyColVariable", ir.FieldArg{Source: "wf"}, "", "wf"},
		{"EmptyColEmptySourceDefaultsRequest", ir.FieldArg{}, "", "params"},
		{"ColRequest", ir.FieldArg{Source: "request"}, "id", "id"},
		{"ColCurrentUser", ir.FieldArg{Source: "currentUser"}, "org_id", `current_user["org_id"]`},
		{"ColVariable", ir.FieldArg{Source: "wf"}, "status", "wf.status"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderArgValueLegacy(tt.arg, tt.col); got != tt.want {
				t.Errorf("renderArgValueLegacy() = %q, want %q", got, tt.want)
			}
		})
	}
}
