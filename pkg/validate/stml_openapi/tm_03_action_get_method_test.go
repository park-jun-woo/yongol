//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM03_ActionGetMethod_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM03_ActionGetMethod_Positive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "form.html",
		Actions: []stml.ActionBlock{{
			OperationID: "ListUsers",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[TM-03]") {
		t.Errorf("expected TM-03 diagnostic, got %v", diags)
	}
}
