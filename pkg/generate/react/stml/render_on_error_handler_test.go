//ff:func feature=stml-gen type=test control=sequence
//ff:what renderOnErrorHandler — data-on-error 유무에 따른 mutation onError 핸들러 렌더 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderOnErrorHandler(t *testing.T) {
	// no data-on-error -> empty (no handler emitted)
	if got := renderOnErrorHandler(stmlparser.ActionBlock{OperationID: "Login"}); got != "" {
		t.Errorf("no on-error = %q, want empty", got)
	}

	// data-on-error -> onError handler updating the error state setter.
	// The thrown ErrorResponse is a plain object and message has no schema
	// guarantee (XOE-01 checks only error/code) — extract a non-empty
	// string message, fall back to String(err).
	a := stmlparser.ActionBlock{OperationID: "Login", OnErrorNode: true}
	want := "    onError: (err) => {\n" +
		"      const msg = (err as any)?.message\n" +
		"      setLoginError(typeof msg === 'string' && msg !== '' ? msg : String(err))\n" +
		"    },\n"
	if got := renderOnErrorHandler(a); got != want {
		t.Errorf("on-error handler = %q, want %q", got, want)
	}
}
