//ff:func feature=gen-react type=test control=sequence
//ff:what sitemapNodePatterns — 페이지 해석/그룹·외부 링크 nil/미해석 페이지 nil 검증

package react

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapNodePatterns(t *testing.T) {
	patterns := map[string]string{"dashboard": "/dashboard"}

	if got := sitemapNodePatterns(stml.SitemapNode{Page: "dashboard"}, patterns); !reflect.DeepEqual(got, []string{"/dashboard"}) {
		t.Errorf("page node = %v", got)
	}
	if got := sitemapNodePatterns(stml.SitemapNode{Label: "그룹"}, patterns); got != nil {
		t.Errorf("group node must be nil, got %v", got)
	}
	if got := sitemapNodePatterns(stml.SitemapNode{Href: "https://x"}, patterns); got != nil {
		t.Errorf("external node must be nil, got %v", got)
	}
	if got := sitemapNodePatterns(stml.SitemapNode{Page: "ghost"}, patterns); got != nil {
		t.Errorf("unresolved page must be nil, got %v", got)
	}
}
