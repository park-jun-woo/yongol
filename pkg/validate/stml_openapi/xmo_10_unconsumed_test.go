//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_Unconsumed_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_Unconsumed_Positive(t *testing.T) {
	// STML references only ListUsers; GetUser is unconsumed.
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users":      getOp("ListUsers", nil, nil),
		"/users/{id}": getOp("GetUser", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[XMO-10]") {
		t.Errorf("expected XMO-10 diagnostic for GetUser, got %v", diags)
	}
	if countDiag(diags, "[XMO-10]") != 1 {
		t.Errorf("expected exactly 1 XMO-10, got %d", countDiag(diags, "[XMO-10]"))
	}
}
