//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestItoa(t *testing.T) {
	if got := itoa(42); got != "42" {
		t.Errorf("itoa(42) = %q", got)
	}
	if got := itoa(0); got != "0" {
		t.Errorf("itoa(0) = %q", got)
	}
}
