//ff:func feature=gen-react type=test control=sequence
//ff:what menuHasNavLink — 중첩 페이지 항목 검출/그룹·외부 링크만이면 false 검증

package react

import "testing"

func TestMenuHasNavLink(t *testing.T) {
	nested := []sitemapMenuItem{{Kind: "group", Children: []sitemapMenuItem{{Kind: "page"}}}}
	if !menuHasNavLink(nested) {
		t.Error("nested page item must be detected")
	}

	noPages := []sitemapMenuItem{{Kind: "group"}, {Kind: "external"}}
	if menuHasNavLink(noPages) {
		t.Error("groups and external links alone must be false")
	}
}
