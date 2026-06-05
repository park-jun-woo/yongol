//ff:func feature=openapi-parse type=test control=iteration dimension=1
//ff:what mapValueType — additionalProperties 의 typed/nil/free-form 분기가 값 타입 또는 "*" 마커를 반환하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestMapValueType(t *testing.T) {
	typed := func(typ string) openapi3.AdditionalProperties {
		return openapi3.AdditionalProperties{
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{Type: &openapi3.Types{typ}},
			},
		}
	}
	has := true

	cases := []struct {
		name string
		ap   openapi3.AdditionalProperties
		want string
	}{
		{"string value", typed("string"), "string"},
		{"integer value", typed("integer"), "integer"},
		{
			"schema without declared type",
			openapi3.AdditionalProperties{Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{}}},
			"*",
		},
		{"free-form has=true", openapi3.AdditionalProperties{Has: &has}, "*"},
		{"unspecified", openapi3.AdditionalProperties{}, "*"},
		{"nil schema value", openapi3.AdditionalProperties{Schema: &openapi3.SchemaRef{Value: nil}}, "*"},
	}
	for _, tc := range cases {
		if got := mapValueType(tc.ap); got != tc.want {
			t.Errorf("%s: mapValueType = %q, want %q", tc.name, got, tc.want)
		}
	}
}
