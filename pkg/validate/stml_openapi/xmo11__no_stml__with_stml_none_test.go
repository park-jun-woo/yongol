//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO11_FrontendOn_WithStml_None — Frontend ON + STML 있음 → XMO-11 없음

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO11_FrontendOn_WithStml_None(t *testing.T) {
	pages := []stml.PageSpec{{FileName: "u.html", Fetches: []stml.FetchBlock{{OperationID: "ListUsers"}}}}
	doc := makeDoc(map[string]*openapi3.PathItem{"/u": getOp("ListUsers", nil, nil)})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-11]") {
		t.Errorf("STML present: expected no XMO-11, got %v", diags)
	}
}
