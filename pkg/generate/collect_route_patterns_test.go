//ff:func feature=generate type=test control=sequence
//ff:what TestCollectRoutePatterns — 페이지명 → 해석 라우트 패턴 맵 구성 검증

package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectRoutePatterns(t *testing.T) {
	pages := []stmlparser.PageSpec{
		{Name: "building-list", FileName: "building-list.html"},
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"},
	}
	got := collectRoutePatterns(pages)
	if got["building-list"] != "/building-list" {
		t.Errorf("building-list = %q", got["building-list"])
	}
	if got["building-detail"] != "/buildings/:BuildingID" {
		t.Errorf("building-detail = %q", got["building-detail"])
	}
	if len(collectRoutePatterns(nil)) != 0 {
		t.Error("nil pages must yield an empty map")
	}
}
