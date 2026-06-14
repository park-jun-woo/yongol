//ff:func feature=stml-gen type=test control=sequence
//ff:what renderOnSuccessHandler — 토큰 가드·캡처 커밋·리다이렉트·invalidate 분기 렌더 검증 (에러 리셋은 onMutate 로 이동)
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderOnSuccessHandler(t *testing.T) {
	// token guard + capture commit + redirect: (data) param; the guard's
	// early return aborts both the commit and the navigate. The error reset
	// lives in onMutate since page-flow Phase004, not here.
	a := stmlparser.ActionBlock{OperationID: "Login", OnErrorNode: true, Redirect: "/home"}
	captures := []stmlparser.CaptureBind{{RespField: "access_token", Sink: "auth.token"}}
	want := "    onSuccess: (data) => {\n" +
		"      if (data?.access_token == null) {\n" +
		"        setLoginError('Unexpected response: missing access_token')\n" +
		"        return\n" +
		"      }\n" +
		"      useAuthStore.getState().setAuth(data.access_token)\n" +
		"      navigate('/home')\n" +
		"    },\n"
	if got := renderOnSuccessHandler(a, captures, nil, nil); got != want {
		t.Errorf("capture+redirect = %q, want %q", got, want)
	}

	// no captures, no remove, no redirect: () param, plain invalidate
	plain := stmlparser.ActionBlock{OperationID: "CreateThing"}
	wantPlain := "    onSuccess: () => {\n" +
		"      queryClient.invalidateQueries({ queryKey: ['ListThings'] })\n" +
		"    },\n"
	if got := renderOnSuccessHandler(plain, nil, []string{"ListThings"}, nil); got != wantPlain {
		t.Errorf("plain invalidate = %q, want %q", got, wantPlain)
	}

	// combined invalidate + navigate: a non-capture mutation refreshes the
	// affected list and then navigates to its data-redirect target (BUG-132
	// 132-1) — they are no longer exclusive branches.
	combined := stmlparser.ActionBlock{OperationID: "CreateThing", Redirect: "/things"}
	wantCombined := "    onSuccess: () => {\n" +
		"      queryClient.invalidateQueries({ queryKey: ['ListThings'] })\n" +
		"      navigate('/things')\n" +
		"    },\n"
	if got := renderOnSuccessHandler(combined, nil, []string{"ListThings"}, nil); got != wantCombined {
		t.Errorf("combined invalidate+navigate = %q, want %q", got, wantCombined)
	}

	// delete: the self GET is removed (not invalidated), sibling list is
	// invalidated, then navigate (BUG-132 132-2).
	del := stmlparser.ActionBlock{OperationID: "DeleteThing", Redirect: "/things"}
	wantDel := "    onSuccess: () => {\n" +
		"      queryClient.removeQueries({ queryKey: ['GetThing'] })\n" +
		"      queryClient.invalidateQueries({ queryKey: ['ListThings'] })\n" +
		"      navigate('/things')\n" +
		"    },\n"
	if got := renderOnSuccessHandler(del, nil, []string{"ListThings"}, []string{"GetThing"}); got != wantDel {
		t.Errorf("delete remove+invalidate+navigate = %q, want %q", got, wantDel)
	}
}
