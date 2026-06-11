//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestDynamicMenuGroup — 동적 그룹 완전성 판정 (4종 전부=true, 누락·정적 노드=false) 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestDynamicMenuGroup(t *testing.T) {
	complete := stml.SitemapNode{Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LabelField: "building_name"}
	if !DynamicMenuGroup(complete) {
		t.Errorf("complete vocabulary set must judge true")
	}
	cases := map[string]stml.SitemapNode{
		"static node":        {Page: "dashboard", Label: "대시보드"},
		"missing fetch":      {Each: "items", Link: "building-detail", LabelField: "building_name"},
		"missing each":       {Fetch: "ListMyBuildings", Link: "building-detail", LabelField: "building_name"},
		"missing link":       {Fetch: "ListMyBuildings", Each: "items", LabelField: "building_name"},
		"missing labelfield": {Fetch: "ListMyBuildings", Each: "items", Link: "building-detail"},
	}
	for name, n := range cases {
		if DynamicMenuGroup(n) {
			t.Errorf("%s must judge false", name)
		}
	}
	// data-link-params is optional — required only by the target route (TM-32).
	withParams := complete
	withParams.LinkParamsRaw = "item.building_id -> BuildingID"
	if !DynamicMenuGroup(withParams) {
		t.Errorf("link-params must not change the completeness judgment")
	}
}
