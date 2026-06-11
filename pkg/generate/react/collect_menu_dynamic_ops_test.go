//ff:func feature=gen-react type=test control=sequence
//ff:what TestCollectMenuDynamicOps — 문서 순서·중첩 포함·중복 op 제거 수집 검증

package react

import (
	"reflect"
	"testing"
)

func TestCollectMenuDynamicOps(t *testing.T) {
	if got := collectMenuDynamicOps(nil); len(got) != 0 {
		t.Errorf("empty menu ops = %v", got)
	}
	items := []sitemapMenuItem{
		{Kind: "page", Label: "대시보드"},
		{Kind: "group", Label: "내 건물", Fetch: "ListMyBuildings"},
		{Kind: "group", Label: "그룹", Children: []sitemapMenuItem{
			{Kind: "group", Label: "내 계약", Fetch: "ListMyContracts"},
			{Kind: "group", Label: "건물 다시", Fetch: "ListMyBuildings"},
		}},
	}
	want := []string{"ListMyBuildings", "ListMyContracts"}
	if got := collectMenuDynamicOps(items); !reflect.DeepEqual(got, want) {
		t.Errorf("ops = %v, want %v (document order, deduplicated)", got, want)
	}
}
