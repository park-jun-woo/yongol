//ff:func feature=validate type=test control=sequence topic=stml-statemachine
//ff:what findRedirectTargetPage — 매칭 페이지 반환 / 미해석 경로 nil 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFindRedirectTargetPage(t *testing.T) {
	pages := []stml.PageSpec{
		{FileName: "login.html"},
		{FileName: "dashboard.html"},
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
}
