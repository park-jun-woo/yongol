//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what hasJSONBProperty 단위 테스트 (JSONB shape 프로퍼티 존재 여부)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestHasJSONBProperty(t *testing.T) {
	boolTrue := true
	jsonbSchema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"meta": &openapi3.SchemaRef{Value: &openapi3.Schema{
				Type:                 &openapi3.Types{"object"},
				AdditionalProperties: openapi3.AdditionalProperties{Has: &boolTrue},
			}},
			"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
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
		{"has jsonb prop", jsonbSchema, true},
		{"no jsonb prop", plainSchema, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := hasJSONBProperty(tc.schema); got != tc.want {
				t.Errorf("hasJSONBProperty(%s) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
