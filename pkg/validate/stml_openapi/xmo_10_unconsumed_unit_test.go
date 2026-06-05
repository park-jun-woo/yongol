//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10Unconsumed_EarlyReturn — frontend OFF / nil doc / 빈 페이지 early return

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMO10Unconsumed_EarlyReturn(t *testing.T) {
	pages := []stml.PageSpec{{FileName: "p.html"}}
	doc := makeDoc(map[string]*openapi3.PathItem{"/u": getOp("ListUsers", nil, nil)})

	// Frontend OFF (nil manifest) → skip.
	if got := xmo10Unconsumed(&yongol.Fullstack{STMLPages: pages, OpenAPIDoc: doc}); got != nil {
		t.Errorf("frontend off: expected nil, got %v", got)
	}
	// Frontend ON but 0 STML pages → yielded to XMO-11, no XMO-10.
	onFS := func(p []stml.PageSpec, d *openapi3.T) *yongol.Fullstack {
		return &yongol.Fullstack{
			STMLPages:  p,
			OpenAPIDoc: d,
			Manifest:   &manifest.ProjectConfig{Frontend: manifest.Frontend{Lang: "typescript"}},
		}
	}
	if got := xmo10Unconsumed(onFS(nil, doc)); got != nil {
		t.Errorf("no pages: expected nil, got %v", got)
	}
	// Frontend ON, nil doc → skip.
	if got := xmo10Unconsumed(onFS(pages, nil)); got != nil {
		t.Errorf("nil doc: expected nil, got %v", got)
	}
}
