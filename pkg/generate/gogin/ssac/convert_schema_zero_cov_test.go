//ff:func feature=gen-gogin type=test control=sequence
//ff:what zz_zerocov_writers — 0% 생성/기록 헬퍼(generateServerHelpers/writeMethodFile/writeConvertFunc/writeConvertListFunc/emitConvert*File) 검증
package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func convertSchemaZeroCov() *openapi3.Schema {
	s := openapi3.NewSchema()
	s.Type = &openapi3.Types{"object"}
	s.Required = []string{"id"}
	s.Properties = openapi3.Schemas{
		"id":   &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
	}
	return s
}
