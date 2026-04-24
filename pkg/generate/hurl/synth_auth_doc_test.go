//ff:func feature=gen-hurl type=test-helper control=selection
//ff:what synthAuthDoc — signup+login 픽스처 OpenAPI doc 생성 (경로/opID 자동 매핑)

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// synthAuthDoc returns an OpenAPI doc with signup + login ops under
// standard auth paths so inferAuthServiceFuncs attaches matching SSaC
// shapes. Chooses the /auth/signup path when the signup op is named
// "Signup" (exercises BUG-023 reproduction case directly).
func synthAuthDoc(signupOpID, loginOpID string) *openapi3.T {
	emptySec := openapi3.SecurityRequirements{}
	regBody := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"email":    {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "email"}},
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
	signup := &openapi3.Operation{
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
	signupPath := "/auth/register"
	switch signupOpID {
	case "Signup":
		signupPath = "/auth/signup"
	case "Join":
		signupPath = "/auth/join"
	}
	loginPath := "/auth/login"
	if loginOpID == "SignIn" {
		loginPath = "/auth/signin"
	}
	return &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath(signupPath, &openapi3.PathItem{Post: signup}),
		openapi3.WithPath(loginPath, &openapi3.PathItem{Post: login}),
	)}
}
