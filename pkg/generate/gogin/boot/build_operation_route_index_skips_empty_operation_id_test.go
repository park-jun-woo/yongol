//ff:func feature=gen-gogin type=test control=sequence topic=dos-guard
//ff:what buildOperationRouteIndex — OpenAPI 의 operationId → "METHOD /path" 매핑
package boot

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
