//ff:func feature=gen-react type=test control=sequence
//ff:what resolveIndexRedirect — frontend.index 선언 분기(보호 인덱스 / optional strip / "/" 우선 / 미존재 폴백) 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestResolveIndexRedirect_DeclaredIndex(t *testing.T) {
	t.Run("declared index overrides file-name fallback", func(t *testing.T) {
		pages := []stml.PageSpec{
			{Name: "forgot-password", FileName: "forgot-password.html"},
			{Name: "dashboard", FileName: "dashboard.html"},
		}
		if got := resolveIndexRedirect(pages, nil, "dashboard"); got != "/dashboard" {
			t.Errorf("got %q, want /dashboard", got)
		}
	})

	t.Run("declared protected index is legal", func(t *testing.T) {
		// ProtectedRoute bounces unauthenticated visits after the redirect —
		// the dashboard-as-index admin pattern (page-flow Phase009).
		pages := []stml.PageSpec{
			{Name: "dashboard", FileName: "dashboard.html"},
			{Name: "login", FileName: "login.html"},
		}
		protected := map[string]bool{"dashboard.html": true}
		if got := resolveIndexRedirect(pages, protected, "dashboard"); got != "/dashboard" {
			t.Errorf("got %q, want /dashboard", got)
		}
	})

	t.Run("declared index strips optional segments", func(t *testing.T) {
		// route.* consumed only by actions → optional segment (":Name?");
		// the redirect has no value to fill it, so the base path is emitted.
		pages := []stml.PageSpec{
			{
				Name:     "unit-list",
				FileName: "unit-list.html",
				Actions: []stml.ActionBlock{{
					OperationID: "DeleteUnit",
					Params:      []stml.ParamBind{{Name: "building_id", Source: "route.BuildingID"}},
				}},
			},
			{Name: "about", FileName: "about.html"},
		}
		if got := resolveIndexRedirect(pages, nil, "unit-list"); got != "/unit-list" {
			t.Errorf("got %q, want /unit-list", got)
		}
	})

	t.Run("slash route wins over declared index", func(t *testing.T) {
		// Simultaneous declaration is a TM-34 ERROR; the emitter still keeps
		// a deterministic priority — the mounted "/" page suppresses the
		// redirect.
		pages := []stml.PageSpec{
			{Name: "home", FileName: "home.html", Route: "/"},
			{Name: "dashboard", FileName: "dashboard.html"},
		}
		if got := resolveIndexRedirect(pages, nil, "dashboard"); got != "" {
			t.Errorf("got %q, want empty (page already routes /)", got)
		}
	})

	t.Run("unknown declared index falls back", func(t *testing.T) {
		// TM-34 blocks an unknown page name before generate; defensively the
		// emitter falls through to the file-name fallback.
		pages := []stml.PageSpec{
			{Name: "settings", FileName: "settings.html"},
		}
		if got := resolveIndexRedirect(pages, nil, "nope"); got != "/settings" {
			t.Errorf("got %q, want /settings", got)
		}
	})
}
