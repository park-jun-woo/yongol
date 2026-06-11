//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm43Cause — 등재×메뉴 비렌더 / 미등재 × 들어오는 간선 유무 4분류 문구 검증

package stml_openapi

import (
	"strings"
	"testing"
)

func TestTM43Cause(t *testing.T) {
	g := &pageGraph{
		InSitemap:   map[string]bool{"building-detail": true},
		MenuBlocked: map[string]string{"building-detail": `route "/buildings/:BuildingID" has required param :BuildingID`},
	}

	t.Run("listed but blocked, no incoming", func(t *testing.T) {
		got := tm43Cause(g, "building-detail", false)
		if !strings.Contains(got, "in sitemap but not menu-rendered") || !strings.Contains(got, ":BuildingID") {
			t.Errorf("cause = %q, want the listed-but-blocked classification with the reason", got)
		}
		if !strings.Contains(got, "no data-link/data-redirect/breadcrumb edge points to it") {
			t.Errorf("cause = %q, want the no-incoming clause", got)
		}
	})

	t.Run("not listed, only unreachable sources", func(t *testing.T) {
		got := tm43Cause(g, "orphan", true)
		if !strings.Contains(got, "not in the sitemap") || !strings.Contains(got, "only unreachable pages link to it") {
			t.Errorf("cause = %q, want the unlisted classification with the unreachable-sources clause", got)
		}
	})
}
