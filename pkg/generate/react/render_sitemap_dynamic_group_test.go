//ff:func feature=gen-react type=test control=sequence
//ff:what TestRenderSitemapDynamicGroup — 0개 시 헤더째 비렌더 게이트·경로 치환 NavLink·label-field·index fallback 키·정적 자식 선행 스냅샷 검증

package react

import (
	"strings"
	"testing"
)

func TestRenderSitemapDynamicGroup(t *testing.T) {
	group := sitemapMenuItem{
		Kind: "group", Label: "내 건물",
		Fetch: "ListMyBuildings", Each: "items", LabelField: "building_name",
		ItemToAttr: "to={`/buildings/${item.building_id}`}", ItemKey: "item.building_id",
	}

	t.Run("full snapshot — zero-item omission gate, mapped NavLinks", func(t *testing.T) {
		var sb strings.Builder
		renderSitemapDynamicGroup(&sb, group, "          ", false)
		want := "          {(listMyBuildingsData?.items ?? []).length > 0 && (\n" +
			"            <li>\n" +
			"              <span>내 건물</span>\n" +
			"              <ul>\n" +
			"                {(listMyBuildingsData?.items ?? []).map((item) => (\n" +
			"                  <li key={item.building_id}><NavLink to={`/buildings/${item.building_id}`} end>{item.building_name}</NavLink></li>\n" +
			"                ))}\n" +
			"              </ul>\n" +
			"            </li>\n" +
			"          )}\n"
		if sb.String() != want {
			t.Errorf("got:\n%s\nwant:\n%s", sb.String(), want)
		}
	})

	t.Run("no item key falls back to the positional index", func(t *testing.T) {
		plain := group
		plain.ItemToAttr = `to="/buildings"`
		plain.ItemKey = ""
		var sb strings.Builder
		renderSitemapDynamicGroup(&sb, plain, "", false)
		got := sb.String()
		if !strings.Contains(got, ".map((item, index) => (") || !strings.Contains(got, "<li key={index}>") {
			t.Errorf("index fallback missing:\n%s", got)
		}
	})

	t.Run("static children render before the dynamic items", func(t *testing.T) {
		mixed := group
		mixed.Children = []sitemapMenuItem{{Kind: "page", Label: "건물 추가", To: "/buildings/new"}}
		var sb strings.Builder
		renderSitemapDynamicGroup(&sb, mixed, "", false)
		got := sb.String()
		static := strings.Index(got, "건물 추가")
		dynamic := strings.Index(got, ".map((item)")
		if static == -1 || dynamic == -1 || static > dynamic {
			t.Errorf("static child must precede the dynamic items:\n%s", got)
		}
	})
}
