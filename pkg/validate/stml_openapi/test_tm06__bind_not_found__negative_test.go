//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM06_BindNotFound_Negative

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM06_BindNotFound_Negative(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "profile.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetProfile",
			Binds: []stml.FieldBind{
				{Name: "Name"},
				{Name: "Email"},
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
	if hasDiag(diags, "[TM-06]") {
		t.Errorf("unexpected TM-06 diagnostic, got %v", diags)
	}
}
