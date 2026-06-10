//ff:func feature=stml-gen type=test control=sequence
//ff:what renderOnErrorHandler — data-on-error 유무와 무관하게 mutation onError 핸들러 상시 렌더 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderOnErrorHandler(t *testing.T) {
	// onError handler updating the error state setter. The thrown
	// ErrorResponse is a plain object and message has no schema guarantee
	// (XOE-01 checks only error/code) — extract a non-empty string message,
	// fall back to String(err).
	want := "    onError: (err) => {\n" +
		"      const msg = (err as any)?.message\n" +
		"      setLoginError(typeof msg === 'string' && msg !== '' ? msg : String(err))\n" +
		"    },\n"

	// data-on-error declared
	a := stmlparser.ActionBlock{OperationID: "Login", OnErrorNode: true}
	if got := renderOnErrorHandler(a); got != want {
		t.Errorf("on-error handler = %q, want %q", got, want)
	}

	// no data-on-error -> still emitted (page-flow Phase004 default onError;
	// the state feeds the default display slot instead)
	if got := renderOnErrorHandler(stmlparser.ActionBlock{OperationID: "Login"}); got != want {
		t.Errorf("default on-error handler = %q, want %q", got, want)
	}
}
