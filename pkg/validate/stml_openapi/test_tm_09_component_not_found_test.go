//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM09_ComponentNotFound_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM09_ComponentNotFound_Positive(t *testing.T) {
	tmpDir := t.TempDir()
	pages := []stml.PageSpec{{
		FileName: "dashboard.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
			Components: []stml.ComponentRef{
				{Name: "NonExistentComponent"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, nil),
	})
	fs := makeFS(pages, doc)
	fs.SpecsDir = tmpDir
	diags := Run(fs)
	if !hasDiag(diags, "[TM-09]") {
		t.Errorf("expected TM-09 diagnostic, got %v", diags)
	}
}
