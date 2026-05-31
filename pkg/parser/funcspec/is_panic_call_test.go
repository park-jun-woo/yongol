//ff:func feature=funcspec type=test control=sequence
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"testing"
)

func TestIsPanicCall(t *testing.T) {
	body := parseBody(t, `panic("TODO")`)
	if !isPanicCall(body.List[0]) {
		t.Error("expected panic call detected")
	}
	body2 := parseBody(t, `return Resp{}, nil`)
	if isPanicCall(body2.List[0]) {
		t.Error("return is not a panic call")
	}
}
