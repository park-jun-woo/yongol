//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO11_FrontendOn_NoStml_Error — Frontend ON + STML 0개 → ERROR 1건

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMO11_FrontendOn_NoStml_Error(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{"/u": getOp("ListUsers", nil, nil)})
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest:   &manifest.ProjectConfig{Frontend: manifest.Frontend{Lang: "typescript"}},
	}
	diags := Run(fs)
	if countDiag(diags, "[XMO-11]") != 1 {
		t.Fatalf("expected 1 XMO-11, got %d (%v)", countDiag(diags, "[XMO-11]"), diags)
	}
}
