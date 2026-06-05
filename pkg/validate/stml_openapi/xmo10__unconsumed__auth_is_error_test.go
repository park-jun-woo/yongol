//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_AuthUnconsumed_IsError — auth(security:[]) 미소비도 ERROR (자동제외 폐지)

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_AuthUnconsumed_IsError(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "dashboard.html",
		Fetches:  []stml.FetchBlock{{OperationID: "GetDashboard"}},
	}}
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
	if countDiag(diags, "[XMO-10]") != 1 {
		t.Fatalf("expected XMO-10 for unconsumed auth op Login, got %v", diags)
	}
}
