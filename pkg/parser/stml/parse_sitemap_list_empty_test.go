//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapList_EmptyUl — 자식 없는 ul 은 빈 항목 목록 반환 검증

package stml

import "testing"

func TestParseSitemapList_EmptyUl(t *testing.T) {
	ul := firstElementNode(t, `<ul></ul>`, "ul")
	var spec SitemapSpec
	if items := parseSitemapList(ul, &spec); len(items) != 0 {
		t.Errorf("expected no items, got %+v", items)
	}
}
