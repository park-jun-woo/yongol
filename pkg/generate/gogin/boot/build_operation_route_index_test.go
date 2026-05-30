//ff:func feature=gen-gogin type=test control=iteration dimension=2 topic=dos-guard
//ff:what buildOperationRouteIndex — OpenAPI 의 operationId → "METHOD /path" 매핑

package boot

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildOperationRouteIndex_Empty(t *testing.T) {
	for _, fs := range []*yongol.Fullstack{
		nil,
		{},
		{Manifest: &pmanifest.ProjectConfig{}},
	} {
		if idx := buildOperationRouteIndex(fs); len(idx) != 0 {
			t.Errorf("expected empty index, got %v", idx)
		}
	}
}

func TestBuildOperationRouteIndex_Maps(t *testing.T) {
	doc := buildDoc([]opSpec{
		{path: "/users/{id}", method: "GET", opID: "GetUser"},
		{path: "/users", method: "POST", opID: "CreateUser"},
	}, false)
	fs := &yongol.Fullstack{OpenAPIDoc: doc}
	idx := buildOperationRouteIndex(fs)
	if idx["GetUser"] != "GET /users/:id" {
		t.Errorf("GetUser = %q, want GET /users/:id", idx["GetUser"])
	}
	if idx["CreateUser"] != "POST /users" {
		t.Errorf("CreateUser = %q, want POST /users", idx["CreateUser"])
	}
}

func TestBuildOperationRouteIndex_SkipsEmptyOperationID(t *testing.T) {
	// An operation without an OperationID must be skipped (continue branch).
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	pi := &openapi3.PathItem{Get: &openapi3.Operation{OperationID: ""}}
	doc.Paths.Set("/anon", pi)
	fs := &yongol.Fullstack{OpenAPIDoc: doc}
	idx := buildOperationRouteIndex(fs)
	if len(idx) != 0 {
		t.Errorf("operation without OperationID must be skipped, got %v", idx)
	}
}
