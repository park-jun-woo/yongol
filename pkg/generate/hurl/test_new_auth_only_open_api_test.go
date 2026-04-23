//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newAuthOnlyOpenAPI — Register+Login 두 operation 만 포함한 테스트용 *openapi3.T 생성

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// newAuthOnlyOpenAPI returns an *openapi3.T with just Register + Login
// auth operations wired under /auth/register and /auth/login. Callers
// typically attach this to Fullstack.OpenAPIDoc and run the smoke order
// builder to assert auth-step ordering without needing the rest of the
// surface.
func newAuthOnlyOpenAPI() *openapi3.T {
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
		OperationID: "Register",
		RequestBody: &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
			WithContent(openapi3.NewContentWithJSONSchema(regBody))},
		Responses: newCreatedResponses(),
	}
	login := &openapi3.Operation{
		OperationID: "Login",
		RequestBody: &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
			WithContent(openapi3.NewContentWithJSONSchema(loginBody))},
		Responses: newOKResponses(),
	}
	return &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/register", &openapi3.PathItem{Post: register}),
		openapi3.WithPath("/auth/login", &openapi3.PathItem{Post: login}),
	)}
}
