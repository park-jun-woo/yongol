//ff:func feature=gen-react type=test control=sequence
//ff:what TestSitemapDynamicItemKey — 첫 item.* 소스 선택·route.* 무시·소스 없음 빈 문자열 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapDynamicItemKey(t *testing.T) {
	if got := sitemapDynamicItemKey(nil); got != "" {
		t.Errorf("no params key = %q, want empty (index fallback)", got)
	}
	params := []stml.LinkParamBind{
		{Source: "item.building_id", Segment: "BuildingID"},
		{Source: "item.unit_id", Segment: "UnitID"},
	}
	if got := sitemapDynamicItemKey(params); got != "item.building_id" {
		t.Errorf("key = %q, want the first item.* source", got)
	}
}
