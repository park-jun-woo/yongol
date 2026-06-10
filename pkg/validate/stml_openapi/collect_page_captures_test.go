//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what collectPageCaptures — 페이지 액션별 data-capture 바인딩을 (파일,바인딩) 쌍으로 수집 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectPageCaptures(t *testing.T) {
	t.Run("no pages returns nil", func(t *testing.T) {
		if got := collectPageCaptures(nil); got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("collects captures across pages", func(t *testing.T) {
		pages := []stml.PageSpec{
			{
				FileName: "login.html",
				Actions: []stml.ActionBlock{
					{Captures: []stml.CaptureBind{
						{RespField: "access_token", Sink: "auth.token"},
						{RespField: "refresh_token", Sink: "auth.refresh"},
					}},
				},
			},
			{
				FileName: "noop.html",
				Actions:  []stml.ActionBlock{{}}, // no captures
			},
		}
		got := collectPageCaptures(pages)
		if len(got) != 2 {
			t.Fatalf("expected 2 captures, got %d: %v", len(got), got)
		}
		if got[0].File != "login.html" || got[0].Bind.Sink != "auth.token" {
			t.Errorf("unexpected first capture: %+v", got[0])
		}
		if got[1].Bind.RespField != "refresh_token" {
			t.Errorf("unexpected second capture: %+v", got[1])
		}
	})
}
