//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what hasJSONContentWithSchema — nil/empty/Ref/Value 경로별 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestHasJSONContentWithSchema(t *testing.T) {
	tests := []struct {
		name string
		ref  *openapi3.ResponseRef
		want bool
	}{
		{
			name: "nil ref returns false",
			ref:  nil,
			want: false,
		},
		{
			name: "nil Value returns false",
			ref:  &openapi3.ResponseRef{Value: nil},
			want: false,
		},
		{
			name: "nil Content returns false",
			ref:  &openapi3.ResponseRef{Value: &openapi3.Response{}},
			want: false,
		},
		{
			name: "no application/json media type returns false",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"text/plain": &openapi3.MediaType{},
					},
				},
			},
			want: false,
		},
		{
			name: "application/json but nil schema returns false",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{Schema: nil},
					},
				},
			},
			want: false,
		},
		{
			name: "application/json with empty SchemaRef (no Value, no Ref) returns false",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "application/json with $ref returns true",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: "#/components/schemas/Error",
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "application/json with inline Value returns true",
			ref: &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Value: &openapi3.Schema{Type: &openapi3.Types{"object"}},
							},
						},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasJSONContentWithSchema(tt.ref)
			if got != tt.want {
				t.Errorf("hasJSONContentWithSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}
