//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what logoutDoc — auth 필수 POST /auth/logout op 가진 OpenAPI doc 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// logoutDoc builds an OpenAPI doc with a POST /auth/logout operation that
// requires auth (op-level security set). The operationId is the given opID.
func logoutDoc(opID string) *openapi3.T {
	sec := openapi3.SecurityRequirements{
		openapi3.SecurityRequirement{"bearerAuth": []string{}},
	}
	op := &openapi3.Operation{
		OperationID: opID,
		Security:    &sec,
		Responses:   openapi3.NewResponses(),
	}
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(openapi3.WithPath("/auth/logout", &openapi3.PathItem{Post: op})),
	}
	return doc
}
