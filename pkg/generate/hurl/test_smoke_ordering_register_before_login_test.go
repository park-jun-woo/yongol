//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestSmokeOrderingSignupBeforeLogin — 명명과 무관하게 signup-shape op 이 login-shape 보다 먼저 배치되는지 검증 (BUG-015 + BUG-023)

package hurl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestSmokeOrderingSignupBeforeLogin pins BUG-015 Phase003 + BUG-023:
// whenever OpenAPI exposes both a signup-shape and a login-shape op,
// signup must appear first in smoke so an empty DB hits POST /auth/...
// (201) before the login runs against a freshly-created user. The
// previous implementation compared operationId literally against
// "Register" — any other name (Signup, Join, ...) silently broke the
// pair. Shape detection replaces that.
func TestSmokeOrderingSignupBeforeLogin(t *testing.T) {
	cases := []struct {
		name             string
		signupOpID       string
		loginOpID        string
	}{
		{"Register+Login", "Register", "Login"},
		{"Signup+Login", "Signup", "Login"},
		{"Join+SignIn", "Join", "SignIn"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			doc := synthAuthDoc(tc.signupOpID, tc.loginOpID)
			fs := newSmokeFullstack(doc)
			steps := buildScenarioOrder(fs)
			opIDs := stepOpIDs(steps)
			sIdx := indexOfString(opIDs, tc.signupOpID)
			lIdx := indexOfString(opIDs, tc.loginOpID)
			if sIdx < 0 {
				t.Fatalf("%s missing from smoke steps: %v", tc.signupOpID, opIDs)
			}
			if lIdx < 0 {
				t.Fatalf("%s missing from smoke steps: %v", tc.loginOpID, opIDs)
			}
			if sIdx > lIdx {
				t.Errorf("%s (idx=%d) must precede %s (idx=%d); got order=%v",
					tc.signupOpID, sIdx, tc.loginOpID, lIdx, opIDs)
			}
		})
	}
}

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
