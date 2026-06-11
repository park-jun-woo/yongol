//ff:func feature=stml-gen type=test control=sequence
//ff:what renderOnErrorHandler — data-on-error 유무·표시 필드 도출·빈 필드 보정에 따른 mutation onError 핸들러 렌더 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderOnErrorHandler(t *testing.T) {
	// The thrown ErrorResponse is a plain object (BUG-113). The schema-derived
	// display field ("error") is read first, message second (real Error
	// instances from network rejects), String(err) only as the final fallback.
	want := "    onError: (err) => {\n" +
		"      const e = err as any\n" +
		"      const msg = e?.error ?? e?.message\n" +
		"      setLoginError(typeof msg === 'string' && msg !== '' ? msg : String(err))\n" +
		"    },\n"

	t.Run("default error field, data-on-error declared", func(t *testing.T) {
		a := stmlparser.ActionBlock{OperationID: "Login", OnErrorNode: true}
		if got := renderOnErrorHandler(a, "error"); got != want {
			t.Errorf("on-error handler = %q, want %q", got, want)
		}
	})

	t.Run("default error field, no data-on-error still emitted", func(t *testing.T) {
		// page-flow Phase004: state feeds the default display slot instead.
		if got := renderOnErrorHandler(stmlparser.ActionBlock{OperationID: "Login"}, "error"); got != want {
			t.Errorf("default on-error handler = %q, want %q", got, want)
		}
		// the stale single `(err as any)?.message` read must not survive
		if strings.Contains(want, "(err as any)?.message") {
			t.Fatalf("golden still contains the BUG-125 single message read")
		}
	})

	t.Run("non-default display field threaded", func(t *testing.T) {
		got := renderOnErrorHandler(stmlparser.ActionBlock{OperationID: "Login"}, "message")
		if !strings.Contains(got, "const msg = e?.message ?? e?.message") {
			t.Errorf("message display field not threaded: %q", got)
		}
	})

	t.Run("empty display field normalized to error", func(t *testing.T) {
		// partially constructed GenerateOptions must not emit a broken expr.
		got := renderOnErrorHandler(stmlparser.ActionBlock{OperationID: "Login"}, "")
		if !strings.Contains(got, "const msg = e?.error ?? e?.message") {
			t.Errorf("empty displayField not normalized to error: %q", got)
		}
		if strings.Contains(got, "e?. ??") {
			t.Errorf("broken empty-field expression emitted: %q", got)
		}
	})
}
