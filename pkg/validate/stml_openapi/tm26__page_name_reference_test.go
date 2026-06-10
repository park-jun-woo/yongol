//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-26 페이지명 참조 분기 — 존재 페이지 침묵·미존재 페이지명 ERROR 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM26PageNameReference(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "contract-edit", FileName: "contract-edit.html"},
	}

	// Page-name reference to an existing page → silent.
	if d := tm26RedirectRouteExists(stml.ActionBlock{OperationID: "CreateContract", Redirect: "contract-edit"}, "contract-new.html", pages); len(d) != 0 {
		t.Errorf("existing page name: expected 0 diagnostics, got %+v", d)
	}

	// Typo'd page name → 1 ERROR.
	got := tm26RedirectRouteExists(stml.ActionBlock{OperationID: "CreateContract", Redirect: "contract-edt"}, "contract-new.html", pages)
	if countDiag(got, "[TM-26]") != 1 {
		t.Errorf("typo page name: expected 1 TM-26, got %+v", got)
	}
}
