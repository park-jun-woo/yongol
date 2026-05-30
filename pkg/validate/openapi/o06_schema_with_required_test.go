//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=openapi-structural
//ff:what o06SchemaWithRequired — props 집합 + required 리스트로 SchemaRef 빌드

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// o06SchemaWithRequired builds a SchemaRef with the given property names (each
// an empty schema) and required list. Used by O-6 unit tests.
func o06SchemaWithRequired(props []string, required []string) *openapi3.SchemaRef {
	p := openapi3.Schemas{}
	for _, name := range props {
		p[name] = &openapi3.SchemaRef{Value: openapi3.NewSchema()}
	}
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Properties: p,
		Required:   required,
	}}
}
