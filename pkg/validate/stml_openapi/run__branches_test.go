//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestRun_Branches — Run early-return / Actions / layouts+manifest 분기 커버

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRun_EarlyReturn(t *testing.T) {
	// nil OpenAPIDoc → nil.
	if got := Run(makeFS([]stml.PageSpec{{Name: "p"}}, nil)); got != nil {
		t.Errorf("nil doc: expected nil, got %v", got)
	}
	// no STML pages → nil.
	doc := makeDoc(map[string]*openapi3.PathItem{"/items": getOp("ListItems", nil, nil)})
	if got := Run(makeFS(nil, doc)); got != nil {
		t.Errorf("no pages: expected nil, got %v", got)
	}
}

func TestRun_ActionsAndLayouts(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/items":  getOp("ListItems", nil, nil),
		"/create": postOp("CreateItem", nil),
	})
	fs := makeFS([]stml.PageSpec{{
		Name:     "page",
		FileName: "page.html",
		Layout:   "app",
		Fetches:  []stml.FetchBlock{{OperationID: "ListItems"}},
		Actions:  []stml.ActionBlock{{OperationID: "CreateItem"}},
	}}, doc)
	// Provide layouts + manifest default layout → exercise TM-11/12/13 branch.
	fs.Layouts = []stml.LayoutSpec{{Name: "app", File: "layouts/app.html", HasOutlet: true}}
	fs.Manifest = &manifest.ProjectConfig{}
	fs.Manifest.Frontend.DefaultLayout = "app"

	diags := Run(fs)
	// All operations are consumed & layout exists → no layout-not-found errors.
	if hasDiag(diags, "[TM-11]") || hasDiag(diags, "[TM-12]") {
		t.Errorf("expected no TM-11/12 for existing layout, got %v", diags)
	}
}
