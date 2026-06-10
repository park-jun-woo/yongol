//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm36StaticNavMatches — 정적 경로 ↔ 해석 라우트 패턴 매칭 판정 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM36StaticNavMatches(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "building-detail", FileName: "building-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "building_id", Source: "route.BuildingID"}},
			}}},
	}

	cases := []struct {
		name, path string
		want       bool
	}{
		{"exact base path", "/dashboard", true},
		{"param segment filled", "/buildings/3", true},
		{"no match", "/nowhere", false},
		{"param segment unfilled", "/buildings", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := tm36StaticNavMatches(c.path, pages); got != c.want {
				t.Errorf("tm36StaticNavMatches(%q) = %v, want %v", c.path, got, c.want)
			}
		})
	}
}
