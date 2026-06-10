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

	// data-on-error -> onError handler updating the error state setter
	a := stmlparser.ActionBlock{OperationID: "Login", OnErrorNode: true}
	want := "    onError: (err) => {\n" +
		"      setLoginError(err instanceof Error ? err.message : String(err))\n" +
		"    },\n"
	if got := renderOnErrorHandler(a); got != want {
		t.Errorf("on-error handler = %q, want %q", got, want)
	}
}
