//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what guardLifecycleToJSX — loading/error/empty 생명주기 노드를 JSX 표현식으로 변환 검증

package stml

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGuardLifecycleToJSX(t *testing.T) {
	tests := []struct {
		name      string
		ref       stml.GuardRef
		lifecycle string
		dataVar   string
		want      string
	}{
		{
			name:      "loading",
			ref:       stml.GuardRef{Model: "items", Field: "list"},
			lifecycle: "loading",
			dataVar:   "data",
			want:      "dataLoading",
		},
		{
			name:      "error",
			ref:       stml.GuardRef{Model: "items", Field: "list"},
			lifecycle: "error",
			dataVar:   "data",
			want:      "dataError",
		},
		{
			name:      "empty",
			ref:       stml.GuardRef{Model: "items", Field: "list"},
			lifecycle: "empty",
			dataVar:   "data",
			want:      "data.items?.list?.length === 0",
		},
		{
			name:      "empty custom var",
			ref:       stml.GuardRef{Model: "rows", Field: "data"},
			lifecycle: "empty",
			dataVar:   "ctx",
			want:      "ctx.rows?.data?.length === 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertGuardLifecycleToJSX(t, tt.ref, tt.lifecycle, tt.dataVar, tt.want)
		})
	}
}
