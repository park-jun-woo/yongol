//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"testing"
)

func TestToLower(t *testing.T) {
	if got := toLower("DeleteWorkflow_123"); got != "deleteworkflow_123" {
		t.Errorf("toLower = %q", got)
	}
	if got := toLower(""); got != "" {
		t.Errorf("toLower empty = %q", got)
	}
}
