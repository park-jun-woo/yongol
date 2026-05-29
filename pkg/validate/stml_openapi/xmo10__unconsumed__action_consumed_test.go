//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_Unconsumed_ActionConsumed

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_Unconsumed_ActionConsumed(t *testing.T) {
	// data-action also counts as consumption.
	pages := []stml.PageSpec{{
		FileName: "form.html",
		Actions: []stml.ActionBlock{{
			OperationID: "CreateWorkflow",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/workflows": postOp("CreateWorkflow", map[string]*openapi3.SchemaRef{
			"name": stringProp(),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("unexpected XMO-10 for consumed action, got %v", diags)
	}
}
