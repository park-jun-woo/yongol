//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO11_FrontendOff_NoStml_None — Frontend OFF + STML 0개 → 진단 없음

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMO11_FrontendOff_NoStml_None(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{"/u": getOp("ListUsers", nil, nil)})
	fs := &yongol.Fullstack{OpenAPIDoc: doc, Manifest: &manifest.ProjectConfig{}}
	if got := xmo11NoStml(fs); got != nil {
		t.Errorf("frontend off: expected no XMO-11, got %v", got)
	}
}
