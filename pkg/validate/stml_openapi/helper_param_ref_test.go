//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what paramRef — 테스트용 parameter ref 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// paramRef creates a parameter ref.
func paramRef(name, in string) *openapi3.ParameterRef {
	return &openapi3.ParameterRef{Value: &openapi3.Parameter{Name: name, In: in}}
}
