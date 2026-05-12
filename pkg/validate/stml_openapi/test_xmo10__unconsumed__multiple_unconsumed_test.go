//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_Unconsumed_MultipleUnconsumed

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_Unconsumed_MultipleUnconsumed(t *testing.T) {
	// Multiple unconsumed ops → multiple XMO-10 diagnostics.
	pages := []stml.PageSpec{{
		FileName: "home.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetDashboard",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/dashboard": getOp("GetDashboard", nil, nil),
		"/users":     getOp("ListUsers", nil, nil),
		"/actions":   getOp("ListActions", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if countDiag(diags, "[XMO-10]") != 2 {
		t.Errorf("expected 2 XMO-10 diagnostics, got %d", countDiag(diags, "[XMO-10]"))
	}
}
