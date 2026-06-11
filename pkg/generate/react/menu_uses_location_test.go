//ff:func feature=gen-react type=test control=sequence
//ff:what menuUsesLocation — 중첩 항목의 prefix 검출/prefix 없음 false 검증

package react

import "testing"

func TestMenuUsesLocation(t *testing.T) {
	nested := []sitemapMenuItem{{Kind: "group", Children: []sitemapMenuItem{
		{Kind: "page", Prefixes: []string{"/buildings/"}},
	}}}
	if !menuUsesLocation(nested) {
		t.Error("nested prefixes must be detected")
	}

	plain := []sitemapMenuItem{{Kind: "page"}, {Kind: "external"}}
	if menuUsesLocation(plain) {
		t.Error("no prefixes anywhere must be false")
	}
}
