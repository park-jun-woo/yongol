//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what hasPgtypexColumn 단위 테스트 (pgtypex bridge 사용 컬럼 존재 여부)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestHasPgtypexColumn(t *testing.T) {
	tables := []ddl.Table{
		{
			Name: "users",
			Columns: map[string]ddl.Column{
				"id":   {Name: "id", RawType: "UUID", NotNull: true},   // pgtypex.FromPgUUID
				"name": {Name: "name", RawType: "TEXT", NotNull: true}, // native
			},
		},
	}
	uuidSchema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"id": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	nativeSchema := &openapi3.Schema{
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
		{"uuid pgtypex column", uuidSchema, true},
		{"native column only", nativeSchema, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := hasPgtypexColumn(tc.schema, tables, "User"); got != tc.want {
				t.Errorf("hasPgtypexColumn(%s) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
