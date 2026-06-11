//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm32SitemapParam — route.* 즉시 거부(명시 세그먼트는 mapped 유지)·item.* 위임 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM32SitemapParam(t *testing.T) {
	ref := linkRefCtx{
		Link:       &stml.LinkRef{TargetPage: "building-detail"},
		ItemFields: map[string]bool{"building_id": true},
		InEach:     true,
	}
	pattern, required := "/buildings/:BuildingID", []string{"BuildingID"}

	t.Run("route.* source is rejected, its explicit segment stays mapped", func(t *testing.T) {
		seg, diags := tm32SitemapParam(stml.LinkParamBind{Source: "route.BuildingID", Segment: "BuildingID"}, ref, "nav[0] > 내 건물", "sitemap.html", pattern, required)
		if seg != "BuildingID" || len(diags) != 1 || !strings.Contains(diags[0].Message, "route segment") {
			t.Errorf("seg = %q, diags = %+v", seg, diags)
		}
	})

	t.Run("item.* source delegates to the page judgment", func(t *testing.T) {
		seg, diags := tm32SitemapParam(stml.LinkParamBind{Source: "item.building_id", Segment: "BuildingID"}, ref, "nav[0] > 내 건물", "sitemap.html", pattern, required)
		if seg != "BuildingID" || len(diags) != 0 {
			t.Errorf("seg = %q, diags = %+v", seg, diags)
		}
	})
}
