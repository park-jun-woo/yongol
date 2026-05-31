//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestNormalizedCallKey(t *testing.T) {
	if got := normalizedCallKey("billing.DeductCredit"); got != "billing.deductCredit" {
		t.Errorf("got %q, want billing.deductCredit", got)
	}
	// no dot → unchanged.
	if got := normalizedCallKey("plain"); got != "plain" {
		t.Errorf("got %q, want plain", got)
	}
}
