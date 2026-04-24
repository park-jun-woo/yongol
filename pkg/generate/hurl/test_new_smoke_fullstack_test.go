//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newSmokeFullstack — OpenAPI doc + state diagram + 자동 SSaC auth funcs 를 Fullstack 으로 조립

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newSmokeFullstack assembles a minimal *yongol.Fullstack wired with
// the supplied OpenAPI doc and (optionally) state diagrams. Manifest.
// Auth is populated so buildAuthSteps fires.
//
// To keep Phase003 shape detection honest without every test spelling
// out the SSaC stanza, any operationId whose path starts with
// "/auth/register" or "/auth/signup" gets a synthetic signup-shape
// ServiceFunc, and "/auth/login" / "/auth/signin" get a login-shape
// one. Tests can call newSmokeFullstackWithFuncs when explicit SSaC
// fixtures are needed.
func newSmokeFullstack(doc *openapi3.T, diagrams ...*statemachine.StateDiagram) *yongol.Fullstack {
	return newSmokeFullstackWithFuncs(doc, inferAuthServiceFuncs(doc), diagrams...)
}

// newSmokeFullstackWithFuncs is the explicit variant — caller provides
// the SSaC ServiceFuncs slice. Used by detect_auth_ops tests.
func newSmokeFullstackWithFuncs(doc *openapi3.T, funcs []ssac.ServiceFunc, diagrams ...*statemachine.StateDiagram) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest:      &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Auth: &pmanifest.Auth{Type: "jwt"}}},
		OpenAPIDoc:    doc,
		ServiceFuncs:  funcs,
		StateDiagrams: diagrams,
	}
}

// inferAuthServiceFuncs walks the doc's paths and returns a SSaC
// ServiceFunc fixture for every auth-ish path (register/signup/join/
// login/signin). Each returned ServiceFunc has the minimal shape needed
// for detectAuthOps to classify the matching operationId — either a
// single @verify-password sequence (login) or a @call auth.HashPassword
// + @post User.Create pair (signup).
func inferAuthServiceFuncs(doc *openapi3.T) []ssac.ServiceFunc {
	if doc == nil {
		return nil
	}
	var out []ssac.ServiceFunc
	for path, pathItem := range doc.Paths.Map() {
		if pathItem == nil || pathItem.Post == nil || pathItem.Post.OperationID == "" {
			continue
		}
		opID := pathItem.Post.OperationID
		switch {
		case pathLooksLikeSignup(path):
			out = append(out, signupServiceFunc(opID))
		case pathLooksLikeLogin(path):
			out = append(out, loginServiceFunc(opID))
		}
	}
	return out
}

func pathLooksLikeSignup(p string) bool {
	return p == "/auth/register" || p == "/auth/signup" || p == "/auth/join"
}

func pathLooksLikeLogin(p string) bool {
	return p == "/auth/login" || p == "/auth/signin"
}

// signupServiceFunc returns a minimal signup-shape SSaC ServiceFunc
// fixture. The HashPassword @call + User.Create @post pair is what
// detectAuthOps matches on.
func signupServiceFunc(opID string) ssac.ServiceFunc {
	return ssac.ServiceFunc{
		Name: opID,
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqCall,
				Model: "auth.HashPassword",
				Inputs: map[string]string{
					"Password": "request.password",
				},
			},
			{
				Type:  ssac.SeqPost,
				Model: "User.Create",
				Inputs: map[string]string{
					"Email":        "request.email",
					"PasswordHash": "hp.HashedPassword",
				},
			},
		},
	}
}

// loginServiceFunc returns a minimal login-shape SSaC ServiceFunc
// fixture — one @verify-password sequence is sufficient.
func loginServiceFunc(opID string) ssac.ServiceFunc {
	return ssac.ServiceFunc{
		Name: opID,
		Sequences: []ssac.Sequence{
			{
				Type:         ssac.SeqVerifyPassword,
				Model:        "User",
				EmailCol:     "email",
				EmailExpr:    "request.email",
				HashCol:      "password_hash",
				PasswordExpr: "request.password",
			},
		},
	}
}
