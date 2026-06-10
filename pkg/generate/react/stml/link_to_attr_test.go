//ff:func feature=stml-gen type=test control=sequence
//ff:what TestLinkToAttr — to 속성 생성(정적·치환·optional 생략·생략형·폴백) 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLinkToAttr(t *testing.T) {
	// No params → plain string attribute.
	got := linkToAttr(stmlparser.LinkRef{TargetPage: "settings", TargetPattern: "/settings"})
	if got != `to="/settings"` {
		t.Errorf("static: got %q", got)
	}

	// item.* source substituted into a template literal; the unmapped
	// optional segment is omitted.
	got = linkToAttr(stmlparser.LinkRef{
		TargetPage:    "building-detail",
		TargetPattern: "/buildings/:BuildingID/:PhotoID?",
		Params:        []stmlparser.LinkParamBind{{Source: "item.id", Segment: "BuildingID"}},
	})
	if got != "to={`/buildings/${item.id}`}" {
		t.Errorf("item source: got %q", got)
	}

	// route.* source uses the useParams() variable.
	got = linkToAttr(stmlparser.LinkRef{
		TargetPage:    "unit-list",
		TargetPattern: "/unit-list/:BuildingID",
		Params:        []stmlparser.LinkParamBind{{Source: "route.BuildingID", Segment: "BuildingID"}},
	})
	if got != "to={`/unit-list/${BuildingID}`}" {
		t.Errorf("route source: got %q", got)
	}

	// Elided segment binds to the single required segment.
	got = linkToAttr(stmlparser.LinkRef{
		TargetPage:    "building-detail",
		TargetPattern: "/buildings/:BuildingID/:PhotoID?",
		Params:        []stmlparser.LinkParamBind{{Source: "item.id"}},
	})
	if got != "to={`/buildings/${item.id}`}" {
		t.Errorf("elided: got %q", got)
	}

	// Mapped optional segment is filled.
	got = linkToAttr(stmlparser.LinkRef{
		TargetPage:    "building-detail",
		TargetPattern: "/buildings/:BuildingID/:PhotoID?",
		Params: []stmlparser.LinkParamBind{
			{Source: "item.id", Segment: "BuildingID"},
			{Source: "item.photo_id", Segment: "PhotoID"},
		},
	})
	if got != "to={`/buildings/${item.id}/${item.photo_id}`}" {
		t.Errorf("optional mapped: got %q", got)
	}

	// Empty pattern falls back to "/<page-name>".
	got = linkToAttr(stmlparser.LinkRef{TargetPage: "login"})
	if got != `to="/login"` {
		t.Errorf("fallback: got %q", got)
	}

	// The index route ("/") keeps the bare root path.
	got = linkToAttr(stmlparser.LinkRef{TargetPage: "dashboard", TargetPattern: "/"})
	if got != `to="/"` {
		t.Errorf("index: got %q", got)
	}
}
