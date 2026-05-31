//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"testing"
)

func TestHasDeleteAction(t *testing.T) {
	if !hasDeleteAction([]string{"read", "delete"}) {
		t.Error("expected delete present")
	}
	if hasDeleteAction([]string{"read", "write"}) {
		t.Error("expected no delete")
	}
	if hasDeleteAction(nil) {
		t.Error("nil should not contain delete")
	}
}
