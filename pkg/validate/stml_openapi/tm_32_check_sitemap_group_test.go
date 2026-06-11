//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm32CheckSitemapGroup — 미존재 대상 침묵·구문 단독 보고·충족 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM32CheckSitemapGroup(t *testing.T) {
	patterns := map[string]string{"building-detail": "/buildings/:BuildingID"}
	raif := map[string]map[string]map[string]bool{
		"ListMyBuildings": {"items": {"building_id": true}},
	}
	entry := func(link, raw string) sitemapEntry {
		return sitemapEntry{
			Node: stml.SitemapNode{Label: "내 건물", Fetch: "ListMyBuildings", Each: "items", Link: link, LinkParamsRaw: raw, LabelField: "name"},
			Path: "nav[0] > 내 건물",
		}
	}

	t.Run("unknown target is silent", func(t *testing.T) {
		if diags := tm32CheckSitemapGroup(entry("nope", "item.building_id"), "sitemap.html", patterns, raif); len(diags) != 0 {
			t.Errorf("diags = %+v", diags)
		}
	})

	t.Run("a syntax error is reported alone", func(t *testing.T) {
		diags := tm32CheckSitemapGroup(entry("building-detail", "broken raw"), "sitemap.html", patterns, raif)
		if len(diags) != 1 || !strings.Contains(diags[0].Message, "data-link-params") {
			t.Errorf("diags = %+v", diags)
		}
	})

	t.Run("a satisfied mapping is silent", func(t *testing.T) {
		if diags := tm32CheckSitemapGroup(entry("building-detail", "item.building_id -> BuildingID"), "sitemap.html", patterns, raif); len(diags) != 0 {
			t.Errorf("diags = %+v", diags)
		}
	})
}
