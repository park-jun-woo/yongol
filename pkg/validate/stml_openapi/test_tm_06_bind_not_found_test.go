//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM06_BindNotFound_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM06_BindNotFound_Positive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "profile.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetProfile",
			Binds: []stml.FieldBind{
				{Name: "NonExistent"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/profile": getOp("GetProfile", nil, map[string]*openapi3.SchemaRef{
			"Name":  stringProp(),
			"Email": stringProp(),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[TM-06]") {
		t.Errorf("expected TM-06 diagnostic, got %v", diags)
	}
}
