//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what jsonResp — 스키마 ref 를 지정 status code 의 application/json 응답으로 래핑

package openapi_ddl

import "github.com/getkin/kin-openapi/openapi3"

// jsonResp wraps a schema ref as a 2xx application/json response with the given
// status code.
func jsonResp(code string, ref *openapi3.SchemaRef) *openapi3.Responses {
	r := openapi3.NewResponses()
	r.Set(code, &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.Content{
			"application/json": &openapi3.MediaType{Schema: ref},
		},
	}})
	return r
}
