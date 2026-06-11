//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapList — ul 직속 li 만 문서 순서대로 변환, li 아닌 자식·텍스트는 무시 검증

package stml

import "testing"

func TestParseSitemapList(t *testing.T) {
	ul := firstElementNode(t, `
<ul>
  <li data-page="first">첫째</li>
  <span>장식</span>
  <li data-page="second">둘째</li>
</ul>`, "ul")
	var spec SitemapSpec
	items := parseSitemapList(ul, &spec)
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d: %+v", len(items), items)
	}
	if items[0].Page != "first" || items[1].Page != "second" {
		t.Errorf("items = %+v, want first then second (document order)", items)
	}
}
