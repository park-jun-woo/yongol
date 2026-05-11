//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-11/12/13 test — 레이아웃 참조 정합성 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// ---------- TM-11 ----------

func TestTM11_LayoutNotFound_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "dashboard-page",
		FileName: "dashboard-page.html",
		Layout:   "main",
	}}
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm11LayoutNotFound(pages, layouts)
	if !hasDiag(diags, "[TM-11]") {
		t.Errorf("expected TM-11 for missing layout 'main', got %v", diags)
	}
}

func TestTM11_LayoutFound_Pass(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "dashboard-page",
		FileName: "dashboard-page.html",
		Layout:   "app",
	}}
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm11LayoutNotFound(pages, layouts)
	if hasDiag(diags, "[TM-11]") {
		t.Errorf("unexpected TM-11 for valid layout reference, got %v", diags)
	}
}

func TestTM11_EmptyLayout_Skip(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "login-page",
		FileName: "login-page.html",
		Layout:   "",
	}}
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm11LayoutNotFound(pages, layouts)
	if hasDiag(diags, "[TM-11]") {
		t.Errorf("unexpected TM-11 for page with empty layout, got %v", diags)
	}
}

func TestTM11_MultiplePages_CountsAll(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "page-a", FileName: "page-a.html", Layout: "missing1"},
		{Name: "page-b", FileName: "page-b.html", Layout: "missing2"},
		{Name: "page-c", FileName: "page-c.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm11LayoutNotFound(pages, layouts)
	count := countDiag(diags, "[TM-11]")
	if count != 2 {
		t.Errorf("expected 2 TM-11 diagnostics, got %d: %+v", count, diags)
	}
}

// ---------- TM-12 ----------

func TestTM12_DefaultLayoutNotFound_Error(t *testing.T) {
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm12DefaultLayoutNotFound("main", layouts)
	if !hasDiag(diags, "[TM-12]") {
		t.Errorf("expected TM-12 for missing defaultLayout 'main', got %v", diags)
	}
}

func TestTM12_DefaultLayoutFound_Pass(t *testing.T) {
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm12DefaultLayoutNotFound("app", layouts)
	if hasDiag(diags, "[TM-12]") {
		t.Errorf("unexpected TM-12 for valid defaultLayout, got %v", diags)
	}
}

func TestTM12_EmptyDefaultLayout_Skip(t *testing.T) {
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm12DefaultLayoutNotFound("", layouts)
	if hasDiag(diags, "[TM-12]") {
		t.Errorf("unexpected TM-12 for empty defaultLayout, got %v", diags)
	}
}

// ---------- TM-13 ----------

func TestTM13_UnusedLayout_Warning(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "dashboard-page",
		FileName: "dashboard-page.html",
		Layout:   "app",
	}}
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html"},
		{Name: "auth", File: "layouts/auth.html"},
	}
	diags := tm13UnusedLayout(pages, layouts, "app")
	if !hasDiag(diags, "[TM-13]") {
		t.Errorf("expected TM-13 for unused layout 'auth', got %v", diags)
	}
	count := countDiag(diags, "[TM-13]")
	if count != 1 {
		t.Errorf("expected 1 TM-13 diagnostic, got %d: %+v", count, diags)
	}
}

func TestTM13_AllUsed_Pass(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard-page", FileName: "dashboard-page.html", Layout: "app"},
		{Name: "login-page", FileName: "login-page.html", Layout: "auth"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html"},
		{Name: "auth", File: "layouts/auth.html"},
	}
	diags := tm13UnusedLayout(pages, layouts, "")
	if hasDiag(diags, "[TM-13]") {
		t.Errorf("unexpected TM-13 when all layouts are used, got %v", diags)
	}
}

func TestTM13_DefaultLayoutCountsAsUsed(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "login-page",
		FileName: "login-page.html",
		Layout:   "auth",
	}}
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html"},
		{Name: "auth", File: "layouts/auth.html"},
	}
	// "app" is only used as defaultLayout, not by any page directly
	diags := tm13UnusedLayout(pages, layouts, "app")
	if hasDiag(diags, "[TM-13]") {
		t.Errorf("unexpected TM-13 when layout is defaultLayout, got %v", diags)
	}
}

func TestTM13_NoPages_AllUnused(t *testing.T) {
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html"},
		{Name: "auth", File: "layouts/auth.html"},
	}
	diags := tm13UnusedLayout(nil, layouts, "")
	count := countDiag(diags, "[TM-13]")
	if count != 2 {
		t.Errorf("expected 2 TM-13 diagnostics, got %d: %+v", count, diags)
	}
}

// ---------- run.go integration: skip when no layouts ----------

func TestRun_NoLayouts_SkipLayoutRules(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/items": getOp("ListItems", nil, nil),
	})
	fs := makeFS([]stml.PageSpec{{
		Name:        "page",
		FileName:    "page.html",
		Layout:      "nonexistent",
		Fetches:     []stml.FetchBlock{{OperationID: "ListItems"}},
	}}, doc)
	// fs.Layouts is nil, fs.Manifest is nil → skip TM-11/12/13
	diags := Run(fs)
	if hasDiag(diags, "[TM-11]") || hasDiag(diags, "[TM-12]") || hasDiag(diags, "[TM-13]") {
		t.Errorf("expected no TM-11/12/13 when Layouts is nil and no manifest, got %v", diags)
	}
}
