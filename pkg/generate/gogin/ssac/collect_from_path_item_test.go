//ff:func feature=gen-gogin type=test control=sequence
//ff:what collectFromPathItem 단위 테스트 (PathItem의 모든 verb 200 $ref 수집)
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectFromPathItem(t *testing.T) {
	item := &openapi3.PathItem{
		Get:  jsonResponseOp("GET", 200, &openapi3.SchemaRef{Ref: "#/components/schemas/Workflow"}),
		Post: jsonResponseOp("POST", 201, &openapi3.SchemaRef{Ref: "#/components/schemas/Action"}),
	}
	out := map[string]bool{}
	collectFromPathItem(item, out)
	if !out["Workflow"] {
		t.Errorf("GET ref missing: %v", out)
	}
	if !out["Action"] {
		t.Errorf("POST ref missing: %v", out)
	}
	if len(out) != 2 {
		t.Errorf("expected 2 refs, got %v", out)
	}
}
