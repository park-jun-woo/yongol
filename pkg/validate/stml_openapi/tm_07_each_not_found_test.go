//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM07_EachNotFound_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM07_EachNotFound_Positive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "list.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
			Eaches: []stml.EachBlock{
				{Field: "NonExistent"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, map[string]*openapi3.SchemaRef{
			"users": arrayProp("object"),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[TM-07]") {
		t.Errorf("expected TM-07 diagnostic, got %v", diags)
	}
}
