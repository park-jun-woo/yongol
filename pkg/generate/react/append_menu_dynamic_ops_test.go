//ff:func feature=gen-react type=test control=sequence
//ff:what TestAppendMenuDynamicOps — seen 집합에 의한 중복 제거와 자식 재귀 누적 검증

package react

import (
	"reflect"
	"testing"
)

func TestAppendMenuDynamicOps(t *testing.T) {
	items := []sitemapMenuItem{
		{Kind: "group", Fetch: "ListA"},
		{Kind: "group", Children: []sitemapMenuItem{{Kind: "group", Fetch: "ListB"}}},
	}
	var ops []string
	appendMenuDynamicOps(items, map[string]bool{"ListA": true}, &ops)
	if !reflect.DeepEqual(ops, []string{"ListB"}) {
		t.Errorf("ops = %v, want already-seen ListA skipped and nested ListB appended", ops)
	}
}
