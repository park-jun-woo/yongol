//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestCallFuncCamelName(t *testing.T) {
	if got := callFuncCamelName("billing.CheckCredits"); got != "checkCredits" {
		t.Errorf("got %q, want checkCredits", got)
	}
	if got := callFuncCamelName("noqualifier"); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}
