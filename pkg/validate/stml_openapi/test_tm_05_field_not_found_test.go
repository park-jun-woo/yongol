//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM05_FieldNotFound_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM05_FieldNotFound_Positive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "form.html",
		Actions: []stml.ActionBlock{{
			OperationID: "CreateUser",
			Fields: []stml.FieldBind{
				{Name: "Email"},
				{Name: "NonExistent"},
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
	if !hasDiag(diags, "[TM-05]") {
		t.Errorf("expected TM-05 diagnostic for NonExistent, got %v", diags)
	}
	if countDiag(diags, "[TM-05]") != 1 {
		t.Errorf("expected exactly 1 TM-05 diagnostic, got %d", countDiag(diags, "[TM-05]"))
	}
}
