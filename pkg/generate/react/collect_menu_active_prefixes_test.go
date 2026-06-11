//ff:func feature=gen-react type=test control=sequence
//ff:what collectMenuActivePrefixes — 렌더 자식 스킵/비렌더 자손 prefix 수집/쓸모없는 prefix 제외 검증

package react

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectMenuActivePrefixes(t *testing.T) {
	patterns := map[string]string{
		"building-detail": "/buildings/:BuildingID",
		"building-rights": "/building-rights",
		"audit-log":       "/:OrgID/audit",
	}

	t.Run("rendered child is skipped whole, hidden ones contribute", func(t *testing.T) {
		nodes := []stml.SitemapNode{
			// renderable at depth 2 — its own subtree highlights it, not the parent
			{Page: "building-rights", Label: "권리관계", Menu: true, Children: []stml.SitemapNode{
				{Page: "building-detail", Label: "상세", Menu: true},
			}},
			// hidden by data-menu="false" — contributes itself and its subtree
			{Page: "building-detail", Label: "상세", Menu: false},
		}
		var out []string
		collectMenuActivePrefixes(nodes, 2, patterns, &out)
		if want := []string{"/buildings/"}; !reflect.DeepEqual(out, want) {
			t.Errorf("prefixes = %v, want %v", out, want)
		}
	})

	t.Run("useless prefixes are dropped", func(t *testing.T) {
		nodes := []stml.SitemapNode{{Page: "audit-log", Label: "감사", Menu: true}}
		var out []string
		collectMenuActivePrefixes(nodes, 3, patterns, &out)
		if len(out) != 0 {
			t.Errorf(`leading-parameter route must not yield the "/" prefix: %v`, out)
		}
	})
}
