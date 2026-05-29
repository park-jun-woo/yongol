//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what extractRefName 단위 테스트 (직접 $ref / array items $ref)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRefName(t *testing.T) {
	cases := []struct {
		name string
		ref  *openapi3.SchemaRef
		want string
	}{
		{"nil", nil, ""},
		{
			name: "direct ref",
			ref:  &openapi3.SchemaRef{Ref: "#/components/schemas/Workflow"},
			want: "Workflow",
		},
		{
			name: "array items ref",
			ref: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:  &openapi3.Types{"array"},
					Items: &openapi3.SchemaRef{Ref: "#/components/schemas/Action"},
				},
			},
			want: "Action",
		},
		{
			name: "scalar no ref",
			ref: &openapi3.SchemaRef{
				Value: &openapi3.Schema{Type: &openapi3.Types{"string"}},
			},
			want: "",
		},
		{
			name: "array without items ref",
			ref: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:  &openapi3.Types{"array"},
					Items: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				},
			},
			want: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := extractRefName(tc.ref); got != tc.want {
				t.Errorf("extractRefName(%s) = %q, want %q", tc.name, got, tc.want)
			}
		})
	}
}
