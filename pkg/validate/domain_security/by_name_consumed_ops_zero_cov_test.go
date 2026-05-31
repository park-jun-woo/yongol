//ff:func feature=validate type=test control=iteration dimension=1 topic=domain-security
//ff:what TestByName_ZeroCov — domain_security 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNameConsumedOps_ZeroCov(t *testing.T) {
	pages := []stml.PageSpec{{
		Fetches: []stml.FetchBlock{{
			OperationID:   "ListItems",
			NestedFetches: []stml.FetchBlock{{OperationID: "ListSub"}},
		}},
		Actions: []stml.ActionBlock{{OperationID: "CreateItem"}},
	}}
	consumed := collectConsumedOpsFromPages(pages)
	for _, op := range []string{"ListItems", "ListSub", "CreateItem"} {
		if _, ok := consumed[op]; !ok {
			t.Errorf("collectConsumedOpsFromPages missing %q", op)
		}
	}

	out := map[string]struct{}{}
	collectFetchOpsRecursive(pages[0].Fetches[0], out)
	if _, ok := out["ListSub"]; !ok {
		t.Errorf("collectFetchOpsRecursive missing nested op")
	}
}
