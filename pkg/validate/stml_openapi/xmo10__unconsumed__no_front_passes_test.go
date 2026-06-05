//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_NoFront_Passes — no-front 태그 op은 미소비여도 XMO-10 면제

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_NoFront_Passes(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches:  []stml.FetchBlock{{OperationID: "ListUsers"}},
	}}
	noFrontOp := &openapi3.PathItem{Get: &openapi3.Operation{OperationID: "GetUser", Tags: []string{"no-front"}}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users":      getOp("ListUsers", nil, nil),
		"/users/{id}": noFrontOp,
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("no-front op should not trigger XMO-10, got %v", diags)
	}
}
