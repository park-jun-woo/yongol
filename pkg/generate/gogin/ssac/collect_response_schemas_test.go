//ff:func feature=gen-gogin type=test control=sequence
//ff:what collectResponseSchemas 단위 테스트 (전체 doc의 200 $ref 스키마 이름 집합)
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectResponseSchemas(t *testing.T) {
	t.Run("nil paths → empty map", func(t *testing.T) {
		got := collectResponseSchemas(&openapi3.T{})
		if got == nil || len(got) != 0 {
			t.Errorf("expected empty non-nil map, got %v", got)
		}
	})
	t.Run("collects across paths", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/wf", &openapi3.PathItem{
					Get: jsonResponseOp("GET", 200, &openapi3.SchemaRef{Ref: "#/components/schemas/Workflow"}),
				}),
				openapi3.WithPath("/act", &openapi3.PathItem{
					Post: jsonResponseOp("POST", 201, &openapi3.SchemaRef{Ref: "#/components/schemas/Action"}),
				}),
			),
		}
		got := collectResponseSchemas(doc)
		if !got["Workflow"] || !got["Action"] {
			t.Errorf("expected Workflow+Action, got %v", got)
		}
	})
}
