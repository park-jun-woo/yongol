//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what resolveHasRefresh — manifest refresh_field / STML auth.refresh 캡처 / 둘 다 없음(dead) / nil 안전 판정 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveHasRefresh(t *testing.T) {
	newFS := func(fa *manifest.FrontendAuth, pages []stml.PageSpec) *yongol.Fullstack {
		return &yongol.Fullstack{
			Manifest:  &manifest.ProjectConfig{Frontend: manifest.Frontend{Auth: fa}},
			STMLPages: pages,
		}
	}
	refreshCapture := []stml.PageSpec{{Actions: []stml.ActionBlock{{
		Captures: []stml.CaptureBind{
			{RespField: "access_token", Sink: "auth.token"},
			{RespField: "refresh_token", Sink: "auth.refresh"},
		},
	}}}}
	tokenOnlyCapture := []stml.PageSpec{{Actions: []stml.ActionBlock{{
		Captures: []stml.CaptureBind{{RespField: "access_token", Sink: "auth.token"}},
	}}}}

	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want bool
	}{
		{"nil fs", nil, false},
		{"manifest refresh_field", newFS(&manifest.FrontendAuth{TokenField: "access_token", RefreshField: "refresh_token"}, nil), true},
		{"STML auth.refresh capture", newFS(&manifest.FrontendAuth{TokenField: "access_token"}, refreshCapture), true},
		{"bearer without refresh (dead)", newFS(&manifest.FrontendAuth{TokenField: "access_token"}, tokenOnlyCapture), false},
		{"no frontend.auth block", newFS(nil, nil), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := resolveHasRefresh(c.fs); got != c.want {
				t.Errorf("resolveHasRefresh = %v, want %v", got, c.want)
			}
		})
	}
}
