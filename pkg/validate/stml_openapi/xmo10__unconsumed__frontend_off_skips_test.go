//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_FrontendOff_Skips — Frontend OFF면 XMO-10 전부 skip

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMO10_FrontendOff_Skips(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches:  []stml.FetchBlock{{OperationID: "ListUsers"}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users":      getOp("ListUsers", nil, nil),
		"/users/{id}": getOp("GetUser", nil, nil),
	})
	fs := &yongol.Fullstack{
		SpecsDir:   "/tmp/test-project",
		STMLPages:  pages,
		OpenAPIDoc: doc,
		Manifest:   &manifest.ProjectConfig{}, // frontend OFF
	}
	diags := Run(fs)
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("frontend OFF should skip XMO-10, got %v", diags)
	}
}
