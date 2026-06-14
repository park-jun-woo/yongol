//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what markChildActionBindsOptional — 중첩 트리(action/fetch/state/static/each) ActionBlock 바인드 optional 표시 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestMarkChildActionBindsOptional(t *testing.T) {
	// build a tree with action blocks nested under every recursive container kind
	topAction := &stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "BuildingID", Source: "route.BuildingID"}, // required → stays false
		{Name: "RoomID", Source: "route.RoomID"},         // not required → optional
		{Name: "Q", Source: "query.q"},                   // non-route → untouched
	}}
	fetchRowAction := &stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "RoomID", Source: "route.RoomID"},
	}}
	eachRowAction := &stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "RoomID", Source: "route.RoomID"},
	}}
	stateAction := &stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "RoomID", Source: "route.RoomID"},
	}}
	staticAction := &stmlparser.ActionBlock{Params: []stmlparser.ParamBind{
		{Name: "RoomID", Source: "route.RoomID"},
	}}

	children := []stmlparser.ChildNode{
		{Kind: "action", Action: topAction},
		{Kind: "fetch", Fetch: &stmlparser.FetchBlock{Children: []stmlparser.ChildNode{
			{Kind: "each", Each: &stmlparser.EachBlock{Children: []stmlparser.ChildNode{
				{Kind: "action", Action: eachRowAction},
			}}},
			{Kind: "action", Action: fetchRowAction},
		}}},
		{Kind: "state", State: &stmlparser.StateBind{Children: []stmlparser.ChildNode{
			{Kind: "action", Action: stateAction},
		}}},
		{Kind: "static", Static: &stmlparser.StaticElement{Children: []stmlparser.ChildNode{
			{Kind: "action", Action: staticAction},
		}}},
		{Kind: "bind"}, // non-recursive kind, ignored
	}

	required := map[string]bool{"BuildingID": true}
	markChildActionBindsOptional(children, required)

	// required route param stays non-optional
	if topAction.Params[0].Optional {
		t.Errorf("required BuildingID should stay Optional=false")
	}
	// unrequired route param flagged optional at every depth
	if !topAction.Params[1].Optional {
		t.Errorf("top RoomID should be Optional=true")
	}
	// non-route source untouched
	if topAction.Params[2].Optional {
		t.Errorf("query.q should stay Optional=false")
	}
	for _, a := range []*stmlparser.ActionBlock{fetchRowAction, eachRowAction, stateAction, staticAction} {
		if !a.Params[0].Optional {
			t.Errorf("nested RoomID should be Optional=true, got false")
		}
	}
}
