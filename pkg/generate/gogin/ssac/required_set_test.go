//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what requiredSet 단위 테스트 (schema.Required → lookup map, nil 안전)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRequiredSet(t *testing.T) {
	t.Run("flattens required list", func(t *testing.T) {
		schema := &openapi3.Schema{Required: []string{"name", "email"}}
		got := requiredSet(schema)
		if !got["name"] || !got["email"] {
			t.Errorf("expected name/email present, got %v", got)
		}
		if got["missing"] {
			t.Errorf("absent key should be false")
		}
		if len(got) != 2 {
			t.Errorf("expected 2 entries, got %d", len(got))
		}
	})

	t.Run("empty required", func(t *testing.T) {
		got := requiredSet(&openapi3.Schema{})
		if got == nil || len(got) != 0 {
			t.Errorf("expected empty non-nil map, got %v", got)
		}
	})
}
