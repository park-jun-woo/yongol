//ff:func feature=stml-gen type=test control=sequence
//ff:what renderOnSuccessHandler — onError 리셋·캡처 커밋·리다이렉트·invalidate 분기 렌더 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderOnSuccessHandler(t *testing.T) {
	// error reset + capture commit + redirect: (data) param, all three lines
	a := stmlparser.ActionBlock{OperationID: "Login", OnErrorNode: true, Redirect: "/home"}
	captures := []stmlparser.CaptureBind{{RespField: "access_token", Sink: "auth.token"}}
	want := "    onSuccess: (data) => {\n" +
		"      setLoginError(null)\n" +
		"      useAuthStore.getState().setAuth(data.access_token)\n" +
		"      navigate('/home')\n" +
		"    },\n"
	if got := renderOnSuccessHandler(a, captures, nil); got != want {
		t.Errorf("capture+redirect = %q, want %q", got, want)
	}

	// no captures, no redirect, no on-error: () param, invalidate fallback
	plain := stmlparser.ActionBlock{OperationID: "CreateThing"}
	wantPlain := "    onSuccess: () => {\n" +
		"      queryClient.invalidateQueries({ queryKey: ['ListThings'] })\n" +
		"    },\n"
	if got := renderOnSuccessHandler(plain, nil, []string{"ListThings"}); got != wantPlain {
		t.Errorf("invalidate fallback = %q, want %q", got, wantPlain)
	}
}
