//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM03_ActionGetMethod_Negative

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM03_ActionGetMethod_Negative(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "form.html",
		Actions: []stml.ActionBlock{{
			OperationID: "CreateUser",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": postOp("CreateUser", nil),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[TM-03]") {
		t.Errorf("unexpected TM-03 diagnostic, got %v", diags)
	}
}
