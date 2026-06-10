//ff:func feature=stml-gen type=test control=sequence
//ff:what renderCaptureCommit — token만/token+refresh/미지정 sink의 setAuth 호출 렌더 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderCaptureCommit(t *testing.T) {
	// token only -> single-arg setAuth
	tokenOnly := []stmlparser.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
	}
	if got, want := renderCaptureCommit(tokenOnly), "useAuthStore.getState().setAuth(data.access_token)"; got != want {
		t.Errorf("token only = %q, want %q", got, want)
	}

	// token + refresh -> two-arg setAuth
	both := []stmlparser.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
		{RespField: "refresh_token", Sink: "auth.refresh"},
	}
	if got, want := renderCaptureCommit(both), "useAuthStore.getState().setAuth(data.access_token, data.refresh_token)"; got != want {
		t.Errorf("both = %q, want %q", got, want)
	}

	// unknown sink only -> token stays undefined, no refresh
	unknown := []stmlparser.CaptureBind{
		{RespField: "whatever", Sink: "auth.other"},
	}
	if got, want := renderCaptureCommit(unknown), "useAuthStore.getState().setAuth(undefined)"; got != want {
		t.Errorf("unknown sink = %q, want %q", got, want)
	}
}
