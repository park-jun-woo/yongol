//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what PageConsumedOps — fetch(재귀)·action·component api.X 소비 op 합집합 / 컴포넌트 스캔 skip 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPageConsumedOps(t *testing.T) {
	page := stml.PageSpec{
		Name: "page",
		Fetches: []stml.FetchBlock{{
			OperationID:   "ListItems",
			NestedFetches: []stml.FetchBlock{{OperationID: "ListSub"}},
			Components:    []stml.ComponentRef{{Name: "DatePicker"}},
		}},
		Actions: []stml.ActionBlock{{OperationID: "CreateItem"}},
	}

	t.Run("unions fetch, action and component ops", func(t *testing.T) {
		specsDir := t.TempDir()
		writeComponent(t, specsDir, "DatePicker", `export function DatePicker() {
	api.LoadDates();
	api.Bogus();
	return null;
}`)
		ops := map[string]struct{}{"LoadDates": {}, "Other": {}}

		out := PageConsumedOps(page, specsDir, ops)
		for _, id := range []string{"ListItems", "ListSub", "CreateItem", "LoadDates"} {
			if _, ok := out[id]; !ok {
				t.Errorf("missing consumed operationId %q in %+v", id, out)
			}
		}
		if _, ok := out["Bogus"]; ok {
			t.Errorf("non-existent op Bogus must not be consumed (intersection only)")
		}
	})

	t.Run("empty specsDir skips component scan", func(t *testing.T) {
		out := PageConsumedOps(page, "", nil)
		for _, id := range []string{"ListItems", "ListSub", "CreateItem"} {
			if _, ok := out[id]; !ok {
				t.Errorf("missing consumed operationId %q in %+v", id, out)
			}
		}
		if _, ok := out["LoadDates"]; ok {
			t.Errorf("component op LoadDates must not be consumed when specsDir is empty")
		}
	})
}
