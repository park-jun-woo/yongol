//ff:func feature=stml-gen type=test control=sequence
//ff:what renderCaptureCommit — 토큰 부재 가드(on-error state/console 분기) + setAuth 커밋 라인 렌더 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderCaptureCommit(t *testing.T) {
	// token only, no data-on-error -> guard surfaces via console.error,
	// then single-arg setAuth
	tokenOnly := []stmlparser.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
	}
	wantTokenOnly := "if (data?.access_token == null) {\n" +
		"  console.error('Unexpected response: missing access_token')\n" +
		"  return\n" +
		"}\n" +
		"useAuthStore.getState().setAuth(data.access_token)"
	if got := strings.Join(renderCaptureCommit(tokenOnly, ""), "\n"); got != wantTokenOnly {
		t.Errorf("token only = %q, want %q", got, wantTokenOnly)
	}

	// token + refresh, with data-on-error state -> guard surfaces via the
	// error state setter; refresh stays optional (token gates the commit)
	both := []stmlparser.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
		{RespField: "refresh_token", Sink: "auth.refresh"},
	}
	wantBoth := "if (data?.access_token == null) {\n" +
		"  setLoginError('Unexpected response: missing access_token')\n" +
		"  return\n" +
		"}\n" +
		"useAuthStore.getState().setAuth(data.access_token, data.refresh_token)"
	if got := strings.Join(renderCaptureCommit(both, "loginError"), "\n"); got != wantBoth {
		t.Errorf("both = %q, want %q", got, wantBoth)
	}

	// unknown sink only -> no token capture, no guard, token stays undefined
	unknown := []stmlparser.CaptureBind{
		{RespField: "whatever", Sink: "auth.other"},
	}
	if got, want := strings.Join(renderCaptureCommit(unknown, ""), "\n"), "useAuthStore.getState().setAuth(undefined)"; got != want {
		t.Errorf("unknown sink = %q, want %q", got, want)
	}
}
