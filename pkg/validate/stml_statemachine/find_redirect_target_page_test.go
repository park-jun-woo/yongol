//ff:func feature=validate type=test control=sequence topic=stml-statemachine
//ff:what findRedirectTargetPage — 정적 경로·페이지명 참조의 매칭 페이지 반환 / 미해석 nil 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFindRedirectTargetPage(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}

	t.Run("resolves to matching page", func(t *testing.T) {
		got := findRedirectTargetPage("/dashboard", pages)
		if got == nil || got.FileName != "dashboard.html" {
			t.Errorf("expected dashboard.html, got %+v", got)
		}
	})

	t.Run("unresolved path returns nil", func(t *testing.T) {
		if got := findRedirectTargetPage("/nope", pages); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("page-name reference resolves by name", func(t *testing.T) {
		got := findRedirectTargetPage("dashboard", pages)
		if got == nil || got.FileName != "dashboard.html" {
			t.Errorf("expected dashboard.html, got %+v", got)
		}
	})

	t.Run("unknown page name returns nil", func(t *testing.T) {
		if got := findRedirectTargetPage("dashbord", pages); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})
}
