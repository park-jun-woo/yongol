//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what sortBreadcrumbRoutes — 정적 패턴 우선 / 동률 문서 순서 안정 / 원본 무변경 검증

package react

import (
	"reflect"
	"testing"
)

func TestSortBreadcrumbRoutes(t *testing.T) {
	trails := []breadcrumbTrail{
		{Page: "building-detail", Pattern: "/buildings/:BuildingID"},
		{Page: "building-new", Pattern: "/buildings/new"},
		{Page: "member-list", Pattern: "/members"},
	}

	sorted := sortBreadcrumbRoutes(trails)
	wantOrder := []string{"building-new", "member-list", "building-detail"}
	for i, page := range wantOrder {
		if sorted[i].Page != page {
			t.Errorf("sorted[%d].Page = %q, want %q (static before param, document order on ties)", i, sorted[i].Page, page)
		}
	}

	wantInput := []breadcrumbTrail{
		{Page: "building-detail", Pattern: "/buildings/:BuildingID"},
		{Page: "building-new", Pattern: "/buildings/new"},
		{Page: "member-list", Pattern: "/members"},
	}
	if !reflect.DeepEqual(trails, wantInput) {
		t.Errorf("input slice mutated: %+v", trails)
	}
}
