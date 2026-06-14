//ff:func feature=gen-react type=test control=iteration dimension=2
//ff:what hasRefreshCaptures — auth.refresh 캡처 존재/타 sink만/캡처 없음/빈 페이지 판정 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestHasRefreshCaptures(t *testing.T) {
	tests := []struct {
		name  string
		pages []stml.PageSpec
		want  bool
	}{
		{
			name:  "nil pages",
			pages: nil,
			want:  false,
		},
		{
			name:  "page without actions",
			pages: []stml.PageSpec{{Name: "home"}},
			want:  false,
		},
		{
			name: "action with token sink only",
			pages: []stml.PageSpec{{Actions: []stml.ActionBlock{
				{Captures: []stml.CaptureBind{{RespField: "access_token", Sink: "auth.token"}}},
			}}},
			want: false,
		},
		{
			name: "action with auth.refresh sink",
			pages: []stml.PageSpec{{Actions: []stml.ActionBlock{
				{Captures: []stml.CaptureBind{
					{RespField: "access_token", Sink: "auth.token"},
					{RespField: "refresh_token", Sink: "auth.refresh"},
				}},
			}}},
			want: true,
		},
		{
			name: "refresh capture on a later page",
			pages: []stml.PageSpec{
				{Name: "first", Actions: []stml.ActionBlock{
					{Captures: []stml.CaptureBind{{RespField: "role", Sink: "auth.claims.role"}}},
				}},
				{Name: "second", Actions: []stml.ActionBlock{
					{Captures: []stml.CaptureBind{{RespField: "refresh_token", Sink: "auth.refresh"}}},
				}},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasRefreshCaptures(tt.pages); got != tt.want {
				t.Errorf("hasRefreshCaptures() = %v, want %v", got, tt.want)
			}
		})
	}
}
