//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestRun_Branches — Run early-return / Actions / layouts+manifest 분기 커버
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRun_EarlyReturn(t *testing.T) {
	// nil OpenAPIDoc → nil.
	if got := Run(makeFS([]stml.PageSpec{{Name: "p"}}, nil)); got != nil {
		t.Errorf("nil doc: expected nil, got %v", got)
	}
	// no STML pages → nil.
	doc := makeDoc(map[string]*openapi3.PathItem{"/items": getOp("ListItems", nil, nil)})
	if got := Run(makeFS(nil, doc)); got != nil {
		t.Errorf("no pages: expected nil, got %v", got)
	}
}
