//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_ComponentConsumed_Passes — 컴포넌트 api.<Op>( 호출 소비도 XMO-10 면제

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMO10_ComponentConsumed_Passes(t *testing.T) {
	specsDir := t.TempDir()
	writeComponent(t, specsDir, "UserCard", `api.GetUser(id);`)
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches:  []stml.FetchBlock{{OperationID: "ListUsers"}},
		Children: []stml.ChildNode{{
			Kind:      "component",
			Component: &stml.ComponentRef{Name: "UserCard"},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users":      getOp("ListUsers", nil, nil),
		"/users/{id}": getOp("GetUser", nil, nil),
	})
	fs := &yongol.Fullstack{
		SpecsDir:   specsDir,
		STMLPages:  pages,
		OpenAPIDoc: doc,
		Manifest:   &manifest.ProjectConfig{Frontend: manifest.Frontend{Lang: "typescript"}},
	}
	diags := Run(fs)
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("component-consumed GetUser should not trigger XMO-10, got %v", diags)
	}
}
