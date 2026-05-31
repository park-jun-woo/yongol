//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestIsAuthEndpoint(t *testing.T) {
	authOp := &openapi3.Operation{Security: &openapi3.SecurityRequirements{}}
	if !isAuthEndpoint(authOp) {
		t.Error("empty security requirement should be an auth endpoint")
	}
	// nil security → not auth.
	if isAuthEndpoint(&openapi3.Operation{}) {
		t.Error("nil security should not be auth endpoint")
	}
	// non-empty security → not auth.
	reqs := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearer": {}}}
	if isAuthEndpoint(&openapi3.Operation{Security: &reqs}) {
		t.Error("non-empty security should not be auth endpoint")
	}
}
