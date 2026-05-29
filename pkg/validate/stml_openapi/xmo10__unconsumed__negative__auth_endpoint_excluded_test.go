//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_Unconsumed_Negative_AuthEndpointExcluded

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_Unconsumed_Negative_AuthEndpointExcluded(t *testing.T) {
	// Auth endpoint (security: []) is excluded from XMO-10.
	pages := []stml.PageSpec{{
		FileName: "dashboard.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetDashboard",
		}},
	}}
	// Login op has explicit empty security → auth endpoint → excluded
	loginOp := &openapi3.Operation{
		OperationID: "Login",
		Security:    &openapi3.SecurityRequirements{},
		Responses:   openapi3.NewResponses(),
	}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/dashboard":  getOp("GetDashboard", nil, nil),
		"/auth/login": {Post: loginOp},
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("unexpected XMO-10 for auth endpoint Login, got %v", diags)
	}
}
