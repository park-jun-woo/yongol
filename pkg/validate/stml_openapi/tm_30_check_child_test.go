//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm30CheckChild — action/fetch/each/static/state/nil/기타 kind 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM30CheckChild(t *testing.T) {
	raif := map[string]map[string]map[string]bool{
		"GetUnit": {"photos": {"id": true}},
	}
	itemAction := func() *stml.ActionBlock {
		return &stml.ActionBlock{
			OperationID: "DeletePhoto",
			Params:      []stml.ParamBind{{Name: "photoId", Source: "item.id"}},
		}
	}

	// each kind resolves the item schema → valid field stays silent
	var out []diagnostic.Diagnostic
	each := stml.EachBlock{Field: "photos", Children: []stml.ChildNode{{Kind: "action", Action: itemAction()}}}
	tm30CheckChild(stml.ChildNode{Kind: "each", Each: &each}, "p.html", "GetUnit", nil, false, raif, &out)
	if len(out) != 0 {
		t.Errorf("each with valid field: %v", out)
	}

	// each under an unknown op → unresolved schema, silent
	out = nil
	tm30CheckChild(stml.ChildNode{Kind: "each", Each: &each}, "p.html", "Unknown", nil, false, raif, &out)
	if len(out) != 0 {
		t.Errorf("unresolved schema must stay silent: %v", out)
	}

	// action outside each → error
	out = nil
	tm30CheckChild(stml.ChildNode{Kind: "action", Action: itemAction()}, "p.html", "", nil, false, raif, &out)
	if len(out) != 1 {
		t.Errorf("outside-each action: %v", out)
	}

	// fetch kind: item.* fetch param flagged, children rechecked with its opID
	out = nil
	fetch := stml.FetchBlock{
		OperationID: "GetUnit",
		Params:      []stml.ParamBind{{Name: "photoId", Source: "item.id"}},
		Children:    []stml.ChildNode{{Kind: "each", Each: &each}},
	}
	tm30CheckChild(stml.ChildNode{Kind: "fetch", Fetch: &fetch}, "p.html", "", nil, false, raif, &out)
	if len(out) != 1 {
		t.Errorf("fetch param + valid each: %v", out)
	}

	// static / state recursion keeps the context; nils are tolerated
	out = nil
	st := stml.StaticElement{Children: []stml.ChildNode{{Kind: "action", Action: itemAction()}}}
	tm30CheckChild(stml.ChildNode{Kind: "static", Static: &st}, "p.html", "GetUnit", map[string]bool{"id": true}, true, raif, &out)
	if len(out) != 0 {
		t.Errorf("static inside each: %v", out)
	}
	sb := stml.StateBind{Children: []stml.ChildNode{{Kind: "action", Action: itemAction()}}}
	tm30CheckChild(stml.ChildNode{Kind: "state", State: &sb}, "p.html", "GetUnit", nil, false, raif, &out)
	if len(out) != 1 {
		t.Errorf("state outside each: %v", out)
	}
	out = nil
	tm30CheckChild(stml.ChildNode{Kind: "static"}, "p.html", "", nil, false, raif, &out)
	tm30CheckChild(stml.ChildNode{Kind: "state"}, "p.html", "", nil, false, raif, &out)
	tm30CheckChild(stml.ChildNode{Kind: "bind", Bind: &stml.FieldBind{Name: "x"}}, "p.html", "", nil, false, raif, &out)
	if len(out) != 0 {
		t.Errorf("nil/bind kinds must stay silent: %v", out)
	}
}
