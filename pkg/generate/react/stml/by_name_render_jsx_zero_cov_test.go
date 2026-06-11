//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNameRenderJSX_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	noBody := map[string]bool{}
	ppt := map[string]map[string]string{}
	f := page.Fetches[0]
	a := page.Actions[0]

	if s := renderFetchJSX(f, 0, noBody, bindCtx{}); s == "" {
		t.Errorf("renderFetchJSX empty")
	}
	_ = renderFetchJSXBody(f, "data", 1, noBody, bindCtx{})
	if s := renderActionJSX(a, 0, noBody); s == "" {
		t.Errorf("renderActionJSX empty")
	}
	if s := renderActionForm(a, 1); s == "" {
		t.Errorf("renderActionForm empty")
	}
	if s := renderActionButton(a, 1, noBody); s == "" {
		t.Errorf("renderActionButton empty")
	}
	_ = renderActionChildNodes(a.Children, "form", "createReservation", "", 2)
	_ = renderChildNodes(f.Children, "data", "item", 2, noBody, bindCtx{})

	if len(f.Eaches) > 0 {
		if s := renderEachJSX(f.Eaches[0], "data", 1, noBody, bindCtx{}); s == "" {
			t.Errorf("renderEachJSX empty")
		}
	}
	if len(a.Fields) > 0 {
		if s := renderFieldJSX(a.Fields[0], "form", "createReservation", 1); s == "" {
			t.Errorf("renderFieldJSX empty")
		}
	}
	if len(f.States) > 0 {
		if s := renderStateJSX(f.States[0], "data", 1, noBody, bindCtx{}); s == "" {
			t.Errorf("renderStateJSX empty")
		}
	}
	// static element
	se := stmlparser.StaticElement{Tag: "h2", Text: "Header"}
	if s := renderStaticJSX(se, "data", "item", 1, noBody, bindCtx{}); s == "" {
		t.Errorf("renderStaticJSX empty")
	}

	_ = renderParamArgs(f.Params, "ListItems", ppt)
	_ = renderParamValues(f.Params)
	_ = renderInvalidateExpr([]string{"ListItems", "ListSub"})
	_ = renderInvalidateExpr(nil)
}
