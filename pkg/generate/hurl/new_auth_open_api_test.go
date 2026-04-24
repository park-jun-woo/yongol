//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newAuthOpenAPI — 임의 operationId 2개로 auth-only OpenAPI 생성

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// newAuthOpenAPI builds an OpenAPI doc with exactly two auth operations
// using the given operationIds. Used by table-driven tests to cover
// alternative names (Signup, Join, SignIn, …) without duplicating the
// schema scaffold.
func newAuthOpenAPI(signupOpID, loginOpID string) *openapi3.T {
	emptySec := openapi3.SecurityRequirements{}
	regBody := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"email":    {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "email"}},
			"org_name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"password": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "password"}},
		},
	}
	loginBody := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"email":    {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "email"}},
			"password": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "password"}},
		},
	}
	register := &openapi3.Operation{
		OperationID: signupOpID,
		Security:    &emptySec,
		RequestBody: &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
			WithContent(openapi3.NewContentWithJSONSchema(regBody))},
		Responses: newCreatedResponses(),
	}
	login := &openapi3.Operation{
		OperationID: loginOpID,
		Security:    &emptySec,
		RequestBody: &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
			WithContent(openapi3.NewContentWithJSONSchema(loginBody))},
		Responses: newOKResponses(),
	}
	return &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/register", &openapi3.PathItem{Post: register}),
		openapi3.WithPath("/auth/login", &openapi3.PathItem{Post: login}),
	)}
}
