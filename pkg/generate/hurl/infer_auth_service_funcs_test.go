//ff:func feature=gen-hurl type=test-helper control=iteration dimension=1
//ff:what inferAuthServiceFuncs — OpenAPI paths 에서 auth-shape SSaC fixtures 추론

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
