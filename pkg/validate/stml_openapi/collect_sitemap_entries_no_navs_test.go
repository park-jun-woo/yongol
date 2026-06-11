//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCollectSitemapEntries_NoNavs — nav 블록 없는 sitemap 은 빈 엔트리 목록 반환 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectSitemapEntries_NoNavs(t *testing.T) {
	if entries := collectSitemapEntries(&stml.SitemapSpec{FileName: "sitemap.html"}); len(entries) != 0 {
		t.Errorf("expected no entries, got %+v", entries)
	}
}
