//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what renderOptionalMutationGuard — optional 없음/단일/다중/혼합 OR 가드식 검증 (BUG-136)

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderOptionalMutationGuard(t *testing.T) {
	tests := []struct {
		name string
		a    stmlparser.ActionBlock
		want string
	}{
		{
			name: "no params → empty",
			a:    stmlparser.ActionBlock{OperationID: "CreateReservation"},
			want: "",
		},
		{
			name: "all required → empty",
			a: stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
				{Name: "ReservationID", Source: "route.ReservationID", Optional: false},
				{Name: "RoomID", Source: "route.RoomID", Optional: false},
			}},
			want: "",
		},
		{
			name: "single optional route param → null guard",
			a: stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
				{Name: "ReservationID", Source: "route.ReservationID", Optional: true},
			}},
			want: "ReservationID == null",
		},
		{
			name: "multiple optional params OR-joined",
			a: stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
				{Name: "ReservationID", Source: "route.ReservationID", Optional: true},
				{Name: "RoomID", Source: "route.RoomID", Optional: true},
			}},
			want: "ReservationID == null || RoomID == null",
		},
		{
			name: "mixed optional and required → only optionals guarded",
			a: stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
				{Name: "RoomID", Source: "route.RoomID", Optional: false},
				{Name: "ReservationID", Source: "route.ReservationID", Optional: true},
			}},
			want: "ReservationID == null",
		},
		{
			name: "optional non-route source kept verbatim",
			a: stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
				{Name: "Slug", Source: "item.Slug", Optional: true},
			}},
			want: "item.Slug == null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderOptionalMutationGuard(tt.a); got != tt.want {
				t.Errorf("renderOptionalMutationGuard() = %q, want %q", got, tt.want)
			}
		})
	}
}
