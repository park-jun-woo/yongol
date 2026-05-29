//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what hasOpenAPITypesCast 단위 테스트 (openapi_types.* 캐스트 필드 존재 여부)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestHasOpenAPITypesCast(t *testing.T) {
	emailSchema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"email": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "email"}},
			"name":  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	plainSchema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	cases := []struct {
		name   string
		schema *openapi3.Schema
		want   bool
	}{
		{"nil", nil, false},
		{"has email cast", emailSchema, true},
		{"no cast", plainSchema, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := hasOpenAPITypesCast(tc.schema); got != tc.want {
				t.Errorf("hasOpenAPITypesCast(%s) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
