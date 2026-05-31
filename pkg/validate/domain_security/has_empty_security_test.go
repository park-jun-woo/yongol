//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestHasEmptySecurity(t *testing.T) {
	if !hasEmptySecurity(&openapi3.Operation{Security: &openapi3.SecurityRequirements{}}) {
		t.Error("security: [] should be empty")
	}
	if hasEmptySecurity(&openapi3.Operation{}) {
		t.Error("nil security is not empty-security")
	}
	reqs := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"x": {}}}
	if hasEmptySecurity(&openapi3.Operation{Security: &reqs}) {
		t.Error("non-empty security should not report empty")
	}
}
