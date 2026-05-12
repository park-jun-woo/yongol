//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM06_BindDottedPath

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM06_BindDottedPath(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "profile.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetProfile",
			Binds: []stml.FieldBind{
				{Name: "User.Name"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/profile": getOp("GetProfile", nil, map[string]*openapi3.SchemaRef{
			"User": stringProp(),
		}),
	})
	diags := Run(makeFS(pages, doc))
	// "User.Name" should resolve to "User" top-level key
	if hasDiag(diags, "[TM-06]") {
		t.Errorf("dotted path should resolve to top-level key, got %v", diags)
	}
}
