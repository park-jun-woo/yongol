//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_Unconsumed_Negative_AllConsumed

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_Unconsumed_Negative_AllConsumed(t *testing.T) {
	// STML references both ops — no XMO-10 expected.
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
		}},
		Actions: []stml.ActionBlock{{
			OperationID: "CreateUser",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, nil),
		"/users/create": postOp("CreateUser", map[string]*openapi3.SchemaRef{
			"name": stringProp(),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("unexpected XMO-10 diagnostic, got %v", diags)
	}
}
