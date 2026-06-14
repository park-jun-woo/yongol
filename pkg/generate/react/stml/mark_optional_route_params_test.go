//ff:func feature=stml-gen type=test control=sequence
//ff:what markOptionalRouteParams — fetch 소비 route param은 required, action-only는 optional 표시 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestMarkOptionalRouteParams(t *testing.T) {
	pageAction := &stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "BuildingID", Source: "route.BuildingID"}, // consumed by fetch → required
		{Name: "RoomID", Source: "route.RoomID"},         // action-only → optional
	}}
	childAction := &stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "RoomID", Source: "route.RoomID"},
	}}

	page := &stmlparser.PageSpec{
		Fetches: []stmlparser.FetchBlock{
			{OperationID: "GetBuilding", Params: []stmlparser.ParamBind{
				{Name: "BuildingID", Source: "route.BuildingID"},
			}},
		},
		Actions: []stmlparser.ActionBlock{*pageAction},
		Children: []stmlparser.ChildNode{
			{Kind: "action", Action: childAction},
		},
	}

	markOptionalRouteParams(page)

	// fetch-consumed BuildingID stays required (Optional=false)
	if page.Actions[0].Params[0].Optional {
		t.Errorf("fetch-consumed BuildingID should stay Optional=false")
	}
	// action-only RoomID becomes optional
	if !page.Actions[0].Params[1].Optional {
		t.Errorf("action-only RoomID should be Optional=true")
	}
	// nested child action's route param flagged too
	if !childAction.Params[0].Optional {
		t.Errorf("child action RoomID should be Optional=true")
	}
}
