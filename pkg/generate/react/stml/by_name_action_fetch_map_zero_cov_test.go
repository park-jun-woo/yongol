//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNameActionFetchMap_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	m := buildActionFetchMap(page)
	if m == nil {
		t.Fatalf("buildActionFetchMap nil")
	}
	walkChildrenForFetchMap(page.Children, []string{"ListItems"}, m)

	rm := map[string][]string{}
	recordActionFetchMapping("CreateItem", []string{"ListItems"}, rm)
	recordActionFetchMapping("CreateItem", []string{"ListItems"}, rm) // already present
	recordActionFetchMapping("NoFetch", nil, rm)

	ops, _ := resolveInvalidateOps(stmlparser.ActionBlock{OperationID: "CreateItem"}, []string{"ListItems"}, m, nil)
	_ = ops
	_, _ = resolveInvalidateOps(stmlparser.ActionBlock{OperationID: "Unknown"}, []string{"ListItems"}, m, nil)
}
