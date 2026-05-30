//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10Unconsumed_EarlyReturn — nil doc / nil paths / 빈 페이지 early return

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10Unconsumed_EarlyReturn(t *testing.T) {
	pages := []stml.PageSpec{{FileName: "p.html"}}

	// nil doc.
	if got := xmo10Unconsumed(pages, nil); got != nil {
		t.Errorf("nil doc: expected nil, got %v", got)
	}
	// nil paths.
	if got := xmo10Unconsumed(pages, &openapi3.T{}); got != nil {
		t.Errorf("nil paths: expected nil, got %v", got)
	}
	// empty pages.
	doc := makeDoc(map[string]*openapi3.PathItem{"/u": getOp("ListUsers", nil, nil)})
	if got := xmo10Unconsumed(nil, doc); got != nil {
		t.Errorf("no pages: expected nil, got %v", got)
	}
}
