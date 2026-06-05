//ff:func feature=gen-ir type=test control=sequence
//ff:what TestParamSchemaType -- 파라미터 첫 schema type 추출, schema/type 부재 시 "" 검증

package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestParamSchemaType(t *testing.T) {
	withType := &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}},
			},
		},
	}
	if got := paramSchemaType(withType); got != "integer" {
		t.Errorf("paramSchemaType(integer) = %q, want integer", got)
	}

	// schema absent
	noSchema := &openapi3.ParameterRef{Value: &openapi3.Parameter{}}
	if got := paramSchemaType(noSchema); got != "" {
		t.Errorf("paramSchemaType(no schema) = %q, want empty", got)
	}

	// schema present but type nil
	noType := &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{}},
		},
	}
	if got := paramSchemaType(noType); got != "" {
		t.Errorf("paramSchemaType(no type) = %q, want empty", got)
	}
}
