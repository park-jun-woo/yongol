//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what sitemapNavPages — 전 깊이 data-page 수집 (그룹·외부 링크 제외, 문서 순서) 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapNavPages(t *testing.T) {
	nodes := []stml.SitemapNode{
		{Page: "login", Menu: true},
		{Label: "그룹", Menu: true, Children: []stml.SitemapNode{
			{Page: "signup", Menu: true, Children: []stml.SitemapNode{
				{Page: "verify-email", Menu: true},
			}},
		}},
		{Href: "https://docs.example.com", Label: "문서", Menu: true},
	}
	got := sitemapNavPages(nodes)
	want := []string{"login", "signup", "verify-email"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sitemapNavPages = %v, want %v", got, want)
	}
}
