//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what resolveRedirectTargets — 페이지명 / "/" 인덱스 / 정적 경로 매칭 / 해석 불가 무간선 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestResolveRedirectTargets(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"},
	}

	t.Run("page-name reference", func(t *testing.T) {
		got := resolveRedirectTargets("dashboard", pages, nil)
		if !reflect.DeepEqual(got, []string{"dashboard"}) {
			t.Errorf("targets = %v, want [dashboard]", got)
		}
	})

	t.Run("slash resolves to the index pages", func(t *testing.T) {
		got := resolveRedirectTargets("/", pages, []string{"dashboard"})
		if !reflect.DeepEqual(got, []string{"dashboard"}) {
			t.Errorf("targets = %v, want the index pages", got)
		}
	})

	t.Run("static path matches a parameterized route", func(t *testing.T) {
		got := resolveRedirectTargets("/buildings/42", pages, nil)
		if !reflect.DeepEqual(got, []string{"building-detail"}) {
			t.Errorf("targets = %v, want [building-detail]", got)
		}
	})

	t.Run("unresolvable values yield no edge", func(t *testing.T) {
		if got := resolveRedirectTargets("ghost", pages, nil); got != nil {
			t.Errorf("targets = %v, want nil for a ghost page name", got)
		}
		if got := resolveRedirectTargets("", pages, nil); got != nil {
			t.Errorf("targets = %v, want nil for empty", got)
		}
		if got := resolveRedirectTargets("/no/such/route", pages, nil); got != nil {
			t.Errorf("targets = %v, want nil for an unmatched static path", got)
		}
	})
}
