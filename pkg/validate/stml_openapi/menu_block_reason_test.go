//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what menuBlockReason — data-menu="false" / 깊이 초과 / 필수 파라미터 사유와 렌더 가능("") 판정 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestMenuBlockReason(t *testing.T) {
	t.Run("data-menu=false wins over everything", func(t *testing.T) {
		node := stml.SitemapNode{Page: "member-list", Menu: false}
		got := menuBlockReason(node, 3, []string{"/members/:MemberID"})
		if !strings.Contains(got, `data-menu="false"`) {
			t.Errorf("reason = %q, want the data-menu=\"false\" reason", got)
		}
	})

	t.Run("depth beyond 2 blocks", func(t *testing.T) {
		node := stml.SitemapNode{Page: "building-detail", Menu: true}
		got := menuBlockReason(node, 3, []string{"/buildings"})
		if !strings.Contains(got, "depth 3") {
			t.Errorf("reason = %q, want the depth reason", got)
		}
	})

	t.Run("required route param blocks", func(t *testing.T) {
		node := stml.SitemapNode{Page: "building-detail", Menu: true}
		got := menuBlockReason(node, 2, []string{"/buildings/:BuildingID"})
		if !strings.Contains(got, ":BuildingID") {
			t.Errorf("reason = %q, want the required param named", got)
		}
	})

	t.Run("optional segment and depth 2 render", func(t *testing.T) {
		node := stml.SitemapNode{Page: "building-list", Menu: true}
		if got := menuBlockReason(node, 2, []string{"/buildings/:Filter?"}); got != "" {
			t.Errorf("reason = %q, want empty (renderable)", got)
		}
	})

	t.Run("group without patterns renders at depth 1", func(t *testing.T) {
		node := stml.SitemapNode{Label: "관리", Menu: true}
		if got := menuBlockReason(node, 1, nil); got != "" {
			t.Errorf("reason = %q, want empty (renderable group)", got)
		}
	})
}
