//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-32 사이트맵 확장 — 구문/route.* 금지/item 스키마/세그먼트 미존재/필수 미충족/생략형 발화·침묵 매트릭스 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM32SitemapGroupLinkParams(t *testing.T) {
	detail := stml.PageSpec{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID/:PhotoID?"}
	settings := stml.PageSpec{Name: "settings", FileName: "settings.html"}
	raif := map[string]map[string]map[string]bool{
		"ListMyBuildings": {"items": {"building_id": true, "building_name": true}},
	}
	sitemapFor := func(link, paramsRaw string) *stml.SitemapSpec {
		return &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "내 건물", Fetch: "ListMyBuildings", Each: "items", Link: link, LinkParamsRaw: paramsRaw, LabelField: "building_name"},
		}}}}
	}
	run := func(link, paramsRaw string) []string {
		fs := makeFS([]stml.PageSpec{detail, settings}, nil)
		fs.Sitemap = sitemapFor(link, paramsRaw)
		var msgs []string
		for _, d := range tm32SitemapGroupLinkParams(fs, raif) {
			msgs = append(msgs, d.Message)
		}
		return msgs
	}

	t.Run("satisfied mapping is silent", func(t *testing.T) {
		if msgs := run("building-detail", "item.building_id -> BuildingID"); len(msgs) != 0 {
			t.Errorf("expected silence, got %v", msgs)
		}
	})

	t.Run("elided form against the single required segment is silent", func(t *testing.T) {
		if msgs := run("building-detail", "item.building_id"); len(msgs) != 0 {
			t.Errorf("expected silence, got %v", msgs)
		}
	})

	t.Run("no params against a target without required segments is silent", func(t *testing.T) {
		if msgs := run("settings", ""); len(msgs) != 0 {
			t.Errorf("expected silence, got %v", msgs)
		}
	})

	t.Run("syntax violation fires once", func(t *testing.T) {
		msgs := run("building-detail", "nonsense format")
		if len(msgs) != 1 || !strings.Contains(msgs[0], "[TM-32]") {
			t.Errorf("msgs = %v", msgs)
		}
	})

	t.Run("route.* source is rejected — no route context in a menu", func(t *testing.T) {
		msgs := run("building-detail", "route.BuildingID -> BuildingID")
		if len(msgs) != 1 || !strings.Contains(msgs[0], "route segment") {
			t.Errorf("msgs = %v", msgs)
		}
	})

	t.Run("item field missing from the item schema fires", func(t *testing.T) {
		msgs := run("building-detail", "item.nope -> BuildingID")
		if len(msgs) != 1 || !strings.Contains(msgs[0], "not in the item schema") {
			t.Errorf("msgs = %v", msgs)
		}
	})

	t.Run("unknown segment name fires", func(t *testing.T) {
		msgs := run("building-detail", "item.building_id -> Nope")
		// the unknown segment plus the now-unmapped required :BuildingID
		if len(msgs) != 2 || !strings.Contains(msgs[0], "not in target page") {
			t.Errorf("msgs = %v", msgs)
		}
	})

	t.Run("unmapped required segment fires", func(t *testing.T) {
		msgs := run("building-detail", "")
		if len(msgs) != 1 || !strings.Contains(msgs[0], "required segment :BuildingID") {
			t.Errorf("msgs = %v", msgs)
		}
	})

	t.Run("unknown target page stays silent (TM-31 owns it)", func(t *testing.T) {
		if msgs := run("nope-detail", "item.building_id -> BuildingID"); len(msgs) != 0 {
			t.Errorf("expected silence, got %v", msgs)
		}
	})
}
