//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestRun_NoLayouts_SkipLayoutRules

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

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
