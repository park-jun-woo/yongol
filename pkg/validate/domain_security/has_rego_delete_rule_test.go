//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"testing"
)

func TestHasRegoDeleteRule(t *testing.T) {
	res := map[string]struct{}{"workflow": {}}
	if !hasRegoDeleteRule("DeleteWorkflow", res) {
		t.Error("DeleteWorkflow should match workflow resource")
	}
	if hasRegoDeleteRule("DeleteUser", res) {
		t.Error("DeleteUser should not match workflow resource")
	}
	if hasRegoDeleteRule("DeleteWorkflow", map[string]struct{}{}) {
		t.Error("empty rego resources → no match")
	}
}
