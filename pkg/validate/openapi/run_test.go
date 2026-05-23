//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what Run — nil doc 빈 결과 + 문서 존재 시 집계 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunOpenAPI(t *testing.T) {
	t.Run("nil doc returns empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("empty valid doc returns empty", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Operations: map[string]int{}, Paths: map[string]int{}},
		}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
