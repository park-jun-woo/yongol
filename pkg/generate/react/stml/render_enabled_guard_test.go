//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what renderEnabledGuard — required만/optional integer/optional 비integer/다중 optional AND 가드 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderEnabledGuard(t *testing.T) {
	pathTypes := map[string]map[string]string{
		"GetRoom": {"RoomID": "integer", "Slug": "string"},
	}

	tests := []struct {
		name string
		f    stmlparser.FetchBlock
		want string
	}{
		{
			name: "all params required → no guard",
			f: stmlparser.FetchBlock{OperationID: "GetRoom", Params: []stmlparser.ParamBind{
				{Name: "RoomID", Source: "route.RoomID", Optional: false},
			}},
			want: "",
		},
		{
			name: "optional integer param → finite guard",
			f: stmlparser.FetchBlock{OperationID: "GetRoom", Params: []stmlparser.ParamBind{
				{Name: "RoomID", Source: "route.RoomID", Optional: true},
			}},
			want: "\n    enabled: Number.isFinite(Number(RoomID)),",
		},
		{
			name: "optional non-integer param → truthy guard",
			f: stmlparser.FetchBlock{OperationID: "GetRoom", Params: []stmlparser.ParamBind{
				{Name: "Slug", Source: "route.Slug", Optional: true},
			}},
			want: "\n    enabled: !!Slug,",
		},
		{
			name: "multiple optional params AND-joined",
			f: stmlparser.FetchBlock{OperationID: "GetRoom", Params: []stmlparser.ParamBind{
				{Name: "RoomID", Source: "route.RoomID", Optional: true},
				{Name: "Slug", Source: "route.Slug", Optional: true},
			}},
			want: "\n    enabled: Number.isFinite(Number(RoomID)) && !!Slug,",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderEnabledGuard(tt.f, pathTypes); got != tt.want {
				t.Errorf("renderEnabledGuard() = %q, want %q", got, tt.want)
			}
		})
	}
}
