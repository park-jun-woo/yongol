//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what hasPasswordField — op 의 request body JSON schema 에 password 프로퍼티 존재 여부

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// hasPasswordField returns true when the operation's JSON request body
// schema declares a top-level `password` property. Matching is
// case-sensitive (OpenAPI field names follow the DSL casing).
//
// Auth-shape detection uses this as a cheap first-pass filter before
// inspecting the SSaC body — ops without a password field cannot be
// signup/login under yongol's convention.
func hasPasswordField(op *openapi3.Operation) bool {
	if op == nil || op.RequestBody == nil || op.RequestBody.Value == nil {
		return false
	}
	content := op.RequestBody.Value.Content
	if content == nil {
		return false
	}
	for _, mt := range content {
		if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		if _, ok := mt.Schema.Value.Properties["password"]; ok {
			return true
		}
	}
	return false
}
