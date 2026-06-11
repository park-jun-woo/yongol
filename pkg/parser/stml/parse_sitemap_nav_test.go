//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapNav — data-layout/data-entry 속성, 복수 ul 병합, 동적 그룹 어휘 1급 승격 검증

package stml

import "testing"

func TestParseSitemapNav(t *testing.T) {
	t.Run("layout, entry and items from multiple uls", func(t *testing.T) {
		nav := firstElementNode(t, `
<nav data-sitemap data-layout="app" data-entry>
  <ul><li data-page="home">홈</li></ul>
  <ul><li data-page="about">소개</li></ul>
</nav>`, "nav")
		var spec SitemapSpec
		out := parseSitemapNav(nav, &spec)
		if out.Layout != "app" || !out.Entry {
			t.Errorf("nav = layout %q entry %v, want app/true", out.Layout, out.Entry)
		}
		if len(out.Items) != 2 || out.Items[0].Page != "home" || out.Items[1].Page != "about" {
			t.Errorf("Items = %+v, want home then about merged from both uls", out.Items)
		}
	})

	t.Run("defaults without attributes", func(t *testing.T) {
		nav := firstElementNode(t, `<nav data-sitemap><ul><li data-page="login">로그인</li></ul></nav>`, "nav")
		var spec SitemapSpec
		out := parseSitemapNav(nav, &spec)
		if out.Layout != "" || out.Entry {
			t.Errorf("nav = layout %q entry %v, want empty/false", out.Layout, out.Entry)
		}
	})

	t.Run("dynamic group vocabulary graduates onto the group node (Phase007)", func(t *testing.T) {
		nav := firstElementNode(t, `<nav data-sitemap><ul><li>건물<ul data-fetch="ListBuildings" data-each="buildings" data-link="building-detail" data-label-field="name"></ul></li></ul></nav>`, "nav")
		var spec SitemapSpec
		out := parseSitemapNav(nav, &spec)
		if len(out.Items) != 1 {
			t.Fatalf("Items = %+v, want one group node", out.Items)
		}
		n := out.Items[0]
		if n.Fetch != "ListBuildings" || n.Each != "buildings" || n.Link != "building-detail" || n.LabelField != "name" {
			t.Errorf("dynamic group node = %+v", n)
		}
	})
}
