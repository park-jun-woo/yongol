//ff:func feature=gen-react type=test control=sequence
//ff:what TestRenderSitemapGroupQuery — 페이지 규약 쿼리 키·bearer enabled 게이트·cookie 무게이트 스냅샷 검증

package react

import (
	"strings"
	"testing"
)

func TestRenderSitemapGroupQuery(t *testing.T) {
	t.Run("bearer mode gates the query on the stored token", func(t *testing.T) {
		var sb strings.Builder
		renderSitemapGroupQuery(&sb, "ListMyBuildings", true)
		want := "  const { data: listMyBuildingsData } = useQuery({\n" +
			"    queryKey: ['ListMyBuildings'],\n" +
			"    queryFn: () => api.ListMyBuildings(),\n" +
			"    enabled: !!token,\n" +
			"  })\n"
		if sb.String() != want {
			t.Errorf("got:\n%s\nwant:\n%s", sb.String(), want)
		}
	})

	t.Run("without the gate the query fires unconditionally (cookie/no-auth)", func(t *testing.T) {
		var sb strings.Builder
		renderSitemapGroupQuery(&sb, "ListMyBuildings", false)
		got := sb.String()
		if strings.Contains(got, "enabled:") {
			t.Errorf("cookie mode must not gate the query:\n%s", got)
		}
		// query key = page fetch convention — data-invalidates hits it
		if !strings.Contains(got, "queryKey: ['ListMyBuildings'],") {
			t.Errorf("query key must follow the page convention:\n%s", got)
		}
	})
}
