//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCollectConsumedOpsLayouts — 레이아웃 data-logout op 의 소비 집합 편입 검증 (값 없는 logout 은 무기여)

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectConsumedOpsLayouts(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:    "page",
		Fetches: []stml.FetchBlock{{OperationID: "ListItems"}},
	}}
	layouts := []stml.LayoutSpec{
		{Name: "app", Logout: &stml.LogoutSpec{OperationID: "Logout"}},
		{Name: "auth"}, // no logout
		{Name: "bare", Logout: &stml.LogoutSpec{}}, // valueless — no op to consume
	}

	out := collectConsumedOps(pages, layouts, "", nil)
	for _, want := range []string{"ListItems", "Logout"} {
		if _, ok := out[want]; !ok {
			t.Errorf("missing consumed op %q", want)
		}
	}
	if _, ok := out[""]; ok {
		t.Errorf("valueless data-logout must not contribute an empty operationId")
	}
	if len(out) != 2 {
		t.Errorf("expected 2 consumed ops, got %d: %+v", len(out), out)
	}
}
