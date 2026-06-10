//ff:func feature=gen-react type=test control=sequence
//ff:what resolveIndexRedirect — 첫 공개 페이지 / param 라우트 스킵 / "/" 보유 / 전부 보호 분기 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestResolveIndexRedirect(t *testing.T) {
	t.Run("first public page by file name", func(t *testing.T) {
		pages := []stml.PageSpec{
			{FileName: "settings.html"},
			{FileName: "dashboard.html"},
		}
		if got := resolveIndexRedirect(pages, nil); got != "/dashboard" {
			t.Errorf("got %q, want /dashboard", got)
		}
	})

	t.Run("protected pages are skipped", func(t *testing.T) {
		pages := []stml.PageSpec{
			{FileName: "dashboard.html"},
			{FileName: "settings.html"},
		}
		protected := map[string]bool{"dashboard.html": true}
		if got := resolveIndexRedirect(pages, protected); got != "/settings" {
			t.Errorf("got %q, want /settings", got)
		}
	})

	t.Run("param-only route is skipped", func(t *testing.T) {
		pages := []stml.PageSpec{
			{
				FileName: "item-detail.html",
				Fetches: []stml.FetchBlock{{
					OperationID: "GetItem",
					Params:      []stml.ParamBind{{Name: "id", Source: "route.ItemID"}},
				}},
			},
			{FileName: "items.html"},
		}
		if got := resolveIndexRedirect(pages, nil); got != "/items" {
			t.Errorf("got %q, want /items", got)
		}
	})

	t.Run("existing slash route wins", func(t *testing.T) {
		pages := []stml.PageSpec{
			{FileName: "about.html"},
			{FileName: "home.html", Route: "/"},
		}
		if got := resolveIndexRedirect(pages, nil); got != "" {
			t.Errorf("got %q, want empty (page already routes /)", got)
		}
	})

	t.Run("all protected falls back to /login", func(t *testing.T) {
		pages := []stml.PageSpec{{FileName: "admin.html"}}
		protected := map[string]bool{"admin.html": true}
		if got := resolveIndexRedirect(pages, protected); got != "/login" {
			t.Errorf("got %q, want /login", got)
		}
	})
}
