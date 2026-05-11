//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-10 test — STML 요소에 class 속성 직접 사용 시 ERROR 검출

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM10_StaticElementWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Children: []stml.ChildNode{{
			Kind:   "static",
			Static: &stml.StaticElement{Tag: "div", ClassName: "bg-red-500"},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 diagnostic for class on StaticElement, got %v", diags)
	}
}

func TestTM10_NoClass_Pass(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Children: []stml.ChildNode{{
			Kind:   "static",
			Static: &stml.StaticElement{Tag: "div"},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if hasDiag(diags, "[TM-10]") {
		t.Errorf("unexpected TM-10 diagnostic for element without class, got %v", diags)
	}
}

func TestTM10_FetchBlockWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListItems",
			ClassName:   "p-4",
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 for FetchBlock class, got %v", diags)
	}
}

func TestTM10_ActionBlockWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Actions: []stml.ActionBlock{{
			OperationID: "CreateItem",
			ClassName:   "mt-2",
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 for ActionBlock class, got %v", diags)
	}
}

func TestTM10_EachBlockWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListItems",
			Eaches: []stml.EachBlock{{
				Field:     "items",
				ClassName: "grid",
			}},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 for EachBlock class, got %v", diags)
	}
}

func TestTM10_FieldBindWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListItems",
			Binds: []stml.FieldBind{{
				Name:      "Title",
				ClassName: "font-bold",
			}},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 for FieldBind class, got %v", diags)
	}
}

func TestTM10_ComponentRefWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListItems",
			Components: []stml.ComponentRef{{
				Name:      "DatePicker",
				ClassName: "w-full",
			}},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 for ComponentRef class, got %v", diags)
	}
}

func TestTM10_StateBindWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListItems",
			States: []stml.StateBind{{
				Condition: "items.empty",
				ClassName: "text-gray-500",
			}},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 for StateBind class, got %v", diags)
	}
}

func TestTM10_MultipleViolations_CountsAll(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Children: []stml.ChildNode{
			{Kind: "static", Static: &stml.StaticElement{Tag: "div", ClassName: "bg-red"}},
			{Kind: "static", Static: &stml.StaticElement{Tag: "span", ClassName: "text-lg"}},
			{Kind: "static", Static: &stml.StaticElement{Tag: "p"}},
		},
	}}
	diags := tm10ClassProhibited(pages)
	count := countDiag(diags, "[TM-10]")
	if count != 2 {
		t.Errorf("expected 2 TM-10 diagnostics, got %d: %+v", count, diags)
	}
}
