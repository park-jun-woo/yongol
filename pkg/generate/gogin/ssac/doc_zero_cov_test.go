//ff:func feature=gen-gogin type=test control=sequence
//ff:what zz_zerocov_extract — 0% OpenAPI 추출 헬퍼(applyOperation/extractFromOpenAPI/tryExtractFromPathItem/extractBodyFormats/extractRespFields) 검증
package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func docZeroCov(operationID string) *openapi3.T {
	strType := &openapi3.Types{"string"}
	reqSchema := openapi3.NewSchema()
	reqSchema.Type = &openapi3.Types{"object"}
	reqSchema.Required = []string{"email"}
	reqSchema.Properties = openapi3.Schemas{
		"email": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType, Format: "email"}},
		"plan":  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType, Enum: []any{"free", "pro"}}},
		"meta":  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, AdditionalProperties: openapi3.AdditionalProperties{Has: boolPtr(true)}}},
	}

	respSchema := openapi3.NewSchema()
	respSchema.Type = &openapi3.Types{"object"}
	respSchema.Required = []string{"widget"}
	respSchema.Properties = openapi3.Schemas{
		"widget": &openapi3.SchemaRef{Ref: "#/components/schemas/Widget"},
		"name":   &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType}},
	}

	op := openapi3.NewOperation()
	op.OperationID = operationID
	op.Parameters = openapi3.Parameters{
		{Value: &openapi3.Parameter{Name: "id", In: "path", Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType}}}},
		{Value: &openapi3.Parameter{Name: "q", In: "query", Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: strType}}}},
	}
	op.RequestBody = &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchema(reqSchema)}
	resp := openapi3.NewResponse().WithJSONSchema(respSchema)
	op.Responses = openapi3.NewResponses()
	op.Responses.Set("200", &openapi3.ResponseRef{Value: resp})

	pathItem := &openapi3.PathItem{Get: op}
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/widgets/{id}", pathItem)
	return doc
}
