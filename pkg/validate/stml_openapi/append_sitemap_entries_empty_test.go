//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestAppendSitemapEntries_EmptyNodes — nil 노드 목록은 엔트리를 추가하지 않음 검증

package stml_openapi

import "testing"

func TestAppendSitemapEntries_EmptyNodes(t *testing.T) {
	var entries []sitemapEntry
	appendSitemapEntries(nil, "nav[0]", &entries)
	if len(entries) != 0 {
		t.Errorf("expected no entries, got %+v", entries)
	}
}
