//ff:func feature=gen-gogin type=test control=sequence topic=dos-guard
//ff:what buildOperationRouteIndex — OpenAPI 의 operationId → "METHOD /path" 매핑
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
