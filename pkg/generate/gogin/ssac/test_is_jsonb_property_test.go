//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isJSONBProperty 단위 테스트

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestIsJSONBProperty(t *testing.T) {
	boolTrue := true
	boolFalse := false

	cases := []struct {
		name string
		ref  *openapi3.SchemaRef
		want bool
	}{
		{
			name: "additionalProperties true",
			ref: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:                 &openapi3.Types{"object"},
					AdditionalProperties: openapi3.AdditionalProperties{Has: &boolTrue},
				},
			},
			want: true,
		},
		{
			name: "additionalProperties unspecified (nil)",
			ref: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: &openapi3.Types{"object"},
				},
			},
			want: true,
		},
		{
			name: "additionalProperties false",
			ref: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:                 &openapi3.Types{"object"},
					AdditionalProperties: openapi3.AdditionalProperties{Has: &boolFalse},
				},
			},
			want: false,
		},
		{
			name: "object with properties",
			ref: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: &openapi3.Types{"object"},
					Properties: openapi3.Schemas{
						"name": &openapi3.SchemaRef{
							Value: &openapi3.Schema{Type: &openapi3.Types{"string"}},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "type string",
			ref: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: &openapi3.Types{"string"},
				},
			},
			want: false,
		},
		{
			name: "nil SchemaRef",
			ref:  nil,
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isJSONBProperty(tc.ref)
			if got != tc.want {
				t.Errorf("isJSONBProperty(%s) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
