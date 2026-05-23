//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what schemeNameSet — nil doc/nil components/스킴 수집 검증

package openapi_manifest

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSchemeNameSet(t *testing.T) {
	t.Run("nil doc returns empty", func(t *testing.T) {
		got := schemeNameSet(nil)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("nil components returns empty", func(t *testing.T) {
		got := schemeNameSet(&openapi3.T{})
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("collects scheme names", func(t *testing.T) {
		doc := &openapi3.T{
			Components: &openapi3.Components{
				SecuritySchemes: openapi3.SecuritySchemes{
					"bearerAuth": &openapi3.SecuritySchemeRef{},
					"apiKey":     &openapi3.SecuritySchemeRef{},
				},
			},
		}
		got := schemeNameSet(doc)
		if len(got) != 2 {
			t.Fatalf("expected 2, got %d", len(got))
		}
		if !got["bearerAuth"] || !got["apiKey"] {
			t.Errorf("missing expected schemes: %v", got)
		}
	})
}
