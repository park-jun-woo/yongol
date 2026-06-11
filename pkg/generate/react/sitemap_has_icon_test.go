//ff:func feature=gen-react type=test control=sequence
//ff:what sitemapHasIcon — 중첩 노드 data-icon 검출/아이콘 없음/nil sitemap 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapHasIcon(t *testing.T) {
	withIcon := &stml.SitemapSpec{Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
		{Label: "그룹", Menu: true, Children: []stml.SitemapNode{
			{Page: "a", Menu: false, Icon: "building"}, // hidden node still pulls the dependency
		}},
	}}}}
	if !sitemapHasIcon(withIcon) {
		t.Error("nested data-icon must be detected")
	}

	without := &stml.SitemapSpec{Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
		{Page: "a", Menu: true},
	}}}}
	if sitemapHasIcon(without) {
		t.Error("no data-icon anywhere must be false")
	}

	if sitemapHasIcon(nil) {
		t.Error("nil sitemap must be false")
	}
}
