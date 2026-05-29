//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM05_FieldNotFound_Negative

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM05_FieldNotFound_Negative(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "form.html",
		Actions: []stml.ActionBlock{{
			OperationID: "CreateUser",
			Fields: []stml.FieldBind{
				{Name: "Email"},
				{Name: "Name"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": postOp("CreateUser", map[string]*openapi3.SchemaRef{
			"Email": stringProp(),
			"Name":  stringProp(),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[TM-05]") {
		t.Errorf("unexpected TM-05 diagnostic, got %v", diags)
	}
}
