//ff:func feature=ground type=test control=iteration dimension=1
//ff:what resolvePrimitiveType — OpenAPI primitive → Go type 매핑 전수 검증

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestResolvePrimitiveType_Cases covers integer/number/string/boolean/array/
// object + int64/float format + nil / empty branches.
func TestResolvePrimitiveType_Cases(t *testing.T) {
	cases := []struct {
		name string
		s    *openapi3.Schema
		want string
	}{
		{"nil", nil, ""},
		{"empty-type", &openapi3.Schema{Type: &openapi3.Types{}}, ""},
		{"integer-default", &openapi3.Schema{Type: &openapi3.Types{"integer"}}, "int"},
		{"integer-int64", &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int64"}, "int64"},
		{"number-default", &openapi3.Schema{Type: &openapi3.Types{"number"}}, "float64"},
		{"number-float", &openapi3.Schema{Type: &openapi3.Types{"number"}, Format: "float"}, "float32"},
		{"string", &openapi3.Schema{Type: &openapi3.Types{"string"}}, "string"},
		{"boolean", &openapi3.Schema{Type: &openapi3.Types{"boolean"}}, "bool"},
		{"object", &openapi3.Schema{Type: &openapi3.Types{"object"}}, "object"},
		{"array-nil-items", &openapi3.Schema{Type: &openapi3.Types{"array"}}, ""},
		{
			"array-string",
			&openapi3.Schema{
				Type: &openapi3.Types{"array"},
				Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Type: &openapi3.Types{"string"},
				}},
			},
			"[]string",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := resolvePrimitiveType(c.s); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}
