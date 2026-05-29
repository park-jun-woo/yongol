//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM02_ActionOpNotFound_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM02_ActionOpNotFound_Positive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "form.html",
		Actions: []stml.ActionBlock{{
			OperationID: "NoSuchAction",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": postOp("CreateUser", nil),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[TM-02]") {
		t.Errorf("expected TM-02 diagnostic, got %v", diags)
	}
}
