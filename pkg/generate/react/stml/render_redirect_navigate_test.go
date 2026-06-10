//ff:func feature=stml-gen type=test control=sequence
//ff:what TestRenderRedirectNavigate — 정적 경로 불변·응답 필드 치환·부재 가드·route 소스·optional 생략·패턴 폴백 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderRedirectNavigate(t *testing.T) {
	// Static "/"-prefixed path: verbatim, byte-identical to the
	// pre-Phase008 output, no data dependency.
	lines, usesData := renderRedirectNavigate(stmlparser.ActionBlock{OperationID: "Login", Redirect: "/"})
	if usesData || len(lines) != 1 || lines[0] != "navigate('/')" {
		t.Errorf("static index: got %v (usesData=%v)", lines, usesData)
	}
	lines, usesData = renderRedirectNavigate(stmlparser.ActionBlock{OperationID: "Login", Redirect: "/home"})
	if usesData || len(lines) != 1 || lines[0] != "navigate('/home')" {
		t.Errorf("static path: got %v (usesData=%v)", lines, usesData)
	}

	// Page-name reference with a respField substitution: missing-field
	// guard precedes the template-literal navigate.
	a := stmlparser.ActionBlock{
		OperationID:     "CreateContract",
		Redirect:        "contract-edit",
		RedirectParams:  []stmlparser.LinkParamBind{{Source: "id", Segment: "ContractID"}},
		RedirectPattern: "/contract-edit/:ContractID",
	}
	lines, usesData = renderRedirectNavigate(a)
	want := []string{
		"if (data?.id == null) {",
		"  setCreateContractError('Unexpected response: missing id')",
		"  return",
		"}",
		"navigate(`/contract-edit/${data.id}`)",
	}
	if !usesData || strings.Join(lines, "\n") != strings.Join(want, "\n") {
		t.Errorf("respField: got %v (usesData=%v), want %v", lines, usesData, want)
	}

	// Elided form binds to the single required segment.
	a.RedirectParams = []stmlparser.LinkParamBind{{Source: "id"}}
	lines, usesData = renderRedirectNavigate(a)
	if !usesData || lines[len(lines)-1] != "navigate(`/contract-edit/${data.id}`)" {
		t.Errorf("elided: got %v (usesData=%v)", lines, usesData)
	}

	// route.<Name> source: no guard (the param is already in scope), the
	// useParams variable is interpolated directly.
	b := stmlparser.ActionBlock{
		OperationID:     "UpdateUnit",
		Redirect:        "building-detail",
		RedirectParams:  []stmlparser.LinkParamBind{{Source: "route.BuildingID", Segment: "BuildingID"}},
		RedirectPattern: "/buildings/:BuildingID/:PhotoID?",
	}
	lines, usesData = renderRedirectNavigate(b)
	if usesData || len(lines) != 1 || lines[0] != "navigate(`/buildings/${BuildingID}`)" {
		t.Errorf("route source + optional omitted: got %v (usesData=%v)", lines, usesData)
	}

	// Page-name target without segments: plain string navigate.
	lines, usesData = renderRedirectNavigate(stmlparser.ActionBlock{
		OperationID:     "Logout",
		Redirect:        "dashboard",
		RedirectPattern: "/dashboard",
	})
	if usesData || len(lines) != 1 || lines[0] != "navigate('/dashboard')" {
		t.Errorf("no segments: got %v (usesData=%v)", lines, usesData)
	}

	// Empty RedirectPattern falls back to "/<page-name>".
	lines, usesData = renderRedirectNavigate(stmlparser.ActionBlock{OperationID: "Logout", Redirect: "dashboard"})
	if usesData || len(lines) != 1 || lines[0] != "navigate('/dashboard')" {
		t.Errorf("fallback: got %v (usesData=%v)", lines, usesData)
	}

	// The same respField mapped to two segments is guarded only once.
	c := stmlparser.ActionBlock{
		OperationID: "CreateThing",
		Redirect:    "thing-pair",
		RedirectParams: []stmlparser.LinkParamBind{
			{Source: "id", Segment: "A"},
			{Source: "id", Segment: "B"},
		},
		RedirectPattern: "/thing-pair/:A/:B",
	}
	lines, usesData = renderRedirectNavigate(c)
	wantDedup := []string{
		"if (data?.id == null) {",
		"  setCreateThingError('Unexpected response: missing id')",
		"  return",
		"}",
		"navigate(`/thing-pair/${data.id}/${data.id}`)",
	}
	if !usesData || strings.Join(lines, "\n") != strings.Join(wantDedup, "\n") {
		t.Errorf("dedup guard: got %v (usesData=%v), want %v", lines, usesData, wantDedup)
	}
}
