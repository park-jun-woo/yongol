//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what apiCastFor 단위 테스트 (enum / email / uuid 캐스트 선택)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestApiCastFor(t *testing.T) {
	cases := []struct {
		name   string
		parent string
		json   string
		ref    *openapi3.SchemaRef
		want   string
	}{
		{"nil ref", "Workflow", "status", nil, ""},
		{
			name:   "enum string",
			parent: "Workflow",
			json:   "status",
			ref:    &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Enum: []interface{}{"a", "b"}}},
			want:   "api.WorkflowStatus",
		},
		{
			name: "email format",
			json: "email",
			ref:  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "email"}},
			want: "openapi_types.Email",
		},
		{
			name: "uuid format",
			json: "ownerId",
			ref:  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "uuid"}},
			want: "openapi_types.UUID",
		},
		{
			name: "plain string no cast",
			json: "name",
			ref:  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			want: "",
		},
		{
			name: "non-string no cast",
			json: "count",
			ref:  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			want: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := apiCastFor(tc.parent, tc.json, tc.ref); got != tc.want {
				t.Errorf("apiCastFor(%q,%q) = %q, want %q", tc.parent, tc.json, got, tc.want)
			}
		})
	}
}
