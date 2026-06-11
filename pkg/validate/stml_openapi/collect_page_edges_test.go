//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what collectPageEdges — data-link 대상 + 페이지 액션·each 행 액션의 data-redirect 대상 수집 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectPageEdges(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "building-list", FileName: "building-list.html"},
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"},
		{Name: "building-create", FileName: "building-create.html"},
	}
	page := stml.PageSpec{
		Name:     "building-list",
		FileName: "building-list.html",
		Actions:  []stml.ActionBlock{{OperationID: "CreateBuilding", Redirect: "building-create"}},
		Children: []stml.ChildNode{
			{Kind: "each", Each: &stml.EachBlock{
				Field:   "items",
				RowLink: &stml.LinkRef{TargetPage: "building-detail"},
				Actions: []stml.ActionBlock{{OperationID: "DeleteBuilding", Redirect: "building-list"}},
			}},
		},
	}
	got := collectPageEdges(page, pages, nil)
	want := []string{"building-detail", "building-create", "building-list"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("edges = %v, want %v", got, want)
	}
}
