//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestJoinReturnTypes(t *testing.T) {
	if got := joinReturnTypes([]string{"int", "error"}); got != "int, error" {
		t.Errorf("got %q", got)
	}
	if got := joinReturnTypes([]string{"string"}); got != "string" {
		t.Errorf("got %q", got)
	}
	if got := joinReturnTypes(nil); got != "" {
		t.Errorf("got %q", got)
	}
}
