//ff:func feature=external type=generator control=iteration dimension=1
//ff:what TestSchemaToGoType: OpenAPI schema type+format → Go 타입 매핑 테이블 드리븐 검증
package external

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSchemaToGoType(t *testing.T) {
	tests := []struct {
		typ    string
		format string
		want   string
	}{
		{"integer", "int64", "int64"},
		{"integer", "int32", "int32"},
		{"integer", "", "int"},
		{"number", "float", "float32"},
		{"number", "double", "float64"},
		{"number", "", "float64"},
		{"string", "", "string"},
		{"string", "date-time", "time.Time"},
		{"boolean", "", "bool"},
	}

	for _, tt := range tests {
		s := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:   &openapi3.Types{tt.typ},
				Format: tt.format,
			},
		}
		got := schemaToGoType(s)
		if got != tt.want {
			t.Errorf("schemaToGoType(%s/%s) = %q, want %q", tt.typ, tt.format, got, tt.want)
		}
	}
}
